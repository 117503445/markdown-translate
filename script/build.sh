#!/usr/bin/env bash

mkdir -p target

for arch in "amd64" "arm64"; do
  for os in "linux" "darwin" "windows"; do
    echo "Building for $os $arch"
    bin=./target/mdtran-$os-$arch
    if [ $os = "windows" ]; then
      bin=$bin.exe
    fi
    GOOS=$os GOARCH=$arch go build -o $bin ./cmd/mdtran
  done
done