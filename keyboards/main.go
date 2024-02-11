package keyboards

import (
	"fmt"
)

func LoadArtTemplate(kb string, layout string) (string, error) {
	layoutValue, err := CheckLayoutForKeyboardExist(kb, layout)
	if err != nil {
		return "", err
	}
	template := Keyboards[kb][layoutValue]
	return template, nil
}

// Return the value of the given layout, layout = Layouts[layout]
func CheckLayoutForKeyboardExist(kb string, layout string) (string, error) {
	if kbLayout, ok := Keyboards[kb]; ok == false {
		return "", fmt.Errorf("Keyboard: %v doesn't exist\n", kb)
	} else if _, ok := kbLayout[layout]; ok == false {
		return "", fmt.Errorf("Keyboard: %v doesn't have %v layout\n", kb, layout)
	}
	return layout, nil
}
