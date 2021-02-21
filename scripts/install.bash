#!/bin/bash

CONFIGPATH=/etc/kronos
DATAPATH=/var/lib/kronos
SOURCE_DAEMON=com.taciomcosta.kronos.plist
TARGET_DAEMON=/Library/LaunchAgents/com.taciomcosta.kronos.plist

mkdir $CONFIGPATH
mkdir $DATAPATH
cp $SOURCE_DAEMON $TARGET_DAEMON

launchctl load $TARGET_DAEMON
