#!/bin/bash

GOOS=linux go build -ldflags "-w -s" -o release/goqut-linux .
GOOS=darwin go build -ldflags "-w -s" -o release/goqut-macos .
GOOS=windows go build -ldflags "-w -s" -o release/goqut-windows.exe .
