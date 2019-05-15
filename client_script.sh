#! bin/bash
# script used for launching all the client nodes

konsole -p 'TerminalColumns=100' -p 'TerminalRows=42' -geometry +1000-400  -e ./control/control_client/control_client &


konsole -p 'TerminalColumns=100' -p 'TerminalRows=42' -geometry +490-400  -e ./log/log_client/log_client &


