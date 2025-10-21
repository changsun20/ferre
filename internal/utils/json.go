package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

type Manifest struct {
	Version     string `json:"version"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
}

func ParseManifestJSON(jsonPath string) (*Manifest, error) {
	file, err := os.Open(jsonPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open manifest file: %w", err)
	}
	defer file.Close()

	var manifest Manifest
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&manifest); err != nil {
		return nil, fmt.Errorf("failed to decode manifest JSON: %w", err)
	}
	return &manifest, nil
}
