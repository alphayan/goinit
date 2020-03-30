package temp

// MAIN ...
const MAIN = `package main

func main() {
	initConfig()
	if conf.DB {
		initDB()
	}
	if conf.Redis {
		initRedis()
	}
	initRouter()
}
`
