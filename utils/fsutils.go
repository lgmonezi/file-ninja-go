package utils

import "path/filepath"

// MustAbsolutePath should be used when the user is sure that filepath.Abs won't fail. Panics otherwise.
// Returns the absolute path of file
func MustAbsolutePath(file string) string {
	path, err := filepath.Abs(file)
	if err != nil {
		panic(err)
	}
	return path
}
