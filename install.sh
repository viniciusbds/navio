#!/bin/bash


# Check wget ...
if  wget --version > /dev/null ; then
    echo "Check wget                       ok"
else 
    echo "Check wget                       fail"
    echo "Install it an run this script again."
    exit 1
fi

# Check mysql ...
if  mysql --version  > /dev/null; then
    echo "Check mysql                      ok"
else 
    echo "Check mysql                      fail"
    echo "Install it an run this script again."
    exit 1
fi

#CREATE NAVIO DATABASE
mysqlversion=`mysql --version`
goversion=`go version`
if [ -n "$goversion"  ] && [ -n "$mysqlversion"  ]; then 
    go run ./database/up.go
fi

# Check navio database ...
navio=`mysql --user=root --password=root -e 'show databases;' | grep navio`
if [ "navio" == $navio ]; then
    echo "Up database                      ok"
else 
    echo "Up database                      fail"
    echo "" 
    echo "Expected Database user: root" 
    echo "Expected Database passwd: root" 
fi


# INSTALL THE BINARY
if sudo cp ./navio /usr/local/bin; then
    echo "Copy ./navio --> /usr/local/bin  ok"
else 
    echo "Copy ./navio --> /usr/local/bin  fail"
fi

