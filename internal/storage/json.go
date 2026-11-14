// Package storage provides deterministic, atomic local artifact IO.
package storage

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Hash(b []byte) string {
	s := sha256.Sum256(b)
	return hex.EncodeToString(s[:])
}

func Encode(v any) ([]byte, error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return nil, err
	}
	return append(b, '\n'), nil
}

func Digest(v any) (string, error) {
	b, err := Encode(v)
	if err != nil {
		return "", err
	}
	return Hash(b), nil
}

func Decode(b []byte, v any) error {
	d := json.NewDecoder(bytes.NewReader(b))
	d.DisallowUnknownFields()
	if err := d.Decode(v); err != nil {
		return err
	}
	var extra any
	if err := d.Decode(&extra); !errors.Is(err, io.EOF) {
		return fmt.Errorf("expected one JSON document")
	}
	return nil
}

func Read(path string, v any) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if err := Decode(b, v); err != nil {
		return fmt.Errorf("decode %s: %w", path, err)
	}
	return nil
}

func WriteJSON(path string, v any) error {
	b, err := Encode(v)
	if err != nil {
		return err
	}
	return Write(path, b)
}

// Write preserves mtime for identical bytes. Rename also avoids modifying the
// target of an existing hard link. A multi-file generation needs its own final
// manifest; this function alone is not a transaction across files.
func Write(path string, b []byte) error {
	if info, err := os.Lstat(path); err == nil && !info.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", path)
	}
	old, err := os.ReadFile(path)
	if err == nil && bytes.Equal(old, b) {
		return nil
	}
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	f, err := os.CreateTemp(filepath.Dir(path), ".write-*")
	if err != nil {
		return err
	}
	defer func() { _ = os.Remove(f.Name()) }()
	if _, err := f.Write(b); err != nil {
		_ = f.Close()
		return err
	}
	if err := f.Chmod(0o644); err != nil {
		_ = f.Close()
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return os.Rename(f.Name(), path)
}
