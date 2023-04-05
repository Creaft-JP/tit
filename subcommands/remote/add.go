package remote

import (
	"encoding/json"
	"fmt"
	"github.com/Creaft-JP/tit/types"
	"github.com/Creaft-JP/tit/types/config"
	"io"
)

func Add(args []string, reader io.Reader, writer io.Writer) error {
	if len(args) != 2 {
		return fmt.Errorf("args should be 2, but received %d", len(args))
	}

	encoder := json.NewEncoder(writer)
	configContent := types.Config{Remotes: []config.Remote{{Name: args[0], Url: args[1]}}}
	if err := encoder.Encode(configContent); err != nil {
		return err
	}
	return nil
}
