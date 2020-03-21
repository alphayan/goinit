# goinit    
goinit  create  a  single package go project  in  `$GOPATH/src/ or use go mod`   

## 2020-03-21 update
add go mod      
add Dockerfile  
add docker-compose.yml  
//TODO:postgresql 
## Start using it

Download and install it:    
`$ go get -u -v github.com/alphayan/goinit`     
Create a go project:    
`$ goinit -c projectname -f echo -o xorm -m`   

## Get help
Read help dockument:    
`$ goinit -h`

## Project directory hierarchy
``` 
<project>/       
    |- -main.go     
    |- -config.go       
    |- -db.go       
    |- -redis.go
    |- -router.go
    |- -.gitignore
    |- -config.toml
    |- -go.mod
    |- -Dockerfile
    |- -docker-compose.yml
```
## Tips
打包成可执行程序并压缩     
1.使用go build -ldflags '-w -s'进行代码编译，得到.exe文件    
2.使用upx小工具进行压缩，使得.exe文件大幅度缩小 [upx官网](https://upx.github.io/ "点击upx下载")
