#!/usr/bin/env bash
set -eufo pipefail

docker-compose -f build/docker-compose.yml --project-name discord-bot down -v