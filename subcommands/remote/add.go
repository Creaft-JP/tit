package remote

import (
	"encoding/json"
	"fmt"
	"github.com/Creaft-JP/tit/types"
	"github.com/Creaft-JP/tit/types/config"
	"io"
)

func Add(args []string, reader io.Reader, writer io.Writer) error {
	name := args[0]

	if len(args) != 2 {
		return fmt.Errorf("args should be 2, but received %d", len(args))
	}

	decoder := json.NewDecoder(reader)
	var configContent types.Config
	if err := decoder.Decode(&configContent); err != nil {
		return err
	}
	encoder := json.NewEncoder(writer)

	for _, remote := range configContent.Remotes {
		if remote.Name == name {
			return fmt.Errorf("remote %s already exists", name)
		}
	}

	configContent.Remotes = append(configContent.Remotes, config.Remote{Name: name, Url: args[1]})
	if err := encoder.Encode(configContent); err != nil {
		return err
	}
	return nil
}
