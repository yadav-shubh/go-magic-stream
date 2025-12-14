package utils

import (
	"os"
	"os/exec"
	"path/filepath"

	"go.uber.org/zap"
)

// StartDockerCompose starts the docker-compose services defined in the given config file.
// configPath should be relative to the current working directory or absolute.
func StartDockerCompose(configPath string) error {
	absPath, err := filepath.Abs(configPath)
	if err != nil {
		Log.Error("Failed to get absolute path for docker-compose file", zap.Error(err))
		return err
	}

	Log.Info("Starting Docker Compose...", zap.String("file", absPath))

	// Project name derived from directory name or fixed to avoid conflicts if needed.
	// We'll let docker compose handle default project naming for now.

	cmd := exec.Command("docker", "compose", "-f", absPath, "up", "-d")

	// Capture output for debugging if needed, but for now just inherit stdout/stderr
	// or we can just log output if it fails.
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		Log.Error("Failed to start docker compose", zap.Error(err))
		return err
	}

	Log.Info("Docker Compose started successfully")
	return nil
}
