package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/example/fmsx-frontend/internal/app"
)

func main() {
	mode := flag.String("mode", "tui", "UI mode: tui or gui")
	configPath := flag.String("config", "configs/sources.json", "Path to JSON sources config")
	dbPath := flag.String("db", "data/fmsx_frontend.db", "Path to sqlite database")
	flag.Parse()

	application, err := app.New(*configPath, *dbPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "initialization error: %v\n", err)
		os.Exit(1)
	}

	if err := application.Run(context.Background(), *mode); err != nil {
		fmt.Fprintf(os.Stderr, "runtime error: %v\n", err)
		os.Exit(1)
	}
}
