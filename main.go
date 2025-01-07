package main

import (
	"fmt"
	"os"
	"time"

	"golang.design/x/clipboard"
)

func copyUnixTume() error {
	err := clipboard.Init()
	if err != nil {
		return fmt.Errorf("Failed to initialize the clipboard: %w", err)
	}

	unixTime := time.Now().Unix()
	unixTimeStr := fmt.Sprintf("%d", unixTime)

	clipboard.Write(clipboard.FmtText, []byte(unixTimeStr))

	return nil
}

func main() {
	err := copyUnixTume()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
