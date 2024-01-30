package keyboards

import (
	"fmt"
	"os"

	letsSplit "github.com/vunhatchuong/go-keymapviz/keyboards/lets_split"
	"github.com/vunhatchuong/go-keymapviz/keyboards/sofle"
)

var Keyboards = map[string]map[string]string{
	"sofle":      sofle.Layout,
	"lets_split": letsSplit.Layout,
}

var Layouts = map[string]string{
	"ascii": "ASCII",
	"fancy": "FANCY",
}

func LoadArtTemplate(kb string, layout string) ([]byte, error) {
	layoutValue, err := CheckLayoutForKeyboardExist(kb, layout)
	if err != nil {
		return nil, err
	}
	template, err := os.ReadFile(Keyboards[kb][layoutValue])
	if err != nil {
		return nil, fmt.Errorf("Can't load template: %v", err)
	}
	return template, nil
}

// Return the value of the given layout, layout = Layouts[layout]
func CheckLayoutForKeyboardExist(kb string, layout string) (string, error) {
	if _, ok := Layouts[layout]; ok == true {
		layout = Layouts[layout]
	}

	if kbLayout, ok := Keyboards[kb]; ok == false {
		return "", fmt.Errorf("Keyboard: %v doesn't exist\n-h for available keyboards", kb)
	} else if _, ok := kbLayout[layout]; ok == false {
		return "", fmt.Errorf("Keyboard: %v doesn't exist\n-h for available keyboards", kb)
	}
	return layout, nil
}
