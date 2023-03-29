package cmd

import (
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var (
	listen       uint16
	goDir        string
	cacheDir     string
	cacheTimeout time.Duration
)

func loadEnv() {
	listen = 8080
	if str, ok := os.LookupEnv("LISTEN"); ok {
		if port, err := strconv.ParseUint(str, 10, 16); err == nil {
			listen = uint16(port)
		}
	}

	goDir = os.Getenv("GO_DIR")
	if len(goDir) == 0 {
		goDir = filepath.Join(os.Getenv("GOROOT"), "bin", "go")
	}
	if len(goDir) == 0 {
		goDir = filepath.Join(os.Getenv("GOPATH"), "bin", "go")
	}
	if len(goDir) == 0 {
		goDir = "go"
	}

	cacheDir = os.Getenv("CACHE_DIR")
	if len(cacheDir) == 0 {
		cacheDir = os.TempDir()
	}

	cacheTimeout = 24 * time.Hour
	if str, ok := os.LookupEnv("CACHE_TIMEOUT"); ok {
		if t, err := time.ParseDuration(str); err == nil {
			cacheTimeout = t
		}
	}
}
