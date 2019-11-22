package create

import (
	"errors"
	"os"
	"path"
	"text/template"

	temp "github.com/alphayan/goinit/template"
)

// GOPATHSRC $GOPATH/src
var GOPATHSRC string

func init() {
	if os.Getenv("GOPATH") == "" {
		GOPATHSRC, _ = os.Getwd()
		return
	}
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
	f.WriteString(temp.MAIN)
	return f.Sync()
}

// NewRedis create redis.go
func NewRedis(dir string) error {
	f, err := os.Create(path.Join(GOPATHSRC, dir, "redis.go"))
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(temp.REDIS)
	return f.Sync()
}

// NewGitignore create .gitignore
func NewGitignore(dir string) error {
	f, err := os.Create(path.Join(GOPATHSRC, dir, ".gitignore"))
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(temp.GITIGNORE)
	return f.Sync()
}

// NewConfig create config.go
func NewConfig(dir string) error {
	f, err := os.Create(path.Join(GOPATHSRC, dir, "config.go"))
	if err != nil {
		return err
	}
	defer f.Close()
	t, err := template.New("config").Parse(temp.CONFIG)
	if err != nil {
		return err
	}
	return t.Execute(f, dir)

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
		f.WriteString(temp.GORM)
	case "xorm":
		f.WriteString(temp.XORM)
	default:
		f.WriteString(temp.DB)
	}
	return f.Sync()
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
		f.WriteString(temp.ECHO)
	case "gin":
		f.WriteString(temp.GIN)
	case "iris":
		f.WriteString(temp.IRIS)
	default:
		f.WriteString(temp.NETHTTP)
	}
	return f.Sync()
}

// NewToml create config.toml
func NewToml(dir string) error {
	f, err := os.Create(path.Join(GOPATHSRC, dir, "config.toml"))
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(temp.TOML)
	return f.Sync()
}
