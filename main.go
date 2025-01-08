package main

import (
	"fmt"
	"os"
	"time"

	"github.com/atotto/clipboard"
)

func copyUnixTume() error {
	unixTime := time.Now().Unix()
	unixTimeStr := fmt.Sprintf("%d", unixTime)

	err := clipboard.WriteAll(unixTimeStr)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := copyUnixTume()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
