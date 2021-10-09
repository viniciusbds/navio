#!/bin/bash

## This scripts check if the [golang, wget and mysql] are instaled.
## If yes, it create the Navio database and copy the binary to /usr/local/bin

echo "Creating navio database and navioUser..."
#Setup database
bash ./setupDatabase.sh
echo "ok"

#Make
make
echo "ok"

# Check golang ...
if  go version  > /dev/null; then
    echo "Check golang                     ok"
else 
    echo "Check golang                     fail"
    echo "Install it an run this script again."
    exit 1
fi

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
go run ./database/up.go

# Check navio database ...
navio=`mysql -uroot -proot -e 'show databases;' | grep navio`
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
