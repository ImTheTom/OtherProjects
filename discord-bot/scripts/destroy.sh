#!/usr/bin/env bash
set -eufo pipefail

project_name='discord-bot'
project_compose_file='build/docker-compose.yml'

while getopts :d option; do
	case $option in
		d)
			echo "Importing dev overrides"
			project_name="${project_name}-dev"
			project_compose_file="${project_compose_file}:build/docker-compose.dev.yml"
		;;
	esac
done

read -p "This will delete the docker volume. Are you sure? " selection

if [ $selection != "y" ]; then
	exit 0
fi

export COMPOSE_PROJECT_NAME=${project_name}
export COMPOSE_FILE=${project_compose_file}

docker-compose down -v