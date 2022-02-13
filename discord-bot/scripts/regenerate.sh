#!/usr/bin/env bash
set -eufo pipefail

while getopts :b option; do
	case $option in
		b)
			echo "Rebuilding generate help command"
			go build -o gen-help cmd/help/main.go
			mv gen-help ~/go/bin
		;;
	esac
done

go generate ./...