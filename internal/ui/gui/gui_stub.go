//go:build !fyne

package gui

import (
	"context"
	"fmt"

	"github.com/example/fmsx-frontend/internal/config"
	"github.com/example/fmsx-frontend/internal/db"
)

func Run(_ context.Context, _ config.SourcesConfig, _ *db.Store) error {
	return fmt.Errorf("GUI desativada neste build; recompile com -tags fyne")
}
