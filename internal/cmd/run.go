package cmd

import (
	"fmt"
	"goproxy/internal/cache"
	"net/http"

	"github.com/goproxy/goproxy"
	"github.com/lwch/runtime"
)

func Run() {
	loadEnv()

	proxy := &goproxy.Goproxy{
		GoBinName:     goDir,
		ProxiedSUMDBs: []string{"sum.golang.org", "sum.golang.google.cn"},
		Cacher:        cache.New(cacheDir, cacheTimeout),
	}
	err := http.ListenAndServe(fmt.Sprintf(":%d", listen), proxy)
	runtime.Assert(err)
}
