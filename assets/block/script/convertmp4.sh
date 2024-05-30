#!/bin/bash
for i in "$@"; do
FILE=$i
# use "sudo apt install -y ffmpeg" to install ffmpeg on Ubuntu.
# extract the base file name with the extension out of the full path name.
FILENAME=$(basename -- "$FILE")
# drop the extension
TARGET="${FILENAME%.*}".mp4
ffmpeg -i $FILE $TARGET
done
