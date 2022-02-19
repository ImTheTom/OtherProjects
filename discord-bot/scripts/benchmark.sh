#!/usr/bin/env bash
set -eufo pipefail

go test -run=Bench -bench=. ./internal/bot/