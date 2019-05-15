#! bin/bash
# script used for launching all the servers

konsole -p 'TerminalColumns=100' -p 'TerminalRows=42' -geometry +1000-30  -e ./control/control_server/control_server &

konsole -p 'TerminalColumns=100' -p 'TerminalRows=42' -geometry +490-30 -e ./log/log_server/log_server &

konsole -p 'TerminalColumns=100' -p 'TerminalRows=42' -geometry +0-30 -e ./car/car_server/car_server &
