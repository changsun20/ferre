package core

import (
	"context"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/changsun20/ferre/internal/pkgs"
	"github.com/urfave/cli/v3"
)

func SearchAction(ctx context.Context, c *cli.Command) error {
	query := c.Args().First()
	if query == "" {
		return fmt.Errorf("please provide a search query")
	}

	homePath, err := GetFerreHome()
	if err != nil {
		return fmt.Errorf("Error: %v\n", err)
	}

	path := fmt.Sprintf("%s/Applications/buckets/main/bucket", homePath)

	err = filepath.WalkDir(path, func(filePath string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if !d.IsDir() && strings.HasSuffix(strings.ToLower(d.Name()), ".json") {
			name := strings.TrimSuffix(d.Name(), ".json")

			if strings.Contains(strings.ToLower(name), strings.ToLower(query)) {
				manifest, parseErr := pkgs.ParseManifestJSON(filePath)
				if parseErr != nil {
					return nil
				}

				fmt.Printf("%s\t\t", name)
				fmt.Printf("%s\n", manifest.Version)
			}
		}

		return nil
	})

	return err
}
