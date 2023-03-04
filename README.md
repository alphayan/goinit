# goinit    
goinit  create  a  single package go project  in  `$GOPATH/src/ or use go mod`   

## 2023-03-04 update
update to go1.20.1

## 2021-03-03 update
update to go1.16

## 2020-03-21 update
add go mod      
add Dockerfile  
add docker-compose.yml  
//TODO:postgresql 
## Start using it

Download and install it:    
`$ go get -u -v github.com/alphayan/goinit`     
Create a go project:    
`$ goinit -c projectname -f echo -o xorm -s User,Log`   

## Get help
Read help dockument:    
`$ goinit -h`

## Project directory hierarchy
``` 
<project>/ 
    |- -c_controller.go
    |- -m_model.go 
    |- -main.go 
    |- -s_config.go       
    |- -s_db.go       
    |- -s_redis.go
    |- -s_router.go
    |- -s_response.go
    |- -.gitignore
    |- -config.toml
    |- -go.mod
    |- -Dockerfile
    |- -docker-compose.yml
```
## Tips
打包成可执行程序并压缩     
1.使用go build -ldflags '-w -s -extldflags "-static"'进行代码编译，得到.exe文件    
2.使用upx小工具进行压缩，使得.exe文件大幅度缩小 [upx官网](https://upx.github.io/ "点击upx下载")
