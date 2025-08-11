#!/bin/bash

MAKE_PACKAGE=false

usage() {
 echo "Usage: $0 [OPTIONS]"
 echo "Options:"
 echo " --package      Make a delivery package instead of plain binary"
}

# Function to handle options and arguments
handle_options() {
  while [ $# -gt 0 ]; do
    case $1 in
      --package)
        MAKE_PACKAGE=true
        ;;
      *)
        echo "Invalid option: $1" >&2
        usage
        exit 1
        ;;
    esac
    shift
  done
}

# Main script execution
handle_options "$@"

if [ ! -d "./build" ]; then
  mkdir ./build
fi

platforms=("linux/386" "linux/amd64" "linux/arm64" "linux/riscv64" "darwin/arm64" "windows/arm64" "windows/amd64" "freebsd/amd64" "android/arm64")

for platform in "${platforms[@]}"; do
    echo "* Compiling for $platform..."
    platform_split=(${platform//\// })

    EXT=""
    if [ "${platform_split[0]}" == "windows" ]; then
      EXT=.exe
    fi

    if [ "$MAKE_PACKAGE" == "true" ]; then
        CGO_ENABLED=0 GOOS=${platform_split[0]} GOARCH=${platform_split[1]} go build -o build/hash_utils$EXT -a
        tar -czf build/hash_utils_${platform_split[0]}_${platform_split[1]}.tar.gz build/hash_utils$EXT
        rm build/hash_utils$EXT
    else
      CGO_ENABLED=0 GOOS=${platform_split[0]} GOARCH=${platform_split[1]} go build -o build/hash_utils_${platform_split[0]}_${platform_split[1]}$EXT -a
    fi
done
