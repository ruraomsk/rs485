#!/bin/bash
echo 'Compiling for bananapi'
GOOS=linux GOARCH=arm  go build
if [ $? -ne 0 ]; then
	echo 'An error has occurred! Aborting the script execution...'
	exit 1
fi
echo 'Copy rs232 to banana'
scp  rs232 root@192.168.88.20:/home/rura
