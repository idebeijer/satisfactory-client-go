//go:build example

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/idebeijer/satisfactory-client-go/satisfactory"
)

// This example demonstrates how to download a save game file from a satisfactory server.
// The save game file is saved to disk with the name "savegame_autosave_0.sav".
func main() {
	ctx := context.Background()
	client := satisfactory.NewClient("https://localhost:7777", nil, true)

	password := os.Getenv("SF_PASSWD") // Replace with your password
	if _, err := client.PasswordLogin(ctx, "Administrator", password); err != nil {
		fmt.Println(err)
		return
	}

	saveName := "savegame_autosave_0"
	saveFile, _, err := client.DownloadSaveGame(ctx, saveName)
	if err != nil {
		fmt.Println(err)
		return
	}

	f, err := os.Create(saveName + ".sav")
	if err != nil {
		fmt.Println(err)
		return
	}
	f.Write(saveFile)
}
