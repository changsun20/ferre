package core

import (
	"context"
	"fmt"
	"os"

	"github.com/changsun20/ferre/internal/utils"
	"github.com/urfave/cli/v3"
)

func InfoAction(ctx context.Context, c *cli.Command) error {
	name := c.Args().First()
	if name == "" {
		return fmt.Errorf("please provide a name")
	}

	homePath, err := GetFerreHome()
	if err != nil {
		return fmt.Errorf("Error: %v\n", err)
	}

	filePath := fmt.Sprintf("%s/Applications/buckets/main/bucket/%s.json", homePath, name)
	_, err = os.Stat(filePath)
	if os.IsNotExist(err) {
		return fmt.Errorf("Error: application %s does not exist in the main bucket\n", name)
	} else if err != nil {
		return fmt.Errorf("Error: failed to access application %s: %v\n", name, err)
	}

	manifest, err := utils.ParseManifestJSON(filePath)
	if err != nil {
		return fmt.Errorf("Error: failed to parse manifest for %s: %v\n", name, err)
	}

	fmt.Printf("Name\t\t: %s\n", name)
	fmt.Printf("Version\t\t: %s\n", manifest.Version)
	fmt.Printf("Description\t: %s\n", manifest.Description)
	fmt.Printf("Homepage\t: %s\n", manifest.Homepage)

	return nil
}
