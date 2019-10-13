#!/bin/bash

GOPATH=$GOPATH
sudo GOPATH=$GOPATH go run $PWD/main/main.go &
pid=$!
sleep 3
sudo ifconfig utun1 10.1.0.10 10.1.0.20 up
echo $pid
wait $pid
