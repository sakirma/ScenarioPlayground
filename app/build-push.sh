#!/bin/sh

#!/bin/sh

# Get the script path
SCRIPT_PATH="$(readlink -f "$0")"

# Get the directory containing the script
SCRIPT_DIR="$(dirname "$SCRIPT_PATH")"

docker build -t poller:v1.0 "$SCRIPT_DIR"

docker tag poller:v1.0 localhost:5001/poller:v1.0

docker push localhost:5001/poller:v1.0