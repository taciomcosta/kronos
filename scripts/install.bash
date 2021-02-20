#!/bin/bash

SOURCE_DAEMON=com.taciomcosta.kronos.plist
TARGET_DAEMON=/Library/LaunchAgents/com.taciomcosta.kronos.plist
SOURCE_KRONOSD=../build/kronosd
TARGET_KRONOSD=/usr/local/bin/kronosd
SOURCE_KRONOSCLI=../build/kronos
TARGET_KRONOSCLI=/usr/local/bin/kronos

sudo mkdir /etc/kronos
sudo mkdir /var/lib/kronos

sudo cp $SOURCE_DAEMON $TARGET_DAEMON
sudo cp $SOURCE_KRONOSD $TARGET_KRONOSD
sudo cp $SOURCE_KRONOSCLI $TARGET_KRONOSCLI

launchctl load $TARGET_DAEMON
