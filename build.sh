#!/bin/bash

set -e

APP_NAME="app"

SOURCE_DIR="./"

OUTPUT_DIR="./build"

GOOS="linux"  # Contoh: linux, darwin, windows
GOARCH="amd64"  # Contoh: amd64, arm64

mkdir -p $OUTPUT_DIR

# Mengatur variabel lingkungan untuk membangun aplikasi
export CGO_ENABLED=0
export GOOS=$GOOS
export GOARCH=$GOARCH

echo "Building $APP_NAME for $GOOS/$GOARCH..."

go build -ldflags "-s -w" -o $OUTPUT_DIR/$APP_NAME . 

if [ -f "$OUTPUT_DIR/$APP_NAME" ]; then
  echo "Build berhasil: $OUTPUT_DIR/$APP_NAME"
else
  echo "Build gagal!"
  exit 1
fi
upx --best --lzma $OUTPUT_DIR/$APP_NAME
echo "Build selesai, file aplikasi berada di $OUTPUT_DIR/$APP_NAME"
