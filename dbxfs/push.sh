#!/bin/bash
if [ "$#" -lt 1 ]; then
    echo "Usage: $0 <tag>"
    echo "Example: $0 v1.0.0"
    exit 1
fi

tag=$1

docker build -t quay.io/woohhan/dbxfs:${tag} -f Dockerfile .
docker push quay.io/woohhan/dbxfs:${tag}
