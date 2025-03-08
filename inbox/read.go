package inbox

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Path is the path to the inbox folder
const Path = "/data/inbox"

// GetAll returns a list of all logs in the inbox folder
func GetAll() ([]Log, error) {
	// Walk the inbox folder, creating a list of paths to log files
	paths, err := walk()
	if err != nil {
		return []Log{}, err
	}

	// Create a list of logs from the paths
	var logs []Log
	for _, path := range paths {
		log, err := get(path)
		if err != nil {
			return []Log{}, fmt.Errorf("failed to get log: %w", err)
		}
		logs = append(logs, log)
	}

	// Return the list of logs
	return logs, nil
}

// get reads the log file at the given path along with the corresponding meta file
// and returns a Log object
func get(path string) (Log, error) {
	// Open the log file
	logFile, err := os.Open(path)
	if err != nil {
		return Log{}, fmt.Errorf("failed to open log file: %w", err)
	}
	defer logFile.Close()

	// Open the meta file
	metaPath := path + ".meta.json"
	metaFile, err := os.Open(metaPath)
	if err != nil {
		return Log{}, fmt.Errorf("failed to open meta file: %w", err)
	}
	defer metaFile.Close()

	// Read the meta file
	metaData, err := io.ReadAll(metaFile)
	if err != nil {
		return Log{}, fmt.Errorf("failed to read meta file: %w", err)
	}

	// Unmarshal the meta data
	var meta Meta
	err = json.Unmarshal(metaData, &meta)
	if err != nil {
		return Log{}, fmt.Errorf("failed to unmarshal meta data: %w", err)
	}

	// Return the log object
	return Log{
		ID:       strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)),
		File:     logFile,
		Meta:     meta,
		LogPath:  path,
		MetaPath: metaPath,
	}, nil
}

// walk walks the inbox folder and returns a list of log files
func walk() ([]string, error) {
	var paths []string
	err := filepath.Walk(Path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("failed to walk '%s': %w", path, err)
		}

		if filepath.Ext(path) == ".log" {
			paths = append(paths, path)
		}

		return nil
	})
	if err != nil {
		return []string{}, err
	}
	return paths, nil
}
