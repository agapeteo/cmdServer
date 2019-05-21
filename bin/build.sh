#!/bin/sh

go build ../main/cmdServer.go && \
mv cmdServer darwin_x86_64 && \
echo darwin is ready

GOOS=linux go build ../main/cmdServer.go && \
mv cmdServer linux_x86_64 && \
echo linux x86_64 is ready

GOOS=linux GOARCH=arm GOARM=7 go build ../main/cmdServer.go && \
mv cmdServer linux_arm && \
echo linux arm is ready
