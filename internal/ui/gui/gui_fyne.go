//go:build fyne

package gui

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/example/fmsx-frontend/internal/config"
	"github.com/example/fmsx-frontend/internal/db"
)

func Run(_ context.Context, sources config.SourcesConfig, _ *db.Store) error {
	a := app.NewWithID("org.msx.fmsx.frontend")
	w := a.NewWindow("fMSX Frontend (Fyne)")

	content := container.NewVBox(
		widget.NewLabel("Frontend inicial para fMSX"),
		widget.NewLabel(fmt.Sprintf("File Hunter: %s", sources.FileHunter.BaseURL)),
		widget.NewLabel(fmt.Sprintf("Game DB: %s", sources.GameDB.ArchiveURL)),
	)
	w.SetContent(content)
	w.Resize(fyne.NewSize(680, 320))
	w.ShowAndRun()
	return nil
}
