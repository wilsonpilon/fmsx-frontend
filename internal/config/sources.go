package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type SourcesConfig struct {
	FileHunter FileHunterConfig `json:"file_hunter"`
	GameDB     GameDBConfig     `json:"game_db"`
	FMSX       FMSXConfig       `json:"fmsx"`
}

type FileHunterConfig struct {
	BaseURL      string `json:"base_url"`
	CatalogPath  string `json:"catalog_path"`
	DownloadPath string `json:"download_path"`
}

type GameDBConfig struct {
	ArchiveURL string `json:"archive_url"`
	LocalZip   string `json:"local_zip"`
	ExtractTo  string `json:"extract_to"`
}

type FMSXConfig struct {
	DocumentationURL string `json:"documentation_url"`
	BinaryPath       string `json:"binary_path"`
	DefaultMachine   string `json:"default_machine"`
}

func LoadSources(path string) (SourcesConfig, error) {
	var cfg SourcesConfig
	content, err := os.ReadFile(path)
	if err != nil {
		return cfg, fmt.Errorf("read %s: %w", path, err)
	}
	if err := json.Unmarshal(content, &cfg); err != nil {
		return cfg, fmt.Errorf("decode json from %s: %w", path, err)
	}
	return cfg, nil
}
