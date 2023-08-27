package remote

import (
	"encoding/json"
	e "github.com/Creaft-JP/tit/error"
	"github.com/Creaft-JP/tit/types"
	"github.com/Creaft-JP/tit/types/config"
	"github.com/morikuni/failure"
	"io"
)

func Add(args []string, configReader io.Reader, configWriter io.Writer) error {
	name := args[0]

	if len(args) != 2 {
		return failure.New(e.Operation, failure.Messagef("args should be 2, but received %d", len(args)))
	}

	decoder := json.NewDecoder(configReader)
	var configContent types.Config
	if err := decoder.Decode(&configContent); err != nil {
		return failure.Translate(err, e.File)
	}
	encoder := json.NewEncoder(configWriter)

	for _, remote := range configContent.Remotes {
		if remote.Name == name {
			return failure.New(e.Operation, failure.Messagef("remote %s already exists", name))
		}
	}

	configContent.Remotes = append(configContent.Remotes, config.Remote{Name: name, Url: args[1]})
	if err := encoder.Encode(configContent); err != nil {
		return failure.Translate(err, e.File)
	}
	return nil
}
