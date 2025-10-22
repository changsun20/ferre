package core

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/changsun20/ferre/internal/utils"
	"github.com/urfave/cli/v3"
)

func InstallAction(ctx context.Context, c *cli.Command) error {
	packageName := c.Args().First()
	if packageName == "" {
		return fmt.Errorf("please provide a package name to install")
	}

	homePath, err := GetFerreHome()
	if err != nil {
		return fmt.Errorf("Error: %v\n", err)
	}

	manifestPath := fmt.Sprintf("%s/Applications/buckets/main/bucket/%s.json", homePath, packageName)

	_, err = os.Stat(manifestPath)
	if os.IsNotExist(err) {
		return fmt.Errorf("Error: package %s does not exist in the main bucket\n", packageName)
	} else if err != nil {
		return fmt.Errorf("Error: failed to access package %s: %v\n", packageName, err)
	}

	manifest, err := utils.ParseInstallInfoJSON(manifestPath)
	if err != nil {
		return fmt.Errorf("Error: failed to parse manifest for %s: %v\n", packageName, err)
	}

	softwareName := filepath.Base(manifest.URL)
	destPath := filepath.Join(homePath, "Cache", softwareName)

	fmt.Printf("Installing package: %s\n", packageName)

	err = utils.DownloadFile(manifest.URL, destPath)
	if err != nil {
		return fmt.Errorf("Error: failed to download package %s: %v\n", packageName, err)
	}

	Success("✓ Package %s installed successfully\n", packageName)
	return nil
}
