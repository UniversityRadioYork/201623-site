#!/bin/bash
#This script will only work on Linux (Specifically using X11)
CURRENT_WID=$(xdotool getwindowfocus)

WID=$(xdotool search --name "Mozilla Firefox")
xdotool windowactivate $WID
xdotool key F5

xdotool windowactivate $CURRENT_WID
make run
