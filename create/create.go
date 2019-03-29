package create

import (
	"errors"
	"fmt"
	"os"
	"path"
)

// GOPATHSRC gopath/src路径
var GOPATHSRC string

func init() {
	GOPATHSRC = path.Join(os.Getenv("GOPATH"), "/src/")
}

// isExist check that the directory exists
func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// Create create a dir in $GOPATH/src/
func Create(dir, frame, db string) error {
	if isExist(path.Join(GOPATHSRC, dir)) {
		return errors.New("this project is already exist,please change the projectname or remove the project")
	}
	err := os.MkdirAll(path.Join(GOPATHSRC, dir), 0777)
	if err != nil {
		return err
	}
	NewMain(dir, db)
	NewConfig(dir)
	NewDB(dir, db)
	NewRedis(dir)
	NewRouter(dir, frame)
	NewGitignore(dir)
	NewToml(dir)
	return nil
}

// NewMain creat main.go
func NewMain(dir, db string) error {
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

// NewRedis create redis.go
func NewRedis(dir string) error {
	f, err := os.Create(path.Join(GOPATHSRC, dir, "redis.go"))
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(REDIS)
	return nil
}

// NewGitignore create .gitignore
func NewGitignore(dir string) error {
	f, err := os.Create(path.Join(GOPATHSRC, dir, ".gitignore"))
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(GITIGNORE)
	return nil
}

// NewConfig create config.go
func NewConfig(dir string) error {
	f, err := os.Create(path.Join(GOPATHSRC, dir, "config.go"))
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(CONFIG)
	return nil
}

// NewDB create db.go
func NewDB(dir, db string) error {
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

// NewRouter create router.go
func NewRouter(dir, frame string) error {
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

// NewToml create config.toml
func NewToml(dir string) error {
	f, err := os.Create(path.Join(GOPATHSRC, dir, "config.toml"))
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(TOML)
	return nil
}
