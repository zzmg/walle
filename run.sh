#!/bin/bash
/server
touch /var/log/cron.log
tail -f /var/log/cron.log
