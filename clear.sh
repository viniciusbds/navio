#!/bin/bash

# THIS SCRIPT CLEAR ALL Navio's data from the system
# obs: this will not uninstall the Navio.


# REMOVE NAVIO FILES (ROOTFS + IMAGES.tar's)
sudo rm -r "/usr/local/navio"

# Re-open (reset) the database
mysqlversion=`mysql --version`
goversion=`go version`
if [ -n "$goversion"  ] && [ -n "$mysqlversion"  ]; then 
    go run ./database/up.go
fi