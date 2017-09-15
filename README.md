# goinit    


goinit  create  a   go  project  in  `$GOPATH/src/`   


## Start using it

Download and install it:    
`$ go get -u -v github.com/alphayan/goinit`     
Create a go project:    
`$ goinit -c projectname`   

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
    |- -rabbitmq.go     
    |- -router.go
    |- -.gitignore
```
## Tips
打包成可执行程序
1.使用go build -ldflags '-w -s'进行代码编译，得到.exe文件
2.使用upx小工具进行压缩，使得.exe文件大幅度缩小 [upx官网](https://upx.github.io/ "点击upx下载")
