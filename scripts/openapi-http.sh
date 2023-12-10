#!/bin/bash
set -e

codegen_version=$(oapi-codegen --version | grep -o 'v[0-9]\+\.[0-9]\+')

if [[ "$codegen_version" != 'v1.15' ]]; then
    echo "codegen version v1.15 is expected. your version: $codegen_version"
    exit 1
fi

directory="api/openapi"

function codegen {
  readonly service="$1"
  readonly output_dir="$2"
  readonly package="$3"

  IFS='/' read -r -a folders <<< $output_dir
  for i in "${!folders[@]}"
  do
      folder_path="${folders[*]:0:i+1}"
      folder_path=$(echo "$folder_path" | tr ' ' '/')
      mkdir -p "$folder_path"
  done

  oapi-codegen -generate types -package $package "${directory}/${service}.yml" > "${output_dir}/openapi_types.gen.go"
  oapi-codegen -generate server -package $package "${directory}/${service}.yml" > "${output_dir}/openapi_server.gen.go"

  echo "openapi types and server generated successfully under $output_dir."
}

for file in "$directory"/*
do
    if [ -f "$file" ]; then
        filename=$(basename "$file")
        filename_no_ext="${filename%.*}"

        codegen $filename_no_ext "internal/${filename_no_ext}/presentation" "presentation"
    fi
done
