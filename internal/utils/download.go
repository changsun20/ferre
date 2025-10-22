package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func CreateClient() *http.Client {
	return &http.Client{
		Timeout: 10 * time.Minute,
		Transport: &http.Transport{
			ResponseHeaderTimeout: 30 * time.Second,
		},
	}
}

func DownloadFile(url, destPath string) error {
	dir := filepath.Dir(destPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	client := CreateClient()

	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed with status: %d", resp.StatusCode)
	}

	tempFile := destPath + ".tmp"
	out, err := os.Create(tempFile)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		out.Close()
		os.Remove(tempFile)
		return fmt.Errorf("failed to write file: %w", err)
	}

	if err := out.Close(); err != nil {
		os.Remove(tempFile)
		return fmt.Errorf("failed to close file: %w", err)
	}

	if err := os.Rename(tempFile, destPath); err != nil {
		os.Remove(tempFile)
		return fmt.Errorf("failed to rename file: %w", err)
	}

	return nil
}
