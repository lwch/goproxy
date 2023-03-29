package cmd

import (
	"goproxy/internal/index"
	"goproxy/internal/storage"
	"goproxy/internal/storage/local"
	"os"
	"time"

	"github.com/dgraph-io/badger/v4"
	"github.com/lwch/runtime"
	"gopkg.in/yaml.v3"
)

type Configure struct {
	Listen  uint16
	GoDir   string
	Storage storage.Storage
}

func loadConf(dir string) *Configure {
	f, err := os.Open(dir)
	runtime.Assert(err)
	defer f.Close()
	var cfg struct {
		Listen uint16 `yaml:"listen"`
		Go     string `yaml:"go"`
		Cache  struct {
			Index   string        `yaml:"index"`
			Timeout time.Duration `yaml:"timeout"`
			Storage string        `yaml:"storage"`
			Local   struct {
				Dir string `yaml:"dir"`
			} `yaml:"local"`
		} `yaml:"cache"`
	}
	err = yaml.NewDecoder(f).Decode(&cfg)
	runtime.Assert(err)

	idx, err := badger.Open(badger.DefaultOptions(cfg.Cache.Index))
	runtime.Assert(err)
	indexer := index.New(idx)

	var storage storage.Storage
	switch cfg.Cache.Storage {
	case "local":
		storage = local.New(indexer, cfg.Cache.Local.Dir, cfg.Cache.Timeout)
	default:
		panic("unsupported storage: " + cfg.Cache.Storage)
	}
	return &Configure{
		Listen:  cfg.Listen,
		GoDir:   cfg.Go,
		Storage: storage,
	}
}
