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
	CheckLayoutForKeyboardExist(kb, layout)
    layout = Layouts[layout]

	template, err := os.ReadFile(Keyboards[kb][layout])
	if err != nil {
		return nil, fmt.Errorf("Can't load template: %v", err)
	}
	return template, nil
}

func CheckLayoutForKeyboardExist(kb string, layout string) bool {
	if _, ok := Layouts[layout]; ok == true {
		layout = Layouts[layout]
	}

	if kbLayout, ok := Keyboards[kb]; ok == false {
		fmt.Printf("Keyboard: %v doesn't exist\n-h for available keyboards", kb)
		os.Exit(0)
	} else if _, ok := kbLayout[layout]; ok == false {
		fmt.Printf("Keyboard: %v doesn't have layout: %v\nAvailable layout: %v", kb, layout, kbLayout)
		os.Exit(0)
	}
	return true
}
