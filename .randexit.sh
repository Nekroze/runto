#!/bin/sh
echo "Random success or failure"
if (( RANDOM % 2 )); then
   exit 1
else
   exit 0
fi
