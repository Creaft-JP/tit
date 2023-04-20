package types

import (
	"github.com/Creaft-JP/tit/types/config"
	"path/filepath"
)

type Config struct {
	Remotes []config.Remote `json:"remotes"`
}

const ConfigFilename = "config.json"

var ConfigFilepath = filepath.Join(RepositoryDirectoryName, ConfigFilename)
