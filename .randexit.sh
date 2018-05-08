#!/bin/sh
echo "Random success or failure"
sleep $[ ( $RANDOM % 5 )  + 1 ]s
if (( RANDOM % 2 )); then
   exit 1
else
   exit 0
fi
