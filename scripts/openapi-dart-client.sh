
readonly service="$1"
readonly output_dir="$2"

PUB_NAME=$(basename "$output_dir")

openapi-generator-cli generate -i "api/openapi/$service.yml" -g dart -o "$output_dir" \
  --additional-properties pubName="$PUB_NAME",pubPublishTo=none