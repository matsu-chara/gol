package exec

import (
	"fmt"
	"github.com/matsu-chara/gol/kvs"
	"github.com/matsu-chara/gol/operations"
	"io"
	"os/exec"
)

// RunPeco filter by prefix and select in peco, then open in browser
func RunPeco(filepath string, prefix string) error {
	entries, err := operations.RunLs(filepath)
	if err != nil {
		return err
	}

	pecoTarget := kvs.Entries(entries).FilterByPrefix(prefix)

	c1 := exec.Command("peco")
	stdin, err := c1.StdinPipe()
	if err != nil {
		return err
	}
	defer func() {
		stdin.Close()
	}()

	go func() {
		for _, t := range pecoTarget {
			io.WriteString(stdin, fmt.Sprintf("%s\n", t.ToPeco()))
		}
		stdin.Close()
	}()

	pecoedBinary, err := c1.Output()
	if err != nil {
		return err
	}
	pecoedStr := string(pecoedBinary)
	if pecoedStr == "" {
		return nil
	}
	pecoed, err := kvs.EntryFromPeco(pecoedStr)
	if err != nil {
		return err
	}

	if err := exec.Command("open", pecoed.Value).Run(); err != nil {
		return err
	}

	return nil
}
