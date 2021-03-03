package create

import (
	"embed"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
	"text/template"

	temp "github.com/alphayan/goinit/create/template"
)

// GOPATHSRC $GOPATH/src
var GOPATHSRC string

//go:embed template/*
var fs embed.FS

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
func Create(dir, frame, orm, db string, module bool, str *[]string) error {

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
	NewDockerfile(pth)
	NewDockerCompose(pth, db)
	for _, v := range *str {
		NewController(pth, v, frame)
		NewModel(pth, v, orm)
	}
	NewResponse(pth)
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
	data, err := fs.ReadFile("template/main.tmpl")
	if err != nil {
		return err
	}
	f.Write(data)
	return f.Sync()
}

// NewRedis create redis.go
func NewRedis(dir string) error {
	f, err := os.Create(path.Join(dir, "s_redis.go"))
	if err != nil {
		return err
	}
	defer f.Close()
	data, err := fs.ReadFile("template/redis.tmpl")
	if err != nil {
		return err
	}
	f.Write(data)
	return f.Sync()
}

// NewGitignore create .gitignore
func NewGitignore(dir string) error {
	f, err := os.Create(path.Join(dir, ".gitignore"))
	if err != nil {
		return err
	}
	defer f.Close()
	data, err := fs.ReadFile("template/gitignore.tmpl")
	if err != nil {
		return err
	}
	f.Write(data)
	return f.Sync()
}

// NewConfig create config.go
func NewConfig(dir string) error {
	f, err := os.Create(path.Join(dir, "s_config.go"))
	if err != nil {
		return err
	}
	defer f.Close()
	t, err := template.ParseFS(fs, "template/config.tmpl")
	if err != nil {
		return err
	}
	return t.Execute(f, dir)
}

// NewDB create db.go
func NewDB(dir, orm, db string) error {
	f, err := os.Create(path.Join(dir, "s_db.go"))
	if err != nil {
		return err
	}
	defer f.Close()
	switch db {
	default:
		switch orm {
		case "xorm":
			err := func() error {
				data, err := fs.ReadFile("template/xorm.tmpl")
				if err != nil {
					return err
				}
				f.Write(data)
				return nil
			}()
			if err != nil {
				return err
			}
		default:
			err := func() error {
				data, err := fs.ReadFile("template/gorm.tmpl")
				if err != nil {
					return err
				}
				f.Write(data)
				return nil
			}()
			if err != nil {
				return err
			}
		}
	}

	return f.Sync()
}

// NewRouter create router.go
func NewRouter(dir, frame string) error {
	f, err := os.Create(path.Join(dir, "s_router.go"))
	if err != nil {
		return err
	}
	defer f.Close()
	switch frame {
	case "echo":
		err := func() error {
			data, err := fs.ReadFile("template/echo.tmpl")
			if err != nil {
				return err
			}
			f.Write(data)
			return nil
		}()
		if err != nil {
			return err
		}
	case "gin":
		err := func() error {
			data, err := fs.ReadFile("template/echo.tmpl")
			if err != nil {
				return err
			}
			f.Write(data)
			return nil
		}()
		if err != nil {
			return err
		}
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
	data, err := fs.ReadFile("template/toml.tmpl")
	if err != nil {
		return err
	}
	f.Write(data)
	return f.Sync()
}

// NewDockerfile create Dockerfile
func NewDockerfile(dir string) error {
	f, err := os.Create(path.Join(dir, "Dockerfile"))
	if err != nil {
		return err
	}
	defer f.Close()
	t, err := template.ParseFS(fs, "template/dockerfile.tmpl")
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
	t, err := template.ParseFS(fs, "template/docker-compose.tmpl")
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

// NewController create controller
func NewController(dir, fn, frame string) error {
	f, err := os.Create(path.Join(dir, "c_"+strings.ToLower(fn)+".go"))
	if err != nil {
		return err
	}
	defer f.Close()
	var t *template.Template
	switch frame {
	case "gin":
		t, err = template.ParseFS(fs, "template/controller_gin.tmpl")
		if err != nil {
			return err
		}
	case "echo":
		t, err = template.ParseFS(fs, "template/controller_echo.tmpl")
		if err != nil {
			return err
		}
	}

	var com = struct {
		Name      string
		ShortName string
		LowName   string
	}{
		Name:      fn,
		ShortName: strings.ToLower(fn)[:1],
		LowName:   strings.ToLower(fn),
	}
	return t.Execute(f, com)
}

// NewModel create model
func NewModel(dir, fn, orm string) error {
	f, err := os.Create(path.Join(dir, "m_"+strings.ToLower(fn)+".go"))
	if err != nil {
		return err
	}
	defer f.Close()
	var t *template.Template
	switch orm {
	case "gorm":
		t, err = template.ParseFS(fs, "template/model_gorm.tmpl")
		if err != nil {
			return err
		}
	case "xorm":
		t, err = template.ParseFS(fs, "template/model_xorm.tmpl")
		if err != nil {
			return err
		}
	}
	var com = struct {
		Name      string
		ShortName string
		LowName   string
	}{
		Name:      fn,
		ShortName: strings.ToLower(fn)[:1],
		LowName:   strings.ToLower(fn),
	}
	return t.Execute(f, com)
}

// NewResponse create response.go
func NewResponse(dir string) error {
	f, err := os.Create(path.Join(dir, "s_response.go"))
	if err != nil {
		return err
	}
	defer f.Close()
	data, err := fs.ReadFile("template/response.tmpl")
	if err != nil {
		return err
	}
	f.Write(data)
	return f.Sync()
}
