package static

import (
	"fmt"
	"os"
	"path/filepath"
)

// BuildStaticDirectories Will build the templates and static directories in the provided path
func BuildStaticDirectories(path string) (err error) {
	if _, err := os.Stat(filepath.Clean(fmt.Sprintf("%s/templates", path))); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}

	if _, err := os.Stat(filepath.Clean(fmt.Sprintf("%s/static", path))); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
