#!/bin/bash

# REMOVE NAVIO FILES (ROOTFS + IMAGES.tar's)
sudo rm -r /usr/local/navio

# REMOVE THE BINARY
sudo rm /usr/local/bin/navio

# REMOVE NAVIO DATABASE
mysqlversion=`mysql --version`
goversion=`go version`
if [ -n "$goversion"  ] && [ -n "$mysqlversion"  ]; then 
    mysql -uroot -proot -e "DROP DATABASE navio";
fi
