#!/bin/sh

echo "Create DynamoDB Table"
awslocal dynamodb create-table --cli-input-json file:////docker-entrypoint-initaws.d/forum.json
