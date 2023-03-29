package cmd

import (
	"flag"
	"fmt"
	"goproxy/internal/storage"
	"goproxy/internal/utils"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/goproxy/goproxy"
	"github.com/lwch/logging"
	"github.com/lwch/runtime"
)

func Run() {
	conf := flag.String("conf", "", "config file")
	flag.Parse()

	if len(*conf) == 0 {
		fmt.Println("missing -conf flag")
		os.Exit(1)
	}

	cfg := loadConf(*conf)

	showGoVersion(cfg.GoDir)
	go clearStorage(cfg.Storage)

	proxy := &goproxy.Goproxy{
		GoBinName:     cfg.GoDir,
		ProxiedSUMDBs: []string{"sum.golang.org", "sum.golang.google.cn"},
		Cacher:        cfg.Storage,
	}
	logging.Info("goproxy listen on %d", cfg.Listen)
	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Listen), proxy)
	runtime.Assert(err)
}

func showGoVersion(dir string) {
	cmd := exec.Command(dir, "version")
	cmd.Stdout = &logging.DefaultLogger
	err := cmd.Run()
	runtime.Assert(err)
}

func clearStorage(storage storage.Storage) {
	tk := time.NewTicker(time.Minute)
	for {
		<-tk.C
		func() {
			defer utils.Recover("clear storage")
			storage.Clear()
		}()
	}
}
