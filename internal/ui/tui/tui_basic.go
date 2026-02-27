//go:build !tview

package tui

import (
	"context"
	"fmt"

	"github.com/example/fmsx-frontend/internal/config"
	"github.com/example/fmsx-frontend/internal/db"
)

func Run(_ context.Context, sources config.SourcesConfig, store *db.Store) error {
	fmt.Println("fMSX Frontend (modo texto básico)")
	fmt.Printf("File Hunter: %s\n", sources.FileHunter.BaseURL)
	fmt.Printf("Game DB: %s\n", sources.GameDB.ArchiveURL)
	fmt.Printf("SQLite path: %s\n", store.Path)
	if !store.DriverReady {
		fmt.Println("SQLite driver ainda não configurado no build atual.")
	}
	return nil
}
