package quote

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

var (
	ErrUnsupportedFormat = errors.New("unsupported file format")
	ErrTextRequired      = errors.New("text is required")
	ErrAuthorRequired    = errors.New("author is required")
)

// Quote represents a single quote
type Quote struct {
	Text   string   `json:"text" yaml:"text"`
	Author string   `json:"author" yaml:"author"`
	Tags   []string `json:"tags" yaml:"tags"`
	Lang   string   `json:"lang" yaml:"lang"`
}

// Loader is responsible for loading quote files
type Loader struct{}

// NewLoader creates a new loader instance
func NewLoader() *Loader {
	return &Loader{}
}

// LoadFromFile loads quotes from a single file
func (l *Loader) LoadFromFile(filePath string) ([]Quote, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	ext := strings.ToLower(filepath.Ext(filePath))
	switch ext {
	case ".json":
		return l.parseJSON(data)
	case ".yaml", ".yml":
		return l.parseYAML(data)
	default:
		return nil, fmt.Errorf("%w: %s", ErrUnsupportedFormat, ext)
	}
}

// LoadFromFiles loads quotes from multiple files
func (l *Loader) LoadFromFiles(filePaths string) ([]Quote, error) {
	files := strings.Split(filePaths, ",")
	var allQuotes []Quote

	for _, file := range files {
		file = strings.TrimSpace(file)
		if file == "" {
			continue
		}

		quotes, err := l.LoadFromFile(file)
		if err != nil {
			return nil, err
		}
		allQuotes = append(allQuotes, quotes...)
	}

	return allQuotes, nil
}

// parseJSON parses JSON format quotes
func (l *Loader) parseJSON(data []byte) ([]Quote, error) {
	var quotes []Quote
	if err := json.Unmarshal(data, &quotes); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	if err := l.validateQuotes(quotes); err != nil {
		return nil, err
	}

	return quotes, nil
}

// parseYAML parses YAML format quotes
func (l *Loader) parseYAML(data []byte) ([]Quote, error) {
	var quotes []Quote
	if err := yaml.Unmarshal(data, &quotes); err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	if err := l.validateQuotes(quotes); err != nil {
		return nil, err
	}

	return quotes, nil
}

// validateQuotes validates quote data
func (l *Loader) validateQuotes(quotes []Quote) error {
	for i, quote := range quotes {
		if quote.Text == "" {
			return fmt.Errorf("quote at index %d: %w", i, ErrTextRequired)
		}
		if quote.Author == "" {
			return fmt.Errorf("quote at index %d: %w", i, ErrAuthorRequired)
		}
	}
	return nil
}
