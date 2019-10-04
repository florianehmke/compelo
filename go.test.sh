#!/usr/bin/env bash

go get github.com/wadey/gocovmerge

echo "" > coverage.txt
rm *.cov

# integration test coverage
go test -coverpkg=./... -coverprofile=integration-test.cov .

for d in $(go list ./... | grep -v vendor); do
    go test -coverprofile=profile.out $d
    if [ -f profile.out ]; then
        mv profile.out "$RANDOM".cov
    fi
done

gocovmerge *.cov > coverage.txt
