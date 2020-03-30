package temp

// MAIN ...
const MAIN = `package main

func main() {
	if connectionDB != nil {
		defer func() {
			connectionDB.Close()
		}()
	}
	if connectionRedis != nil {
		defer func() {
			connectionRedis.Close()
		}()
	}
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
