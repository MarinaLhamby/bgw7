#!/bin/bash
set -e

# Access the environment variables directly
MYSQL_USER="${DB_USER}"
MYSQL_PASSWORD="${DB_PASSWORD}"
MYSQL_DATABASE="${DB_NAME}"
MYSQL_ROOT_PASSWORD="${DB_ROOT_PASSWORD}" # Need root password to grant permissions

echo "Attempting to grant permissions for user: $MYSQL_USER on database: $MYSQL_DATABASE"

mysql -u root -p"${MYSQL_ROOT_PASSWORD}" <<-EOSQL
ALTER USER '${MYSQL_USER}'@'%' IDENTIFIED WITH mysql_native_password BY '${MYSQL_PASSWORD}';
GRANT ALL PRIVILEGES ON ${MYSQL_DATABASE}.* TO '${MYSQL_USER}'@'%';
FLUSH PRIVILEGES;
EOSQL

echo "Permissions granted successfully!"