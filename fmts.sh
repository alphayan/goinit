#!/bin/bash
cd $GOPATH/src/$1
#echo n|glide init
#export http_proxy=http://192.168.31.99:8123
#export https_proxy=$http_proxy
#export HTTP_PROXY=$http_proxy
#export HTTPS_PROXY=$HTTP_PROXY
#glide up
goimports -w -l .|grep -v vendor
interfacer
misspell .|grep -v vendor
megacheck
goconst ./...|grep -v vendor
go vet
golint .|grep -v vendor
echo 格式化完成


