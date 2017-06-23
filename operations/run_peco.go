package operations

import (
	"fmt"
	"io"
	"os/exec"
	"strings"

	"github.com/matsu-chara/gol/kvs"
)

// RunPeco filter by prefix and select in peco, then open in browser
func RunPeco(filepath string, prefix string) error {
	entries, err := RunLs(filepath)
	if err != nil {
		return err
	}

	pecoTarget := filterByPrefix(entries, prefix)

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

func filterByPrefix(entries []kvs.Entry, prefix string) []kvs.Entry {
	filtered := make([]kvs.Entry, 0)
	for _, entry := range entries {
		if strings.HasPrefix(entry.Key, prefix) {
			filtered = append(filtered, entry)
		}
	}
	return filtered
}
