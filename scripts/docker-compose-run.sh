#!/bin/bash
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"

ROOT_DIR="$(dirname "$SCRIPT_DIR")"
cd "$ROOT_DIR"

sudo docker-compose up --build
sudo docker image prune -f

####### WORK IN PROGRESS #######      