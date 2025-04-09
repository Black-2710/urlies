// main.go
package main

import (
	"fmt"
	"urlies/cmd"
	"urlies/modules/archive"
	"urlies/modules/engine"
	"urlies/modules/headless"
)

func main() {
	cmd.InitFlags()

	fmt.Println("URLies v0.1")
	fmt.Println("Target:", cmd.TargetURL)
	fmt.Println("Mode:", cmd.Mode)
	fmt.Println("Output Directory:", cmd.OutputDir)

	switch cmd.Mode {
		case "archive":
			archive.Run(cmd.TargetURL, cmd.OutputDir, cmd.Depth)
		case "engine":
			engine.Run(cmd.TargetURL, cmd.OutputDir)
		case "headless":
			headless.Run(cmd.TargetURL, cmd.OutputDir)
		case "all":
			archive.Run(cmd.TargetURL, cmd.OutputDir, cmd.Depth)
			engine.Run(cmd.TargetURL, cmd.OutputDir)
			headless.Run(cmd.TargetURL, cmd.OutputDir)
		default:
			fmt.Println("[!] Invalid mode. Use: archive, engine, headless, or all")
	}
}
