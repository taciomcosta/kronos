#!/bin/bash

SOURCE_DAEMON=com.taciomcosta.kronos.plist
TARGET_DAEMON=/Library/LaunchDaemons/com.taciomcosta.kronos.plist
SOURCE_KRONOSD=../build/kronosd
TARGET_KRONOSD=/usr/local/bin/kronosd
SOURCE_KRONOSCLI=../build/kronos
TARGET_KRONOSCLI=/usr/local/bin/kronos

cp $SOURCE_DAEMON $TARGET_DAEMON
cp $SOURCE_KRONOSD $TARGET_KRONOSD
cp $SOURCE_KRONOSCLI $TARGET_KRONOSCLI

launchctl load $TARGET_DAEMON
