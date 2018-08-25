#!/bin/bash

set -e

sudo cp -rf etc/nginx.conf /etc/nginx/nginx.conf
sudo cp -rf etc/my.cnf /etc/mysql/my.cnf
# sudo rm /var/log/nginx/access.log
# sudo rm /var/log/mysql/slow_query.log
sudo service nginx restart
sudo service mysql restart
# killall webapp
# cd $HOME/webapp/go && make
# $HOME/webapp/go/webapp
