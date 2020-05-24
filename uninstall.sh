#!/bin/bash


# REMOVE NAVIO DATABASE
mysqlversion=`mysql --version`
goversion=`go version`
if [ -n "$goversion"  ] && [ -n "$mysqlversion"  ]; then 
    mysql -uroot -proot -e "DROP DATABASE navio";
fi

# ...
sudo rm /usr/local/bin/navio
