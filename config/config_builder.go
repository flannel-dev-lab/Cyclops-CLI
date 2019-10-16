package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// BuildConfig takes a path and configFileType and builds configuration directories
func BuildConfig(path, configFileType string) (err error) {
	switch configFileType {
	case "yaml":
		err = generateYAMLFiles(path)
		if err != nil {
			return err
		}
	}

	err = generateConfigurationGoFiles(path)
	if err != nil {
		return err
	}
	return nil
}

// generateYAMLFiles will generate yaml config files in the path given
func generateYAMLFiles(path string) (err error) {
	// Creates a directory called conf
	path = filepath.Clean(fmt.Sprintf("%s/conf", path))
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}

	// Creates the necessary YAML files with basic listen addr
	configFiles := []string{"dev.yaml", "prod.yaml", "stage.yaml", "local.yaml"}

	for _, configFile := range configFiles {
		configFilePath := filepath.Clean(fmt.Sprintf("%s/%s", path, configFile))
		d1 := []byte("listen_addr: \":80\"")
		err := ioutil.WriteFile(configFilePath, d1, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

// generateConfigurationGoFiles generates Go files to build configurations
func generateConfigurationGoFiles(path string) (err error) {
	// Creates a directory called config
	path = filepath.Clean(fmt.Sprintf("%s/config", path))
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}

	d1 := []byte(`package config

import "errors"
					
type ConfigurationService interface {
	ParseConfigFile(filePath string) error
	GetListenerAddress() string
}
					
// Retrieves configuration Object based on the file type
func CreateConfiguration(fileType string) (ConfigurationService, error) {
	switch fileType {
		case "yaml":
			return new(YAMLConfiguration), nil
		default:
			return nil, errors.New("undefined file type")
	}
}
					
`)
	err = ioutil.WriteFile(filepath.Clean(fmt.Sprintf("%s/configuration_service.go", path)), d1, 0644)
	if err != nil {
		return err
	}

	d1 = []byte(`package config
import (
	"io/ioutil"
)

type YAMLConfiguration struct {
	Listener       string         ` + `yaml:"listen_addr"` + `
}

func (yamlConfiguration *YAMLConfiguration) ParseConfigFile(filePath string) error {
	configFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(configFile, yamlConfiguration)

	if err != nil {
		return err
	}

	return nil
}

func (yamlConfiguration *YAMLConfiguration) GetListenerAddress() string {
	return yamlConfiguration.Listener
}
`)
	err = ioutil.WriteFile(filepath.Clean(fmt.Sprintf("%s/yaml_config.go", path)), d1, 0644)
	if err != nil {
		return err
	}

	return nil
}
