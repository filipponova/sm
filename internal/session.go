package internal

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

func StartSession(region, profile, instanceID string) {
	cmd := exec.CommandContext(
		context.Background(),
		"aws", "ssm", "start-session",
		"--color", "on",
		"--region", region,
		"--profile", profile,
		"--target", instanceID,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	fmt.Printf("Running: aws ssm start-session --color on --region %s --profile %s --target %s\n", region, profile, instanceID)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Failed to start session: %v\n", err)
		return err
	}
	return nil
}
