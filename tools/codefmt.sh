#!/bin/sh

find ../src/socket -iname "*.go" | xargs gofmt -w -s -l;
find ../src/client -iname "*.go" | xargs gofmt -w -s -l;
