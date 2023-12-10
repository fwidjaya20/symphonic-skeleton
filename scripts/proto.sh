#!/bin/bash
set -e

mkdir -p pkg/genproto/$service

directory="api/protobuf"

function codegen {
  readonly service="$1"
  readonly output_dir=pkg/genproto

  protoc \
    --proto_path=api/protobuf "$directory/$service.proto" \
    "--go_out=$output_dir/$service" --go_opt=paths=source_relative \
    --go-grpc_opt=require_unimplemented_servers=false \
    "--go-grpc_out=$output_dir/$service" --go-grpc_opt=paths=source_relative

  echo "$service.proto generated successfully under $output_dir."
}

for file in "$directory"/*
do
    if [ -f "$file" ]; then
        filename=$(basename "$file")
        filename_no_ext="${filename%.*}"

        codegen $filename_no_ext
    fi
done