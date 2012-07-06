~khamer/bin
===========

These are various scripts and such I've written and use on a daily basis.

wmsnap
------

wmsnap is a simple window management and positioning script that I use with to
emulate the behavior in Windows Vista/7 for two, side-by-side, 1920x1080
monitors. I've tried making the script more extensible and configurable in the
past, but at present, its quite hardcoded. On the plus, side, its currently
works the best it ever has. It depends only on wmctrl, xdotool, and xprop.

overload
--------

overload effectively allows bash functions to have names that contain spaces as
long as they start with well known words by translating them into underscores
and checking for the existence of such a function. I use this primarily to
improve svn's command line interface to have additional functionality or
memorable shortcuts that don't cloud up the ENVironment.
