#!/bin/bash

# THIS SCRIPT CLEAR ALL Navio's data from the system
# obs: this will not uninstall the Navio.


# REMOVE NAVIO FILES (ROOTFS + IMAGES.tar's)
if sudo rm -r /usr/local/navio; then
    echo "Remove /usr/local/navio  ok"
else
    echo "Remove /usr/local/navio  fail"
fi


go run ./database/up.go
