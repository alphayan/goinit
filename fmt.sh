#!/bin/bash
cd $GOPATH/src/$1
go fmt
goimports -w -l .
echo 格式化完成