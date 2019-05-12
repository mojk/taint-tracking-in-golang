// This is the part 5
#include "TinyTimber.h"
#include "sciTinyTimber.h"
#include "canTinyTimber.h"
#include "sioTinyTimber.h"
#include "buffer.h"
#include <stdlib.h>
#include <stdio.h>

#define START 0
#define STOP 1
#define KEY 2
#define TEMPO 3
#define VOLUME 4
#define REQUEST 5
#define RESPONSE 6

char *DAC_OUTPUT = (char *) 0x4000741C;
int LEADER = 1;
int freqind[32] = {0, 2, 4, 0, 0, 2, 4, 0, 4, 5, 7, 4, 5, 7, 7, 9, 7, 5, 4, 0, 7, 9, 7, 5, 4, 0, 0, -5, 0, 0, -5, 0};
int periods[25] = {2024, 1908, 1805, 1701, 1608, 1515, 1433, 1351, 1276, 1205, 1136, 1073, 1012, 956, 903, 852, 804, 759, 716, 676, 638, 602, 568, 536, 506};
float beats[32] = {1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 2, 0.5, 0.5, 0.5, 0.5, 1, 1, 0.5, 0.5, 0.5, 0.5, 1, 1, 1, 1, 2, 1, 1, 2};

typedef struct {
    Object super;
    int count;
    int myNum;
	int runSum;
	char c;
	char inpStr[50];
	char bufNum[50];
	char bufSum[50];
	int first_message;
	Time time;
} App;

typedef struct {
	Object super;
	int period;
	int volume;
	int mute;
    int background_loop_range;
    int deadline;
    int const_deadline;
	int killed;
	int count;
	int key;
} Sound;

typedef struct {
	Object super;
    int tempo;
	int count;
	int killed;
	int started;
} Melody;

typedef struct {
	Object super;
	int counter;
	int pressed;
	Time time;
	Time buffer[3];
} UserButton;

App app = { initObject(), 0, 0, 0, '\0' };

Sound s = { initObject(), 1136, 10, 0, 0, 100, 100, 0 };
Sound b = { initObject(), 1300, 5, 0, 1000, 1300, 1300 };
Melody m = { initObject(), 120, 0};
UserButton btn = {initObject(), 0, 0, 0};
Buffer circbuff = {initObject(), 0, 0, 0};

/*********************************** Function and method delcaration ***********************************/
void tone_generator(Sound *self, int unused);
void background_generator(Sound *self, int unused);
void deadline_control(Sound *self, int unused);
void volume_control(Sound *self, char inp);
void key_control(Sound *self, char inp);
void reader(App*, int);
void receiver(App*, int);
void kill_tone(Sound *self, int unused);
void set_tempo(Melody *self, int t);
void set_key(Sound *self, int key);
void set_period(Sound *self, int p);
void unkill(Sound *self, int unused);
void unkill_player(Melody *self, int unused);
void kill_player(Melody *self, int unused);
void melody_player(Melody *self, int unused);
void start_player(Melody *self, int unused);

void prepare_can_msg(CANMsg *self, int value);
void button_receiver(UserButton *self, int unused);
void unset_button(UserButton *self, int unused);
void set_button(UserButton *self, int unused);
void deliver(App *self, int unused);
void check_tempo(UserButton *self, int unused);
int from_micro_to_milli(Time t);
void led_control(Melody *self, int unused);
void detect_neighbors(App *self, int unused);

Serial sci0 = initSerial(SCI_PORT0, &app, reader);
Can can0 = initCan(CAN_PORT0, &app, receiver);
SysIO sysio = initSysIO(SIO_PORT0, &btn, button_receiver);

int from_micro_to_milli(Time t) {
	return t/100;
}

