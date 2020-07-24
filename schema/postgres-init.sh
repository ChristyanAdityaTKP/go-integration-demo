#!/bin/sh
# do not use this script on production
# this script is for development phase only

db="postgres"
port=${PGPORT:-5434}
user="postgres"
password="pass"
host=${PGHOST:-localhost}

echo "host $host port: $port"
PGPASSWORD=$password psql -U $user -d "postgres" -h $host -p $port -c "CREATE DATABASE $db" 2>/dev/null
for filename in schema/*.sql; do
    PGPASSWORD=$password psql -h $host -p $port -d $db -U $user -f "$filename"
done
