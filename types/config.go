package types

import "github.com/Creaft-JP/tit/types/config"

type Config struct {
	Remotes []config.Remote `json:"remotes"`
}
