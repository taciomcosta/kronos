#!/bin/bash

launchctl unload /Library/LaunchAgents/com.taciomcosta.kronos.plist

rm /Library/LaunchDaemons/com.taciomcosta.kronos.plist
rm -rf /etc/kronos
rm -rf /var/lib/kronos
rm /var/log/kronos.log
rm /usr/local/bin/kronos*
rm ../build/*
