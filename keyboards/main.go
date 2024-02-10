package keyboards

import (
	"fmt"
	"os"
)

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
	if kbLayout, ok := Keyboards[kb]; ok == false {
		return "", fmt.Errorf("Keyboard: %v doesn't exist\n-h for available keyboards", kb)
	} else if _, ok := kbLayout[layout]; ok == false {
		return "", fmt.Errorf("Keyboard: %v doesn't have %v layout\n-h for available keyboards", kb, layout)
	}
	return layout, nil
}
