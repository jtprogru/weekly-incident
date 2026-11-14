package storage

import (
	"fmt"
	"os"
	"path/filepath"
)

// Lock serializes local writers in addition to the GitHub concurrency group.
// It does not guess that an existing lock is stale. After a process crash the
// operator can inspect its recorded PID before removing the lock file.
func Lock(root string) (func(), error) {
	if err := os.MkdirAll(root, 0o755); err != nil {
		return nil, err
	}
	path := filepath.Join(root, ".archive.lock")
	f, err := os.OpenFile(path, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o600)
	if err != nil {
		return nil, fmt.Errorf("archive writer lock %s: %w", path, err)
	}
	if _, err := fmt.Fprintf(f, "pid=%d\n", os.Getpid()); err != nil {
		_ = f.Close()
		_ = os.Remove(path)
		return nil, err
	}
	if err := f.Close(); err != nil {
		_ = os.Remove(path)
		return nil, err
	}
	return func() { _ = os.Remove(path) }, nil
}
