package inbox

import (
	"fmt"
	"os"
)

type Log struct {
	ID       string
	File     *os.File
	Meta     Meta
	LogPath  string
	MetaPath string
}

// Read reads the log file and returns its contents
func (l *Log) Read() ([]byte, error) {
	return os.ReadFile(l.LogPath)
}

// Close closes the log file
func (l *Log) Close() error {
	return l.File.Close()
}

// Delete deletes the log file and its meta data file
func (l *Log) Delete() error {
	l.Close()
	if err := os.Remove(l.LogPath); err != nil {
		return fmt.Errorf("failed to delete log file: %w", err)
	}
	if err := os.Remove(l.MetaPath); err != nil {
		return fmt.Errorf("failed to delete meta data file: %w", err)
	}
	return nil
}
