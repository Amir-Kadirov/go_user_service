#!/bin/bash

# Get the current directory from the first argument
CURRENT_DIR=$1

# Remove old generated files
rm -rf "${CURRENT_DIR}/genproto/*"

# Create the genproto directory if it doesn't exist
mkdir -p "${CURRENT_DIR}/genproto"

# Set paths to the plugins
PROTOC_GEN_GO=$(which protoc-gen-go)
PROTOC_GEN_GO_GRPC=$(which protoc-gen-go-grpc)

# Check that plugins are found
if [[ -z "$PROTOC_GEN_GO" || -z "$PROTOC_GEN_GO_GRPC" ]]; then
  echo "protoc-gen-go or protoc-gen-go-grpc not found in PATH"
  exit 1
fi

# Generate code for each subdirectory in the proto folder
for x in $(find ${CURRENT_DIR}/proto/* -type d); do
  protoc --plugin="protoc-gen-go=${PROTOC_GEN_GO}" \
         --plugin="protoc-gen-go-grpc=${PROTOC_GEN_GO_GRPC}" \
         -I=${x} -I=${CURRENT_DIR}/proto -I /usr/local/include \
         --go_out=${CURRENT_DIR}/genproto \
         --go-grpc_out=${CURRENT_DIR}/genproto \
         ${x}/*.proto
done
