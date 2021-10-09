read -p "Enter mysql root password: " DATABASE_PASSWORD

mysql -u root --password=$DATABASE_PASSWORD -e "CREATE DATABASE IF NOT EXISTS navio"
mysql -u root --password=$DATABASE_PASSWORD -e "CREATE USER IF NOT EXISTS 'navioUser'@'localhost' IDENTIFIED BY 'PmO001-nav'"
mysql -u root --password=$DATABASE_PASSWORD -e "GRANT ALL PRIVILEGES ON navio . * TO 'navioUser'@'localhost';"
mysql -u root --password=$DATABASE_PASSWORD -e "FLUSH PRIVILEGES"
