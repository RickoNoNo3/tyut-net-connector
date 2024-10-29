#!/bin/bash

if [ -z "$1" ]; then
    echo "Usage: $0 <text>"
    exit 1
fi


result=""
for ((i=0; i<${#1}; i++)); do
    c=${1:$i:1}
    result+=$(printf "%02x" $(echo "$(( $(printf "%d" "'$c") ^ 0x77 ))"))
done
echo "$result"