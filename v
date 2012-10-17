#!/bin/bash
if [ -z $1 ]; then
	gvim
else
	gvim --remote-tab-silent "$@"
fi
