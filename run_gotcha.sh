#!/bin/bash

gotcha -ssf="gotcha/ssf.txt" -src="control/control_client/control_client.go" -path="taint-tracking-in-golang"

if [ -z "$1" ]
then
    echo "No argument supplied. For deletion of log files add 'delete' as an argument"
else
    if [ $1 = "delete" ]; then
        rm 201*.txt
        rm 201*.stat
    fi

fi


