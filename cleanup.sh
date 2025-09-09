#!/bin/bash

# Load environment variables
if [ -f .env ]; then
    export $(cat .env | xargs)
fi

# Set default values if not defined in .env
IMAGE_NAME=${IMAGE_NAME:-"atlassian-dc-mcp-go"}
CONTAINER_NAME="atlassian-dc-mcp"

echo "Stopping and removing container: $CONTAINER_NAME"
docker stop $CONTAINER_NAME 2>/dev/null || echo "Container $CONTAINER_NAME is not running"
docker rm $CONTAINER_NAME 2>/dev/null || echo "Container $CONTAINER_NAME does not exist"

echo "Removing image: $IMAGE_NAME"
docker rmi $IMAGE_NAME 2>/dev/null || echo "Image $IMAGE_NAME does not exist"

echo "Cleanup completed"