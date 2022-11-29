#!/bin/sh
version="v1.0"
path="main"
flags="-X '$path.Version=$version' -X '$path.GoVersion=$(go version)' -X '$path.BuildTime=`date +"%Y-%m-%d %H:%m:%S"`'"
go build -ldflags "$flags" -o httpService
