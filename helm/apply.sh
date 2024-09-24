#!/bin/sh

# Get the script path
SCRIPT_PATH="$(readlink -f "$0")"

# Get the directory containing the script
SCRIPT_DIR="$(dirname "$SCRIPT_PATH")"

helm upgrade --install prober-a "$SCRIPT_DIR/prober" -n team-a
helm upgrade --install prober-b "$SCRIPT_DIR/prober" -n team-b