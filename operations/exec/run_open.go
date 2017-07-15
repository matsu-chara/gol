package exec

import (
	"github.com/matsu-chara/gol/operations"
	"os/exec"
)

// RunOpen get entry with key, then open in browser
func RunOpen(filepath string, key string) error {
	entry, err := operations.RunGet(filepath, key)
	if err != nil {
		return err
	}

	if err := exec.Command("open", entry.Value).Run(); err != nil {
		return err
	}
	return nil
}
