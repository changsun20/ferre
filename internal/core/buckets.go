package core

import (
	"context"
	"fmt"
	"os"

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
		Success("Main bucket created at %s\n", mainBucketPath)

		if err := initializeMainBucket(mainBucketPath); err != nil {
			return fmt.Errorf("Error: %v\n", err)
		}
		Success("Main bucket initialized with default apps\n")
		return nil
	}

	return nil
}

func initializeMainBucket(path string) error {
	_, err := git.PlainClone(path, &git.CloneOptions{
		URL:   "https://github.com/ScoopInstaller/Main.git",
		Depth: 1,
	})
	if err != nil {
		return fmt.Errorf("failed to initialize main bucket: %w", err)
	}
	return nil
}

func BucketRemoveAction(ctx context.Context, c *cli.Command) error {
	if c.NArg() < 1 {
		return fmt.Errorf("Error: Bucket name is required\n")
	}

	home, err := GetFerreHome()
	if err != nil {
		return fmt.Errorf("Error: %v\n", err)
	}

	bucketsDir := fmt.Sprintf("%s/Applications/buckets", home)

	name := c.Args().Get(0)
	bucketPath := fmt.Sprintf("%s/%s", bucketsDir, name)

	if _, err := os.Stat(bucketPath); os.IsNotExist(err) {
		return fmt.Errorf("Error: Bucket %s does not exist\n", name)
	}

	if err := os.RemoveAll(bucketPath); err != nil {
		return fmt.Errorf("Error: Failed to remove bucket %s: %v\n", name, err)
	}

	Success("Bucket %s removed successfully\n", name)
	return nil
}
