#!/bin/bash

# INSTALL THE BINARY
sudo cp ./navio /usr/local/bin

# CREATE NAVIO DATABASE
mysqlversion=`mysql --version`
goversion=`go version`
if [ -n "$goversion"  ] && [ -n "$mysqlversion"  ]; then 
    go run ./database/up.go
fi
