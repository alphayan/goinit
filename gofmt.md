# go语言代码优化第三方库
### goimport 官方工具，比gofmt多了自动导入包名的功能  
安装：
     
`go get -u -v golang.org/x/tools/cmd/goimports`  

用法： 
    
`goimport -w -l .`
### gocyclo 用来检查函数的复杂度。   
安装  :

`go get -u -v github.com/fzipp/gocyclo
`   

用法
  
`gocyclo -over 12 $(ls -d */ | grep -v vendor)                  
`       

`gocyclo -top 10 $(ls -d */ | grep -v vendor)
`
### interfacer 检查是否可以定义为接口类型   
安装：     
       
`go get -u -v github.com/mvdan/interfacer/cmd/interfacer
`

用法：

`interfacer  
`
### deadcode  检查没有用到的代码   
安装：
  
`go get -u -v github.com/tsenart/deadcode
`

用法：

`deadcode`
### gotype 对go文件和包进行语义(semantic)和句法(syntactic)的分析,这是google提供的一个工具。
安装：

`go get -u -v golang.org/x/tools/cmd/gotype
`

用法：

`gotype *.go`
### misspell 用来拼写检查，对国内英语不太熟练的同学很有帮助。
安装：

`go get -u github.com/client9/misspell`

用法：

`misspell *.go`
### go-tools 强大的go工具   
安装：

`go get -u -v honnef.co/go/tools/...`


命令 | 用法  
-----|------  
gosimple    |	Detects code that could be rewritten in a simpler way. (检查哪些代码能被优化)
keyify	|Transforms an unkeyed struct literal into a keyed one.(将一个非键结构的结构转换为一个键控的)
rdeps	|Find all reverse dependencies of a set of packages.(查找一组包的所有反向依赖项)
staticcheck|	Detects a myriad of bugs and inefficiencies in your code.(检查大量bug或者低性能的代码)
structlayout|	Displays the layout (field sizes and padding) of structs.(显示结构体布局)
structlayout-optimize|	Reorders struct fields to minimize the amount of padding.(重新设计struct字段，以减少填充的数量)
structlayout-pretty	|Formats the output of structlayout with ASCII art.(用ASCII格式化输出)
unused|	Reports unused identifiers (types, functions, ...) in your code.(检查代码中未使用的标识)
megacheck|	Run staticcheck, gosimple and unused in one go(同时运行三个命令)


### goconst 会查找重复的字符串，这些字符串可以抽取成常量。
安装：

`go get -u -v github.com/jgautheron/goconst/cmd/goconst
`

用法：

`goconst ./...|grep -v vendor`