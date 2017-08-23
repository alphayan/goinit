package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
)

var echoRouter = `package main
	func initRouter(){
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
	}
`
var ginRouter = `package main
	func initRouter(){
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
	}
`

var goJSONRestRouter = `package main
	func initRouter(){
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	api.SetApp(rest.AppSimple(func(w rest.ResponseWriter, r *rest.Request) {
		w.WriteJson(map[string]string{"Body": "Hello World!"})
	}))
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
	}
`

// gocreate create a dir in $GOPATH/src/
func gocreate(dir, frame string) error {
	if isExist(path.Join(GOPATHSRC, dir)) {
		return errors.New("project is already exist,please change the projectname or remove the project")
	}
	err := os.MkdirAll(path.Join(GOPATHSRC, dir), 0777)
	if err != nil {
		return err
	}
	fmt.Println(newMain(dir), newConfig(dir), newDB(dir), newRabbitmq(dir), newRedis(dir), newRouter(dir, frame), newGitgnore(dir))
	return nil
}

func newMain(dir string) error {
	f, err := os.Create(path.Join(GOPATHSRC, dir, "main.go"))
	if err != nil {
		return err
	}
	f.WriteString(`package main

	func main(){

	}
	`)
	return nil
}
func newConfig(dir string) error {
	f, err := os.Create(path.Join(GOPATHSRC, dir, "config.go"))
	if err != nil {
		return err
	}
	f.WriteString(`package main
    // Config struct
	type Config struct {
	DBHost     string
	DBPort     string
	DBUsername string
	DBPassword string
	DBName     string

	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDb       int

	MQHost     string
	MQUsername string
	MQPassword string
	MQPort     string
}

var config Config

// initConfig 初始化配置信息
func initConfig() {
	// load .env file for develop env
	err := godotenv.Overload()
	if err != nil {
		log.Infof("Err loading .env file: %+v", err)
	}
	// load config
	m := multiconfig.MultiLoader(
		&multiconfig.TagLoader{},
		&multiconfig.EnvironmentLoader{
			// ErrorMap the verification lib errors
			CamelCase: true,
		},
	)
	m.Load(&config)
	// log settings
	if config.Debug {
		log.SetFormatter(&log.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "06-01-02 15:04:05.00",
		})
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetFormatter(&log.JSONFormatter{
			TimestampFormat: "06-01-02 15:04:05.00",
		})
		log.SetLevel(log.InfoLevel)
	}
}
	`)
	return nil
}
func newRouter(dir, frame string) error {
	f, err := os.Create(path.Join(GOPATHSRC, dir, "router.go"))
	if err != nil {
		return err
	}
	switch frame {
	case "echo":
		f.WriteString(echoRouter)
	case "gin":
		f.WriteString(ginRouter)
	case "go-json-rest":
		f.WriteString(goJSONRestRouter)
	default:
		f.WriteString(`package main
	func initRouter(){

	}
	`)
	}
	return nil
}
func newDB(dir string) error {
	f, err := os.Create(path.Join(GOPATHSRC, dir, "db.go"))
	if err != nil {
		return err
	}
	f.WriteString(`package main

	func initDB(){
	path := config.DBUsername + ":" + config.DBPassword + "@tcp(" + config.DBHost + ":" + config.DBPort + ")/" + config.DBName + "?charset=utf8"
	var err error
	for {
		db, err = gorm.Open("mysql", path)
		if err != nil {
			log.Error("DB connect error:", err, "Retry in 2 seconds!")
			time.Sleep(time.Second * 2)
			continue
		}
		log.Info("DB connect successful!")
		break
	}
	if config.Debug {
		db.LogMode(true)
	}
	}
	`)
	return nil
}
func newRedis(dir string) error {
	f, err := os.Create(path.Join(GOPATHSRC, dir, "redis.go"))
	if err != nil {
		return err
	}
	f.WriteString(`package main

	func initRedis(){
	for {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     config.RedisHost + ":" + config.RedisPort,
			Password: config.RedisPassword, // no password set
			DB:       config.RedisDb,       // use default DB
		})
		_, err := redisClient.Ping().Result()
		if err != nil {
			log.Error(err, "Retry in 2 seconds!")
			time.Sleep(time.Second * 2)
			continue
		}
		log.Info("Redis connect successful!")
		break
	}
	}
	`)
	return nil
}
func newRabbitmq(dir string) error {
	f, err := os.Create(path.Join(GOPATHSRC, dir, "rabbitmq.go"))
	if err != nil {
		return err
	}
	f.WriteString(`package main

	func initRabbitmq(){
	}
	`)
	return nil
}

func newGitgnore(dir string) error {
	f, err := os.Create(path.Join(GOPATHSRC, dir, ".gitignore"))
	if err != nil {
		return err
	}
	f.WriteString(`
	# Binaries for programs and plugins
	*.exe
	*.dll
	*.so
	*.dylib

	# Test binary, build with 'go test -c'
	*.test

	# Output of the go coverage tool, specifically when used with LiteIDE
	*.out

	# Project-local glide cache, RE: https://github.com/Masterminds/glide/issues/736
	.glide/

	#jetbrains
	.idea/

	#.init
	.init/

	#vendor
	.vendor/

	#gitignore
	.gitignore

	`)
	return nil
}

func goinitfmt(dir string) {
	mycmd := "/d/gopath/src/goinit/fmt.sh" + " " + dir
	f := exec.Command("sh", "-c", mycmd)
	stdout, err := f.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return
	}
	f.Start()
	reader := bufio.NewReader(stdout)
	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Print(line)
	}
	f.Wait()

}

// isExist check that the directory exists
func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
