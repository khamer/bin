# ~khamer/bin #
These are various scripts and such I've written and use on a daily basis.

## overload ##
overload effectively allows bash functions to have names that contain spaces as
long as they start with well known words by translating them into underscores
and checking for the existence of such a function. I use this primarily to
improve svn's command line interface to have additional functionality or
memorable shortcuts that don't cloud up the ENVironment.

I reworked this script fairly heavily recently after solving some critical
issues that the previous version had. I believe this verison is significantly
improved over overload 1.x, so I'd recommend the upgrade.

## wmsnap ##
wmsnap is a simple window management and positioning script that I use with to
emulate the behavior in Windows Vista/7 for two, side-by-side, 1920x1080
monitors. I've tried making the script more extensible and configurable in the
past, but at present, its quite hardcoded. On the plus, side, its currently
works the best it ever has. It depends only on wmctrl, xdotool, and xprop.

## upload ##
This is the script I use with vim to upload files to our development
environments. It uses `rsync` by default, but it can be configured to use
`lftp` for environments that only support sftp.

This is how I call the upload script from vim:
```vim
nnoremap <Leader>r :execute '!upload %'<CR><CR>
nnoremap <Leader>W :w<CR>:execute '!upload %'<CR><CR>
nnoremap <Leader>R :execute '!upload %:h'<CR><CR>
nnoremap <Leader>T :execute '!upload .'<CR><CR>
```
## colortest ##
Just displays all the available colors in a terminal for testing terminal color
schemes.

## go ##
go just automates some of the repetition of connection to remote hosts.

## v ##
just a launcher script for gvim that handles multiple file arguments at once.
