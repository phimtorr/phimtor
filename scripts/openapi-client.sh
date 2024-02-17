#!/bin/bash
set -e

readonly service="$1"
readonly output_dir="$2"
readonly package="$3"

oapi-codegen -generate types -o "$output_dir/openapi_generated_types.go" -package "$package" "api/openapi/$service.yml"
oapi-codegen -generate client -o "$output_dir/openapi_generated_client.go" -package "$package" "api/openapi/$service.yml"

