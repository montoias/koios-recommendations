#!/usr/bin/env bash

go fmt $(go list ./... | grep -v /*vendor*/)
find . -path ./vendor -prune -o -name '*.go' -print | xargs goimports -w -l
go test -cover $(go list ./... | grep -v /*vendor*/) | \
  grep -v "\[no test files\]"
find . -type d \( -path "./*vendor*" -o -path "./.*" -o -path "\.$" \) -prune -o -mindepth 1 -type d -print | \
       xargs -n 1 golint | \
       grep -v "receiver name should be a reflection of its identity" | \
       grep -v "don't use an underscore in package name" | \
       grep -v ", and that stutters; consider" | \
       grep -v "mocks"
find . -type d \( -path "./*vendor*" -o -path "./__vendor" -o -path "./.*" -o -path "\.$" \) \
       -prune -o -name "*.go" -print | \
       xargs -n 1 go vet
find . -type d \( -path "./*vendor*" -o -path "./__vendor" -o -path "./.*" -o -path "\.$" \) \
       -prune -o -name "*.go" -print | \
       xargs -n 1 go fix
