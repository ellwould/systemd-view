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

cp /root/systemd-view/systemd/systemdview.service /usr/lib/systemd/system/;
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

mkdir -p /etc/systemdview/html-css;

# Copy HTML/CSS start and end files

cp /root/systemd-view/html-css/* /etc/systemdview/html-css/;

# Copy Systemd View configuration file

cp /root/systemd-view/env/systemdview.env /etc/systemdview/systemdview.env

# Create Go directories in root home directory for compiling the source code

mkdir -p /root/go/{bin,pkg,src/systemdview};

# Copy Systemd View source code

cp /root/systemd-view/go/systemdview.go /root/go/src/systemdview/systemdview.go;

# Create Go mod for Systemd View

export PATH=$PATH:/usr/local/go/bin;
cd /root/go/src/systemdview;
go mod init root/go/src/systemdview;
go mod tidy;

# Compile systemdview.go

cd /root/go/src/systemdview;
go build systemdview.go;
cd /root;

# Create system user named systemdview with no shell, no home directory and lock the account

useradd -r -s /bin/false systemdview;
usermod -L systemdview;

# Change executables file permissions, owner, group and move executables

chown root:systemdview /root/go/src/systemdview/systemdview;
chmod 050 /root/go/src/systemdview/systemdview;
mv /root/go/src/systemdview/systemdview /usr/bin/systemdview;

# Change resource file permissions, owner and group

chown -R root:systemdview /etc/systemdview;
chmod 050 /etc/systemdview;
chmod 040 /etc/systemdview/systemdview.env;
chmod 050 /etc/systemdview/html-css;
chmod 040 /etc/systemdview/html-css/*;

# Srart Systemd View programs and enable on boot

systemctl start systemdview;
systemctl enable systemdview;
