package basic

import (
	"log"
	"os"
	"path/filepath"
)

var rootPath string

func Init() {
	var err error
	//初始化根目录
	rootPath, err = initRootPath()
	if err != nil {
		log.Fatal(err)
	}

}

func initRootPath() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", os.ErrNotExist
		}
		dir = parent
	}
}

func GetRootPath() string {
	return rootPath
}
