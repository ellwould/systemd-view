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

systemctl stop systemdview.service;
systemctl disable systemdview.service;

# Remove Systemd View unit files and reload systemd deamon

rm /usr/lib/systemd/system/systemdview.service;
systemctl daemon-reload;

#----------------------------------------------------------------------

# Remove Systemd View binary

rm /usr/bin/systemdview;

# Remove all other directores and files used by Systemd View

rm -r /etc/systemdview;

# Remove Systemd View source code in root home directory

rm -r /root/go/src/systemdview;

# Remove the user and group systemdview from the system

userdel systemdview;
