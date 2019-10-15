package router

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// BuildRouter takes a path and builds router directories and files
func BuildRouter(path string) (err error) {
	err = generateRouterGoFiles(path)
	if err != nil {
		return err
	}
	return nil
}

// generateRouterGoFiles generates Go files to build configurations
func generateRouterGoFiles(path string) (err error) {
	// Creates a directory called routes
	path = filepath.Clean(fmt.Sprintf("%s/routes", path))
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}

	d1 := []byte(`package routes

import (
	"github.com/flannel-dev-lab/cyclops/router"
	"net/http"
)

func GetRoutes() *router.Router {

	routes := router.New()

	routes.Get("/health", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		return
	})


	return routes
}
`)
	err = ioutil.WriteFile(filepath.Clean(fmt.Sprintf("%s/routes.go", path)), d1, 0644)
	if err != nil {
		return err
	}

	return nil
}
