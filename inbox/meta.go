package inbox

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// readMetaFile reads the data from the meta file for the given log file
func readMetaFile(path string) (Meta, error) {
	metaPath := path + ".meta.json"
	meta := Meta{LogPath: path, MetaPath: &metaPath}

	// Open the meta data file
	file, err := os.Open(metaPath)
	if err != nil {
		return meta, fmt.Errorf("failed to open meta data file: %w", err)
	}

	// Read the meta file
	data, err := io.ReadAll(file)
	if err != nil {
		return meta, fmt.Errorf("failed to read meta file: %w", err)
	}

	// Unmarshal the meta data
	err = json.Unmarshal(data, &meta)
	if err != nil {
		return meta, fmt.Errorf("failed to unmarshal meta data: %w", err)
	}
	return meta, nil
}

type Meta struct {
	Source       Source  `json:"source"`
	OriginalPath string  `json:"path"`
	LogPath      string  `json:"-"`
	MetaPath     *string `json:"-"`
	raw          []byte  `json:"-"`
}

// Bytes returns the raw bytes of the meta data
func (m *Meta) Bytes() []byte {
	return m.raw
}

// Source represents a log source
type Source struct {
	System    string   `json:"system"`
	Path      string   `json:"path"`
	Tags      []string `json:"tags"`
	Recursive bool     `json:"recursive"`
}