void check_tempo(UserButton *self, int unused) {
	char buffer[50];

	int first = self->buffer[0] - self->buffer[1];
	int second = self->buffer[0] - self->buffer[2];
	int third = self->buffer[1] - self->buffer[2];

	snprintf(buffer, 100, "First: %d\n", first);
	SCI_WRITE(&sci0, buffer);
	snprintf(buffer, 100, "Second: %d\n", second);
	SCI_WRITE(&sci0, buffer);
	snprintf(buffer, 100, "Third: %d\n", third);
	SCI_WRITE(&sci0, buffer);

	if(first > -100 && first < 100 && second > -100 && second < 100 && third > -100 && third < 100) {
		int average = (self->buffer[0] + self->buffer[1] + self->buffer[2]) / 3;
		int tempo = 60000 / average;

		snprintf(buffer, 100, "Average was: %d\n", average);
		SCI_WRITE(&sci0, buffer);

		snprintf(buffer, 100, "Setting tempo to: %d\n", tempo);
		SCI_WRITE(&sci0, buffer);
		SYNC(&m, set_tempo, tempo);
	}
	else {
		SCI_WRITE(&sci0, "Failed to set tempo...\n");
	}
}

/* Called function when the User button is pressed
 * Calculates the inter-arrival times between two button press
 * TODO: 	Implement sending a CAN message each time the button is pressed
 *
 * */
void button_receiver(UserButton *self, int unused) {

	// First initial press
	if(!self->pressed) {
		self->time = from_micro_to_milli(CURRENT_BASELINE());
		self->pressed = 1;
	}
	// Button already pressed and interval is larger than 100ms
	else if(CURRENT_BASELINE()/100 - self->time > 100) {

		self->buffer[self->counter % 3] = from_micro_to_milli(CURRENT_BASELINE()) - self->time;
		if((self->counter % 4) == 0) {
			// Check tempo
			SCI_WRITE(&sci0, "Check tempo!\n!");
			check_tempo(self, 0);
		}
		self->time = from_micro_to_milli(CURRENT_BASELINE());
	}
	self->counter = (self->counter + 1) % 128;
}

/* The recevier method for the CAN bus */
void receiver(App *self, int unused) {
    CANMsg msg;
    CAN_RECEIVE(&can0, &msg);

	if(msg.msgId == REQUEST && msg.nodeId == 15) {
		SCI_WRITE(&sci0, "Recevied request message!\n");
	}
	else if(msg.msgId == RESPONSE && msg.nodeId == 15) {
		SCI_WRITE(&sci0, "Recevied response message!\n");
	}

	if(!self->first_message) {
		self->time = from_micro_to_milli(CURRENT_BASELINE());
		self->first_message = 1;

		if(SYNC(&circbuff, put, msg.msgId) == -1) {
			SCI_WRITE(&sci0, "Buffer full!\n");
		}
		ASYNC(self, deliver, 0);

	}
	else {
		if(SYNC(&circbuff, put, msg.msgId) == -1) {
			SCI_WRITE(&sci0, "Buffer full!\n");
		}
	}
}

void deliver(App *self, int unused) {
	int val = get(&circbuff, 0);

	if(val != -1) {
		char buf[100];
		snprintf(buf, 100, "Delivered msg: %d\n", val);
		SCI_WRITE(&sci0, buf);
	}

	AFTER(SEC(1), self, deliver, 0);
}

/*********************************** Reader for controlling the keyboad ***********************************/

/* This is the function that handles keyboard inputs */
void reader(App *self, int c) {
	CANMsg msg;
	SCI_WRITECHAR(&sci0, c);
	SCI_WRITE(&sci0,"\n");

	// Controlling the tempo
	if(c == 't' && LEADER) {
		prepare_can_msg(&msg, 3);

		self->inpStr[self->count] = '\0';
		int tempo = atoi(self->inpStr);
		self->count = 0;

		can_send(can0, &msg);
		SYNC(&m, set_tempo, tempo);
	}

	// Controlling the key
	else if (c == 'k' && LEADER) {
		prepare_can_msg(&msg, 2);

		self->inpStr[self->count] = '\0';
		int key = atoi(self->inpStr);
		self->count = 0;

		can_send(can0, &msg);
		SYNC(&s, set_key, key);
	}
	else if(c == 'p' && LEADER) {
		prepare_can_msg(&msg, 0);
		can_send(can0, &msg);
		SYNC(&m, start_player, 0);
	}

	// Controlling the stop
	else if(c == 's') {
		prepare_can_msg(&msg, 1);
		can_send(can0, &msg);
		SYNC(&m, kill_player, 0);
	}

	else {
		self->inpStr[self->count] = c;
		self->count++;
	}
}

/*********************************** Melody and tone generators ***********************************/

