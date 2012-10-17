#!/bin/bash
## Kevin Hamer [kh] <kevin@imarc.net>

if [ $1 == "-h" ]; then
	echo "Quickly login to systems and switch to that site's directory or connect"
	echo "to its database."
	echo ""
	echo "Usage: go [OPTION] SERVERNAME"
	echo "    -h  Show this help."
	echo "    -p  Log in directly to postgres as postgres with psql."
	echo "    -s  Immediately call su, to switch to root."
	echo ""
	echo "Examples: go -s imarc.net"
	echo "          go -p dev.imarc.net"

elif [ $1 == "-s" ]; then
	ssh -t $2 "cd www/$2; cd \$(pwd -P); bash -ic su"
elif [ $1 == "-p" ]; then
	ssh -t $2 "psql -U postgres ${2//[^a-zA-Z0-9]/_}"
else
	ssh -t $1 "cd www/$1; cd \$(pwd -P); bash -i"
fi
