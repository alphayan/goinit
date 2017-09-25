#!/bin/bash
cd $GOPATH/src/$1
goimports -w -l .|grep -v vendor
echo n|glide init
export http_proxy=http://192.168.31.99:8123
export https_proxy=$http_proxy
export HTTP_PROXY=$http_proxy
export HTTPS_PROXY=$HTTP_PROXY
glide up
echo 格式化完成