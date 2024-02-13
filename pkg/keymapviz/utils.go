package keymapviz

import (
	"fmt"
	"os"

	"github.com/gookit/ini/v2"
)

func LoadLegends(legendPath string) (map[string]string, error) {
	err := ini.LoadExists(legendPath)
	if err != nil {
		return nil, fmt.Errorf("Can't load legends file: %v", err)
	}

	legends := ini.StringMap("legends")

	return legends, nil
}

func LoadWrapper(wrapperPath string) (string, error) {
	wrapper, err := os.ReadFile(wrapperPath)
	if err != nil {
		return "", fmt.Errorf("Can't load wrapper: %v", err)
	}

	return string(wrapper), nil
}
