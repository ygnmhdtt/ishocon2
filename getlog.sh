now=$(date +%Y%m%d_%H%M%S)
sudo cp /var/log/nginx/access.log ./logs/access.log.${now}
sudo cp /var/log/mysql/slow_query.log ./logs/slow_query.log.${now}
sudo chown ishocon:ishocon ./logs/*
sudo chmod 755 ./logs/*
