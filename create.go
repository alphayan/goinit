package main

import (
	"errors"
	"fmt"
	"os"
	"path"
)

// gocreate create a dir in $GOPATH/src/
func gocreate(dir, frame, db string) error {
	if isExist(path.Join(GOPATHSRC, dir)) {
		return errors.New("project is already exist,please change the projectname or remove the project")
	}
	err := os.MkdirAll(path.Join(GOPATHSRC, dir), 0777)
	if err != nil {
		return err
	}
	fmt.Println(newMain(dir, db), newConfig(dir), newDB(dir, db), newRedis(dir), newRouter(dir, frame), newGitignore(dir), newToml(dir))
	return nil
}

// newMain creat main.go
func newMain(dir, db string) error {
	f, err := os.Create(path.Join(GOPATHSRC, dir, "main.go"))
	if err != nil {
		return err
	}
	defer f.Close()
	switch db {
	case "xorm":
		f.WriteString(fmt.Sprintf(MAIN, "engine", "engine"))
	default:
		f.WriteString(fmt.Sprintf(MAIN, "db", "db"))
	}
	return nil
}

// newRedis create redis.go
func newRedis(dir string) error {
	f, err := os.Create(path.Join(GOPATHSRC, dir, "redis.go"))
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(REDIS)
	return nil
}

// newGitignore create .gitignore
func newGitignore(dir string) error {
	f, err := os.Create(path.Join(GOPATHSRC, dir, ".gitignore"))
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(GITIGNORE)
	return nil
}

// isExist check that the directory exists
func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// newConfig create config.go
func newConfig(dir string) error {
	f, err := os.Create(path.Join(GOPATHSRC, dir, "config.go"))
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(CONFIG)
	return nil
}

// newDB create db.go
func newDB(dir, db string) error {
	f, err := os.Create(path.Join(GOPATHSRC, dir, "db.go"))
	if err != nil {
		return err
	}
	defer f.Close()
	switch db {
	case "gorm":
		f.WriteString(GORM)
	case "xorm":
		f.WriteString(XORM)
	default:
		f.WriteString(DB)
	}
	return nil
}

// newRouter create router.go
func newRouter(dir, frame string) error {
	f, err := os.Create(path.Join(GOPATHSRC, dir, "router.go"))
	if err != nil {
		return err
	}
	defer f.Close()
	switch frame {
	case "echo":
		f.WriteString(ECHO)
	case "gin":
		f.WriteString(GIN)
	case "go-json-rest":
		f.WriteString(GOJSONREST)
	case "iris":
		f.WriteString(IRIS)
	default:
		f.WriteString(NETHTTP)
	}
	return nil
}

// newToml create config.toml
func newToml(dir string) error {
	f, err := os.Create(path.Join(GOPATHSRC, dir, "config.toml"))
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(TOML)
	return nil
}
