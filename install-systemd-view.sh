#!/bin/bash

# Install Script for Systemd View

#----------------------------------------------------------------------

# Check user is root otherwise exit script

if [ "$EUID" -ne 0 ]
then
  printf "\nPlease run as root\n\n";
  exit;
fi;

cd /root;

#----------------------------------------------------------------------

# Check Systemd View has been cloned from GitHub

if [ ! -d "/root/systemd-view" ]
then
  printf "\nDirectory systemd-view does not exist in /root.\n";
  printf "Please run commands: \"cd /root; git clone https://github.com/ellwould/systemd-view\"\n";
  printf "and run install script again\n\n";
  exit;
fi;

#----------------------------------------------------------------------

# Copy unit file and reload systemd deamon

cp /root/systemd-view/systemd/systemd-view.service /usr/lib/systemd/system/;
systemctl daemon-reload;

#----------------------------------------------------------------------

# Install wget

apt update;
apt install wget;

#----------------------------------------------------------------------

# Remove any previous version of Go, download and install Go 1.24.4

wget -P /root https://go.dev/dl/go1.24.4.linux-amd64.tar.gz;
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.24.4.linux-amd64.tar.gz;

#----------------------------------------------------------------------

# Create directores used for Systemd View

mkdir -p /etc/systemd-view/html-css;

# Copy HTML/CSS start and end files

cp /root/systemd-view/html-css/* /etc/systemd-view/html-css/;

# Copy Systemd View configuration file

cp /root/systemd-view/env/systemd-view.env /etc/systemd-view/systemd-view.env

# Create Go directories in root home directory for compiling the source code

mkdir -p /root/go/{bin,pkg,src/systemd-view};

# Copy Systemd View source code

cp /root/systemd-view/go/systemd-view.go /root/go/src/systemd-view/systemd-view.go;

# Create Go mod for Systemd View

export PATH=$PATH:/usr/local/go/bin;
cd /root/go/src/systemd-view;
go mod init root/go/src/systemd-view;
go mod tidy;

# Compile systemd-view.go

cd /root/go/src/systemd-view;
go build systemd-view.go;
cd /root;

# Create system user named systemd-view with no shell, no home directory and lock the account

useradd -r -s /bin/false systemd-view;
usermod -L systemd-view;

# Change executables file permissions, owner, group and move executables

chown root:systemd-view /root/go/src/systemd-view/systemd-view;
chmod 050 /root/go/src/systemd-view/systemd-view;
mv /root/go/src/systemd-view/systemd-view /usr/bin/systemd-view;

# Change resource file permissions, owner and group

chown -R root:systemd-view /etc/systemd-view;
chmod 050 /etc/systemd-view;
chmod 040 /etc/systemd-view/systemd-view.env;
chmod 050 /etc/systemd-view/html-css;
chmod 040 /etc/systemd-view/html-css/*;

# Srart Systemd View programs and enable on boot

systemctl start systemd-view;
systemctl enable systemd-view;