/* This is the tone gernerator mehod */
void tone_generator(Sound *self, int unused) {

	// Check if tone is killed
	if(!self->killed) {
		if(*DAC_OUTPUT) {
			*DAC_OUTPUT = 0;
		}
		else {
			*DAC_OUTPUT = self->volume;
		}
	}
	// After the correct period we call this tone_generator again
	AFTER(USEC(self->period), self, tone_generator, 0);
}

void led_control(Melody *self, int unused) {
	AFTER(MSEC(60000 / self->tempo), &sysio, sio_toggle, 0);
	AFTER(MSEC((60000 / self->tempo) / 2), &sysio, sio_toggle, 1);
	AFTER(MSEC(60000 / self->tempo), self, led_control, 0);
}

/* This is the melody player that controls the tone_generator, giving it the correct beat */
void melody_player(Melody *self, int unused) {

	// Used for starting and stopping the player
	self->started = 1;

	// If user has pressed the stop button, won't play anymore
	if(!self->killed) {

		// Disable killing of tone
		unkill(&s, 0);
		set_period(&s, self->count);

		// After the beat period - 50 kill the current tone
		AFTER(MSEC((60000/self->tempo) * beats[self->count] - 50), &s, kill_tone, 0);

		// After beat period call melody_player again but with new tone
		AFTER(MSEC((60000/self->tempo) * beats[self->count]), self, melody_player, 0);

		// Increment the counting of tone indices
		self->count = (self->count + 1) % 32;
	}
}

/*********************************** Methods for controlling ***********************************/

void start_player(Melody *self, int unused) {
	SYNC(self, unkill_player, 0);

	if(self->started) {
		SYNC(self, melody_player, 0);
	}
	else {
		SYNC(&s, tone_generator, 0);
		SYNC(self, melody_player, 0);
	}
}

void set_period(Sound *self, int p) {
	self->period = periods[freqind[p] + self->key + 10];
}

void set_key(Sound *self, int k){
	self->key = k - 5;
}

void set_tempo(Melody *self, int t) {
	self->tempo = t;
}

/* Method for unkilling the tone generator */
void unkill(Sound *self, int unused) {
	self->killed = 0;
}

/* Method for killing the tone generator */
void kill_tone(Sound *self, int unused) {
	self->killed = 1;
}

/* Method for killing the melody player */
void kill_player(Melody *self, int unused) {
	self->killed = 1;
	self->count = 0;
	kill_tone(&s, 0);
}

/* Method for unkilling the melody player */
void unkill_player(Melody *self, int unused) {
	self->killed = 0;
}

void volume_control(Sound *self, char inp) {

	switch(inp) {
        // UP - increment the volume by one
		case 'u':
			if(self->volume < 20)
				self->volume += 1;
			break;
        // DOWN - decrement the volume by one
		case 'd':
			if(self->volume > 5)
				self->volume -= 1;
			break;
        // MUTE or UNMUTE - mute the sound
		case 'm':
			if(!self->volume)
				self->volume = 5;
			else
				self->volume = 0;
			break;

		default:
			SCI_WRITE(&sci0, "Enter another character.\n");
		}
}

void prepare_can_msg(CANMsg *self, int value) {
	self->msgId = value;
	self->nodeId = ;
	self->length;
	self->buff
}

void detect_neighbors(App *self, int unused) {
	CANMsg msg;

	msg.msgId = REQUEST;
	msg.nodeId = 15;
	msg.length = 1;
	msg.buff[0] = 0;
	can_send(can0, &msg);
}


/*********************************** Starting methods for the system ***********************************/

/* Main method for starting the system */
void startApp(App *self, int arg) {
    CAN_INIT(&can0);
    SCI_INIT(&sci0);
	SIO_INIT(&sysio);
	SCI_WRITE(&sci0, "Startapp...\n");
	SYNC(&app, detect_neighbors, 0);

	//ASYNC(&m, melody_player, 0);
	//ASYNC(&s, tone_generator, 0);
	//ASYNC(&m, led_control, 0);
}

/* Main function for installing interrupt handerls and starting TinyTimber*/
int main() {
    INSTALL(&sci0, sci_interrupt, SCI_IRQ0);
	INSTALL(&can0, can_interrupt, CAN_IRQ0);
	INSTALL(&sysio, sio_interrupt, SIO_IRQ0);
    TINYTIMBER(&app, startApp, 0);
    return 0;
}
