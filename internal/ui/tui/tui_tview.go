//go:build tview

package tui

import (
	"context"
	"fmt"

	"github.com/rivo/tview"

	"github.com/example/fmsx-frontend/internal/config"
	"github.com/example/fmsx-frontend/internal/db"
)

func Run(_ context.Context, sources config.SourcesConfig, store *db.Store) error {
	app := tview.NewApplication()
	status := "ativo"
	if !store.DriverReady {
		status = "pendente (driver não embutido)"
	}

	summary := tview.NewTextView().
		SetDynamicColors(true).
		SetText(fmt.Sprintf("[yellow]fMSX Frontend (TUI/tview)\n\n[white]File Hunter: %s\nGame DB: %s\nSQLite: %s\nDriver: %s",
			sources.FileHunter.BaseURL,
			sources.GameDB.ArchiveURL,
			store.Path,
			status,
		))

	summary.SetBorder(true).SetTitle("Status")
	if err := app.SetRoot(summary, true).SetFocus(summary).Run(); err != nil {
		return fmt.Errorf("run tview: %w", err)
	}
	return nil
}
