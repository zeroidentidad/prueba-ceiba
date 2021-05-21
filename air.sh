#!/bin/bash
set -e 

# Rollback shell options before exiting or returning
trap "set +e" EXIT RETURN

echo "[+] Run container"
docker start mysql_pruebaceiba

echo "[+] Wait to container healthy status..."
./wait-hc.sh mysql_pruebaceiba

echo "[+] Wait 30s more please to db ready connections..."
sleep 30

echo "[+] Live-reload Air binary"
air -c .air.linux.conf

echo "[+] Stop container on exit..."
docker stop mysql_pruebaceiba