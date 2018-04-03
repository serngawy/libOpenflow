#!/usr/bin/env bash

Update=$1

export GOPATH=$PWD/vendor
echo "GOPATH=" $GOPATH

if [ ! -d "vendor/src" ] || [ "$Update" = "update" ]; then
    mkdir -p bin
    mkdir -p vendor/src

    glide update -v
    # glide is wired, create the src dir and move dependencies under it.
    mkdir vendor/src
    for dir in vendor/*; do
      if [ "$dir" != "vendor/src" ]; then
         cp -r $dir vendor/src;
      fi;
    done
fi

# Copy the source code to the gopath dir in order to compile the new changes.
# For any reason of there is a new dir added, it should be copied also as below
cp -rf common/* vendor/src/github.com/serngawy/libOpenflow/common/
cp -rf ofctrl/* vendor/src/github.com/serngawy/libOpenflow/ofctrl/
cp -rf openflow13/* vendor/src/github.com/serngawy/libOpenflow/openflow13/
cp -rf protocol/* vendor/src/github.com/serngawy/libOpenflow/protocol/
cp -rf util/* vendor/src/github.com/serngawy/libOpenflow/util/
cp -rf libOpenflow.go vendor/src/github.com/serngawy/libOpenflow.go

go build -o bin/ofctrl

