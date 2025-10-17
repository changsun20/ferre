package core

import (
	"context"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/go-git/go-git/v6"
	"github.com/urfave/cli/v3"
)

func BucketAddAction(ctx context.Context, c *cli.Command) error {
	if c.NArg() < 1 {
		return fmt.Errorf("Error: Bucket name is required\n")
	}

	home, err := GetFerreHome()
	if err != nil {
		return fmt.Errorf("Error: %v\n", err)
	}

	bucketsDir := fmt.Sprintf("%s/Applications/buckets", home)

	name := c.Args().Get(0)
	if name == "main" {
		mainBucketPath := fmt.Sprintf("%s/main", bucketsDir)
		if err := os.MkdirAll(mainBucketPath, 0755); err != nil {
			return fmt.Errorf("Error: %v\n", err)
		}
		color.Green("Main bucket created at %s\n", mainBucketPath)

		if err := InitializeMainBucket(mainBucketPath); err != nil {
			return fmt.Errorf("Error: %v\n", err)
		}
		color.Green("Main bucket initialized with default apps\n")
		return nil
	}

	return nil
}

func InitializeMainBucket(path string) error {
	_, err := git.PlainClone(path, &git.CloneOptions{
		URL:   "https://github.com/ScoopInstaller/Main.git",
		Depth: 1,
	})
	if err != nil {
		return fmt.Errorf("failed to initialize main bucket: %w", err)
	}
	return nil
}
