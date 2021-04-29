#!/bin/sh

LEFT=10
RIGHT=40

SET_ONLINE="./set_online.py"
GRANDOM="./random"

test -f "$SET_ONLINE" || exit 1
test -f "$GRANDOM" || exit 1

while true; do
	t=$($GRANDOM $LEFT $RIGHT)
	python "$SET_ONLINE" || exit
	sleep "$t"
done
