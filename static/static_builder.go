package static

import (
	"fmt"
	"os"
	"path/filepath"
)

func BuildStaticDirectories(path string) (err error) {
	path = filepath.Clean(fmt.Sprintf("%s/templates", path))
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}

	path = filepath.Clean(fmt.Sprintf("%s/static", path))
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
