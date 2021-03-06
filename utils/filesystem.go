package utils

import (
	"os"
	"path/filepath"
)

// GetDiskUsageStats accepts a path to a directory or file
// and returns the number of bytes and inodes used by the path
func GetDiskUsageStats(path string) (uint64, uint64, error) {
	var dirSize, inodeCount uint64

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		// Walk does not follow symbolic links
		if err != nil {
			return err
		}

		dirSize += uint64(info.Size())
		inodeCount++

		return nil
	})

	if err != nil {
		return 0, 0, err
	}

	return dirSize, inodeCount, err
}
