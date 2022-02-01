package utils

import (
	"fmt"
	"os"
	"path"
	"time"
)

const healthFileName = "healthy"

func WriteHealthy(healthProbeDir string) error {
	err := makeIfNotExists(healthProbeDir)
	if err != nil {
		return fmt.Errorf("error ensuring health probe directory exists: %w", err)
	}

	healthPath := path.Join(healthProbeDir, healthFileName)

	err = makeIfNotExists(healthPath)
	if err != nil {
		return fmt.Errorf("error ensuring the health probe file exists: %w", err)
	}

	err = touch(healthPath)

	if err != nil {
		return fmt.Errorf("error updating mtime on health probe file: %w", err)
	}
	return nil
}

func Unhealthy(healthProbeDir string) error {
	healthPath := path.Join(healthProbeDir, healthFileName)
	_, err := os.Stat(healthPath)
	if os.IsNotExist(err) {
		return nil
	}

	if err != nil {
		return fmt.Errorf("failed to get file stats for health probe: %w", err)
	}

	err = os.Remove(healthPath)
	if err != nil {
		return fmt.Errorf("failed to delete the health probe file")
	}
	return nil
}

func makeIfNotExists(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.Mkdir(path, os.ModeDir)
	}
	if err != nil {
		return err
	}
	return nil
}

func touch(path string) error {
	currentTime := time.Now().Local()
	err := os.Chtimes(path, currentTime, currentTime)
	return err
}
