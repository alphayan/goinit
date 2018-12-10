package main

const (
	// CONFIG ...
	CONFIG = `package main

import (
	"github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
	"github.com/sirupsen/logrus"
)

// Config 配置文件结构体
type Config struct {
	RedisHost     string //redis地址
	RedisPort     string //redis端口
	RedisPassword string //redis密码
	RedisDB       int    //redis数据库 0-15

	DBHost     string //DB地址
	DBPort     string //DB端口
	DBUsername string //DB用户
	DBPassword string //DB密码
	DBName     string //DB数据库
}

// TomlConfig toml的配置文件
type TomlConfig struct {
	Runmode Runmode
	Pro     Pro
	Dev     Dev
}

type Pro = Config
type Dev = Config

type Runmode struct {
	Runmode string
}

var conf = new(Config)

func initConfig() {
	tc := new(TomlConfig)
	if err := config.Load(file.NewSource(
		file.WithPath("config.toml"),
	)); err != nil {
		logrus.Fatal(err)
	}
	if err := config.Scan(&tc); err != nil {
		logrus.Fatal(err)
	}
	if tc.Runmode.Runmode == "dev" {
		*conf = tc.Dev
	} else {
		*conf = tc.Pro
	}
	tc = nil
	logrus.Infof("%+v", conf)
	logrus.Info("初始化config成功")
}
`
	// TOML ...
	TOML = `#运行模式
[Runmode]
Runmode="dev"
#production
[pro]
#redis配置
RedisHost="192.168.0.100"
RedisPort="6379"
RedisPassword=""
RedisDB=0
#DB配置
DBHost="192.168.0.100"
DBPort="3306"
DBUsername="root"
DBPassword="password"
DBName="test"

#develop
[dev]
#redis配置
RedisHost="192.168.199.248"
RedisPort="6379"
RedisPassword=""
RedisDB=0
#DB配置
DBHost="192.168.199.248"
DBPort="3306"
DBUsername="root"
DBPassword="password"
DBName="test"
`

	// GORM ...
	GORM = `package main

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var db *gorm.DB

func initDB() {
	path := conf.DBUsername + ":" + conf.DBPassword + "@tcp(" + conf.DBHost + ":" + conf.DBPort + ")/" + conf.DBName + "?charset=utf8"
	var err error
	for {
		db, err = gorm.Open("mysql", path)
		if err != nil {
			logrus.Error("DB connect error:", err, "Retry in 2 seconds!")
			time.Sleep(time.Second * 2)
			continue
		}
		logrus.Info("DB connect successful!")
		break
	}
	db.LogMode(true)
}
`
	// XORM ...
	XORM = `package main

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/sirupsen/logrus"
)

var engine *xorm.Engine

func initDB() {
	path := conf.DBUsername + ":" + conf.DBPassword + "@tcp(" + conf.DBHost + ":" + conf.DBPort + ")/" + conf.DBName + "?charset=utf8"
	var err error
	for {
		engine, err = xorm.NewEngine("mysql", path)
		if err != nil {
			logrus.Error("DB connect error:", err, "Retry in 2 seconds!")
			time.Sleep(time.Second * 2)
			continue
		}
		logrus.Info("DB connect successful!")
		break
	}
	engine.ShowSQL(true)
	engine.ShowExecTime(true)
}
`
	// ECHO ...
	ECHO = `package main

import (
	"github.com/labstack/echo"
)

func initRouter() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
`
	// GIN ...
	GIN = `package main

import (
	"github.com/gin-gonic/gin"
)

func initRouter() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
`
	// IRIS ...
	IRIS = `package main

import "github.com/kataras/iris"

func initRouter() {
    app := iris.Default()
    app.Get("/ping", func(ctx iris.Context) {
        ctx.JSON(iris.Map{
            "message": "pong",
        })
    })
    // listen and serve on http://0.0.0.0:8080.
    app.Run(iris.Addr(":8080"))
}`
	// GOJSONREST ...
	GOJSONREST = `package main

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

func initRouter() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	api.SetApp(rest.AppSimple(func(w rest.ResponseWriter, r *rest.Request) {
		w.WriteJson(map[string]string{"Body": "Hello World!"})
	}))
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}`
	// GITIGNORE ...
	GITIGNORE = `# Binaries for programs and plugins
*.exe
*.dll
*.so
*.dylib

# Test binary, build with 'go test -c'
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

#jetbrains
.idea/
`
	// MAIN
	MAIN = `package main

func main() {
	if %s != nil {
		defer %s.Close()
	}
	if redisClient != nil {
		defer redisClient.Close()
	}
	initConfig()
	initDB()
	initRedis()
	initRouter()
}
`
	// REDIS ...
	REDIS = `package main

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

var redisClient *redis.Client

func initRedis() {
	for {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     conf.RedisHost + ":" + conf.RedisPort,
			Password: conf.RedisPassword, // no password set
			DB:       conf.RedisDB,       // use default DB
		})
		_, err := redisClient.Ping().Result()
		if err != nil {
			logrus.Error(err, "Retry in 2 seconds!")
			time.Sleep(time.Second * 2)
			continue
		}
		logrus.Info("Redis connect successful!")
		break
	}
}
`
	// DB ...
	DB = `package main

import (
	_ "github.com/go-sql-driver/mysql"

	"database/sql"
	"time"

	"github.com/sirupsen/logrus"
)

var db *sql.DB

func initDB() {
	path := conf.DBUsername + ":" + conf.DBPassword + "@tcp(" + conf.DBHost + ":" + conf.DBPort + ")/" + conf.DBName + "?charset=utf8"
	var err error
	for {
		db, err = sql.Open("mysql", path)
		if err != nil {
			logrus.Error("DB connect error:", err, "Retry in 2 seconds!")
			time.Sleep(time.Second * 2)
			continue
		}
		logrus.Info("DB connect successful!")
		break
	}
}
`
	// NETHTTP
	NETHTTP = `package main

import "net/http"

func initRouter() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("hello world"))
	})
	http.ListenAndServe(":8080", nil)
}`
)
