package basic

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
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

func IsSecure(c *gin.Context) bool {
	// 1️⃣ 直连 HTTPS
	if c.Request.TLS != nil {
		return true
	}

	// 2️⃣ 反向代理（Nginx / CDN）
	if proto := c.GetHeader("X-Forwarded-Proto"); proto == "https" {
		return true
	}

	// 3️⃣ RFC 标准 Forwarded
	if fwd := c.GetHeader("Forwarded"); strings.Contains(fwd, "proto=https") {
		return true
	}

	return false
}
