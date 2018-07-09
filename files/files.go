package files

import "os"

// Exists test if files exist
func Exists(name string) bool {
	_, err := os.Stat(name)
	return err == nil
}
