mysql -uroot -proot -e "CREATE DATABASE IF NOT EXISTS navio"
mysql -uroot -proot -e "CREATE USER IF NOT EXISTS 'navioUser'@'localhost' IDENTIFIED BY 'PmO001-nav'"
mysql -uroot -proot -e "GRANT ALL PRIVILEGES ON navio . * TO 'navioUser'@'localhost';"
mysql -uroot -proot -e "FLUSH PRIVILEGES"
