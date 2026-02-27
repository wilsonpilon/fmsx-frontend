package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
)

type Store struct {
	Path        string
	DB          *sql.DB
	DriverReady bool
}

func NewStore(path string) (*Store, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, fmt.Errorf("create db dir: %w", err)
	}

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("open sqlite: %w", err)
	}

	store := &Store{Path: path, DB: db}
	if err := db.Ping(); err == nil {
		store.DriverReady = true
	}
	return store, nil
}

func (s *Store) EnsureSchema() error {
	if !s.DriverReady {
		return nil
	}

	stmts := []string{
		`CREATE TABLE IF NOT EXISTS app_config (
			id INTEGER PRIMARY KEY,
			key TEXT NOT NULL UNIQUE,
			value TEXT NOT NULL,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);`,
		`CREATE TABLE IF NOT EXISTS file_hunter_entries (
			id INTEGER PRIMARY KEY,
			title TEXT NOT NULL,
			path TEXT NOT NULL,
			size_bytes INTEGER,
			crc32 TEXT,
			collected_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);`,
		`CREATE TABLE IF NOT EXISTS game_db_roms (
			id INTEGER PRIMARY KEY,
			title TEXT NOT NULL,
			publisher TEXT,
			year TEXT,
			sha1 TEXT,
			md5 TEXT,
			crc32 TEXT,
			source TEXT DEFAULT 'msxromsdb'
		);`,
	}

	for _, stmt := range stmts {
		if _, err := s.DB.Exec(stmt); err != nil {
			return fmt.Errorf("apply schema: %w", err)
		}
	}
	return nil
}
