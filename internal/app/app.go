package app

import (
	"context"
	"fmt"

	"github.com/example/fmsx-frontend/internal/config"
	"github.com/example/fmsx-frontend/internal/db"
	"github.com/example/fmsx-frontend/internal/ui/gui"
	"github.com/example/fmsx-frontend/internal/ui/tui"
)

type App struct {
	Sources config.SourcesConfig
	Store   *db.Store
}

func New(configPath, dbPath string) (*App, error) {
	sources, err := config.LoadSources(configPath)
	if err != nil {
		return nil, fmt.Errorf("load source config: %w", err)
	}

	store, err := db.NewStore(dbPath)
	if err != nil {
		return nil, fmt.Errorf("open sqlite store: %w", err)
	}

	if err := store.EnsureSchema(); err != nil {
		return nil, fmt.Errorf("ensure schema: %w", err)
	}

	return &App{Sources: sources, Store: store}, nil
}

func (a *App) Run(ctx context.Context, mode string) error {
	switch mode {
	case "tui":
		return tui.Run(ctx, a.Sources, a.Store)
	case "gui":
		return gui.Run(ctx, a.Sources, a.Store)
	default:
		return fmt.Errorf("unsupported mode %q", mode)
	}
}
