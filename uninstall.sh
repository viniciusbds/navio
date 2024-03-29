#!/bin/bash

# REMOVE NAVIO FILES (ROOTFS + IMAGES.tar's)
if sudo rm -r /usr/local/navio; then
    echo "Remove /usr/local/navio  ok"
else
    echo "Remove /usr/local/navio  fail"
fi

# REMOVE THE BINARY
if sudo rm /usr/local/bin/navio; then
    echo "Remove /usr/local/bin/navio  ok"
else
    echo "Remove /usr/local/bin/navio  fail"
fi

# REMOVE NAVIO DATABASE
if mysql -uroot -proot -e "DROP DATABASE navio"; then
    echo "Drop database                      ok"
else 
    echo "Drop database                      fail"    
fi

# REMOVE NAVIO USER
if mysql -uroot -proot -e "DROP USER navioUser@localhost"; then
    echo "Drop navioUser                      ok"
else 
    echo "Drop navioUser                      fail"    
fi
