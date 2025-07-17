#!/bin/bash

# Uninstall script for Systemd View

#----------------------------------------------------------------------

# Check user is root otherwise exit script

if [ "$EUID" -ne 0 ]
then
  printf "\nPlease run as root\n\n";
  exit;
fi;

cd /root;

#----------------------------------------------------------------------

# Stop Systemd View automatically starting on boot

systemctl stop systemd-view.service;
systemctl disable systemd-view.service;

# Remove Systemd View unit files and reload systemd deamon

rm /usr/lib/systemd/system/systemd-view.service;
systemctl daemon-reload;

#----------------------------------------------------------------------

# Remove Systemd View binary

rm /usr/bin/systemd-view;

# Remove all other directores and files used by Systemd View

rm -r /etc/systemd-view;

# Remove Systemd View source code in root home directory

rm -r /root/go/src/systemd-view;

# Remove the user and group systemd-view from the system

userdel systemd-view;
