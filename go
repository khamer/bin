#!/bin/bash
## Kevin Hamer [kh] <kevin@imarc.net>

shopt -s extglob

declare -A OPTS
ARGS=()

while [[ $# > 0 ]]; do
	case $1 in
		-*)
			KEY="${1##+(-)}"
			OPTS[$KEY]="yes"
			;;
		*)
			ARGS+=("$1")
			;;
	esac
	shift
done

if [[ ${ARGS[1]} == "" ]]; then
	SITE="$(basename $PWD)"
else
	SITE="${ARGS[1]}"
fi

if [[ ${OPTS[h]+1} == 1 || ${OPTS[help]+1} == 1 ]]; then

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

elif [[ ${OPTS[s]+1} == 1 || ${OPTS[su]+1} == 1 ]]; then

	ssh -t $SITE "cd www/$SITE; cd \$(pwd -P); bash -ic su"

elif [[ ${OPTS[p]+1} == 1 || ${OPTS[psql]+1} == 1 ]]; then

	ssh -t $SITE "psql -U postgres ${SITE//[^a-zA-Z0-9]/_}"

else
	ssh -t $SITE "cd www/$SITE; cd \$(pwd -P); bash -i"
fi
