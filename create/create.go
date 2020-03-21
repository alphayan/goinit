package create

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
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
func Create(dir, frame, orm, db string, module bool) error {
	var pth = dir
	if module {
		pth = dir
	} else {
		pth = path.Join(GOPATHSRC, dir)
	}
	fmt.Println("pwd:", pth)
	if isExist(pth) {
		return errors.New("this project is already exist,please change the projectname or remove the project")
	}
	err := os.MkdirAll(pth, 0777)
	if err != nil {
		return err
	}
	NewMain(pth)
	NewConfig(pth)
	NewDB(pth, orm, db)
	NewRedis(pth)
	NewRouter(pth, frame)
	NewGitignore(pth)
	NewToml(pth)
	NewDockerfile(dir)
	NewDockerCompose(dir, db)
	if module {
		os.Chdir(pth)
		os.Setenv("GO111MODULE", "on")
		os.Setenv("GOPROXY", "https://goproxy.cn,https://goproxy.io,direct")
		cmd := exec.Command("/bin/bash", "-c", "go mod init "+dir)
		fmt.Println("go mod init ", dir)
		cmd.Run()
		cmd2 := exec.Command("/bin/bash", "-c", "go mod tidy")
		fmt.Println("go mod tidy")
		cmd2.Run()
		cmd3 := exec.Command("/bin/bash", "-c", "go fmt ./...")
		fmt.Println("go fmt ./...")
		cmd3.Run()
		cmd4 := exec.Command("/bin/bash", "-c", "go mod vendor")
		fmt.Println("go mod vendor")
		cmd4.Run()
	}
	return nil
}

// NewMain creat main.go
func NewMain(dir string) error {
	f, err := os.Create(path.Join(dir, "main.go"))
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(temp.MAIN)
	return f.Sync()
}

// NewRedis create redis.go
func NewRedis(dir string) error {
	f, err := os.Create(path.Join(dir, "redis.go"))
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(temp.REDIS)
	return f.Sync()
}

// NewGitignore create .gitignore
func NewGitignore(dir string) error {
	f, err := os.Create(path.Join(dir, ".gitignore"))
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(temp.GITIGNORE)
	return f.Sync()
}

// NewConfig create config.go
func NewConfig(dir string) error {
	f, err := os.Create(path.Join(dir, "config.go"))
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
func NewDB(dir, orm, db string) error {
	f, err := os.Create(path.Join(dir, "db.go"))
	if err != nil {
		return err
	}
	defer f.Close()
	switch db {
	case "postgresql":
		switch orm {
		case "gorm":
			f.WriteString(temp.GORM_POSTGRESQL)
		case "xorm":
			f.WriteString(temp.XORM_POSTGRESQL)
		default:
			f.WriteString(temp.DB_POSTGRESQL)
		}
	case "mongodb":
	default:
		switch orm {
		case "gorm":
			f.WriteString(temp.GORM_MYSQL)
		case "xorm":
			f.WriteString(temp.XORM_MYSQL)
		default:
			f.WriteString(temp.DB_MYSQL)
		}
	}

	return f.Sync()
}

// NewRouter create router.go
func NewRouter(dir, frame string) error {
	f, err := os.Create(path.Join(dir, "router.go"))
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
	f, err := os.Create(path.Join(dir, "config.toml"))
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(temp.TOML)
	return f.Sync()
}

// NewDockerfile create Dockerfile
func NewDockerfile(dir string) error {
	f, err := os.Create(path.Join(dir, "Dockerfile"))
	if err != nil {
		return err
	}
	defer f.Close()
	t, err := template.New("Dockerfile").Parse(temp.DOCKERFILE)
	if err != nil {
		return err
	}
	return t.Execute(f, dir)
}

// NewDockerCompose create docker-compose.yml
func NewDockerCompose(dir, db string) error {
	f, err := os.Create(path.Join(dir, "docker-compose.yml"))
	if err != nil {
		return err
	}
	defer f.Close()
	t, err := template.New("docker-compose.yml").Parse(temp.DOCKER_COMPOSE)
	if err != nil {
		return err
	}
	var com = struct {
		APP string
		DB  string
	}{
		APP: dir,
		DB:  db,
	}
	return t.Execute(f, com)
}
