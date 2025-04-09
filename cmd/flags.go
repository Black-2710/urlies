// cmd/flags.go
package cmd

import (
	"flag"
)

var (
	TargetURL string
	OutputDir string
	Mode      string
	Depth     int
)

func InitFlags() {
	flag.StringVar(&TargetURL, "u", "", "Target URL (required)")
	flag.StringVar(&Mode, "m", "all", "Recon mode: archive, engine, headless, or all")
	flag.StringVar(&OutputDir, "o", "output", "Output directory")
	flag.IntVar(&Depth, "depth", 20, "Number of archive pages to fetch (each page ~1000 URLs)")
	flag.Parse()
}
