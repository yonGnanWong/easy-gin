#!/bin/sh

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o  application -v main.go

#array=()
#
#for((i = 0;i < 2;i++));do
#
#ip = ${array[i]}
#
#rsync -az ./* ${ip}:: --exclude='\.[a-zA-Z]*'
#
#done

upx -9v application

chmod +x application
