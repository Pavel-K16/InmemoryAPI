#!/bin/bash
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"

ROOT_DIR="$(dirname "$SCRIPT_DIR")"
cd "$ROOT_DIR"

go run cmd/taskapi.go