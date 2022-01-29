#!/usr/bin/env bash
set -eufo pipefail

docker-compose -f build/docker-compose.dev.yml --project-name discord-bot down -v