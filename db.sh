#!/bin/bash
set -e 

# Rollback shell options before exiting or returning
trap "set +e" EXIT RETURN

echo "[+] Copy db.sql into container"
docker cp db.sql mysql_pruebaceiba:/db.sql

echo "[+] Login into container and Setup DB"
docker exec -itd mysql_pruebaceiba /bin/sh -c 'mysql --user=ceiba --password=ceiba < /db.sql 2>/dev/null | grep -v "mysql: [Warning] Using a password on the command line interface can be insecure."'

echo "[+] Success running script :)"