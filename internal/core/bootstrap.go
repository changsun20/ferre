package core

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"golang.org/x/sys/windows/registry"
)

func GetFerreHome() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %w", err)
	}
	return filepath.Join(homeDir, "ferre"), nil
}

func CreateDirectoryStructure(ferreHome string) error {
	dirs := []string{
		filepath.Join(ferreHome, "Applications", "apps"),
		filepath.Join(ferreHome, "Applications", "buckets"),
		filepath.Join(ferreHome, "Applications", "persist"),
		filepath.Join(ferreHome, "Cache"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	return nil
}

func CopyExecutable(targetDir string) error {
	exePath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to get executable path: %w", err)
	}

	ferreAppDir := filepath.Join(targetDir, "ferre")
	if err := os.MkdirAll(ferreAppDir, 0755); err != nil {
		return fmt.Errorf("failed to create app directory: %w", err)
	}

	targetPath := filepath.Join(ferreAppDir, "ferre.exe")

	sourceFile, err := os.Open(exePath)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer sourceFile.Close()

	destFile, err := os.Create(targetPath)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, sourceFile); err != nil {
		return fmt.Errorf("failed to copy file: %w", err)
	}

	return nil
}

func AddToUserPath(dir string) error {
	key, err := registry.OpenKey(registry.CURRENT_USER, `Environment`, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		return fmt.Errorf("failed to open registry key: %w", err)
	}
	defer key.Close()

	currentPath, _, err := key.GetStringValue("Path")
	if err != nil && err != registry.ErrNotExist {
		return fmt.Errorf("failed to read PATH: %w", err)
	}

	if contains(currentPath, dir) {
		fmt.Printf("Ferre already add to path.\n")
		return nil
	}

	var newPath string
	if currentPath == "" {
		newPath = dir
	} else {
		newPath = currentPath + ";" + dir
	}

	if err := key.SetStringValue("Path", newPath); err != nil {
		return fmt.Errorf("failed to write PATH: %w", err)
	}

	return nil
}

func contains(path, dir string) bool {
	pathElements := filepath.SplitList(path)
	for _, element := range pathElements {
		if filepath.Clean(element) == filepath.Clean(dir) {
			return true
		}
	}
	return false
}

func Bootstrap() error {
	ferreHome, err := GetFerreHome()
	if err != nil {
		return err
	}

	fmt.Printf("Setting up Ferre at: %s\n", ferreHome)

	fmt.Println("Creating directory structure...")
	if err := CreateDirectoryStructure(ferreHome); err != nil {
		return err
	}
	fmt.Println("✓ Directory structure created")

	appsDir := filepath.Join(ferreHome, "Applications", "apps")
	fmt.Println("Copying ferre.exe...")
	if err := CopyExecutable(appsDir); err != nil {
		return err
	}
	fmt.Println("✓ Executable copied")

	ferreBinDir := filepath.Join(appsDir, "ferre")
	fmt.Printf("Adding %s to PATH...\n", ferreBinDir)
	if err := AddToUserPath(ferreBinDir); err != nil {
		return err
	}
	fmt.Println("✓ Added to PATH")

	fmt.Println("\n✓ Ferre has been successfully initialized!")
	fmt.Println("\nPlease restart your terminal to use ferre.")
	return nil
}
