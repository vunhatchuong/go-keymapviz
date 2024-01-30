package keymapviz

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/vunhatchuong/go-keymapviz/keyboards"
)

type Keymapviz struct {
	Keymap   string
	Keyboard string
	Layout   string
	Legends  map[string]string
	Wrapper  string
}

func NewKeymapviz(
	keymapPath string,
	kb string,
	layout string,
	legendPath string,
	wrapperPath string,
) (*Keymapviz, error) {
	fmt.Println(keymapPath)
	fmt.Println(kb)
	fmt.Println(layout)
	fmt.Println(legendPath)
	fmt.Println(wrapperPath)

	keymap, err := os.ReadFile(keymapPath)
	if err != nil {
		fmt.Printf("Can't load keymap file: %v", err)
		os.Exit(0)
	}

	var wrapper string
	if len(wrapperPath) != 0 {
		wrapper, err = LoadWrapper(wrapperPath)
		if err != nil {
			fmt.Printf("Can't load wrapper: %v", err)
			os.Exit(0)
		}
	}

	_, err = keyboards.CheckLayoutForKeyboardExist(kb, layout)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	var legend map[string]string
	if len(legendPath) != 0 {
		legend, err = LoadLegends(legendPath)
		if err != nil {
			fmt.Printf("Can't load legends file: %v", err)
			os.Exit(0)
		}
	}

	return &Keymapviz{
		Keymap:   string(keymap),
		Keyboard: kb,
		Layout:   layout,
		Legends:  legend,
		Wrapper:  wrapper,
	}, nil
}

func (kmz *Keymapviz) ExtractWrapper() (map[string]string, error) {
	wrapperExtractor := regexp.MustCompile(`(?m)#define (_[\S]+_)(.+)`)

	wrapperDefinition := wrapperExtractor.FindAllStringSubmatch(kmz.Wrapper, -1)
	wrapperMap := make(map[string]string, len(wrapperDefinition))
	for i := range wrapperDefinition {
		k := wrapperDefinition[i][1]
		v := wrapperDefinition[i][2]
		wrapperMap[k] = v
	}
	return wrapperMap, nil
}

func (kmz *Keymapviz) ExtractKeymaps() ([][]string, error) {
	getKeymapZone := regexp.MustCompile(
		`const uint16_t PROGMEM keymaps\[\]\[MATRIX_ROWS\]\[MATRIX_COLS\] = \{([\s\S]*?)\};`,
	)
	removeBlockComment := regexp.MustCompile(`/\*[\s\S]*?\*/`)
	getKeymapLayers := regexp.MustCompile(`\(([\s\S]*?)\n\s*\)`)
	getKeymaps := regexp.MustCompile(`([^\s,]+)`)

	found := getKeymapZone.FindStringSubmatch(kmz.Keymap)[1]
	found = removeBlockComment.ReplaceAllString(found, "")

	if len(kmz.Wrapper) != 0 {
		wrapper, err := kmz.ExtractWrapper()
		if err != nil {
			log.Fatalf("Can't extract wrapper: %v", err)
		}
		for k, v := range wrapper {
			found = strings.ReplaceAll(found, k, v)
		}
	}

	keymaps := getKeymapLayers.FindAllStringSubmatch(found, -1)

	keymap := make([][]string, len(keymaps))

	for i := range keymaps {
		keys := getKeymaps.FindAllStringSubmatch(keymaps[i][1], -1)
		keymap[i] = make([]string, len(keys))

		for j := range keys {
			if elem, ok := kmz.Legends[keys[j][1]]; ok == true {
				keymap[i][j] = elem
			} else {
				keymap[i][j] = keys[j][1]
			}
		}
	}
	return keymap, nil
}

func (kmz *Keymapviz) OutputStdout(keymaps [][]string) {
	template, err := keyboards.LoadArtTemplate(kmz.Keyboard, kmz.Layout)
	if err != nil {
		fmt.Printf("Can't load template: :%v", err)
		os.Exit(0)
	}
	templateStr := string(template)

	getPlaceHolder := regexp.MustCompile(`\{.*?\}`)
	fmt.Println(getPlaceHolder)

	for i, layer := range keymaps {
		currentLayer := templateStr
		for j := range layer {
			placeholder := getPlaceHolder.FindString(currentLayer)
			placeholderLen := len(placeholder)
			fmt.Printf("PlaceholderLen:%v\n", placeholderLen)
			key := keymaps[i][j]
			if len(key) > placeholderLen {
				key = key[:placeholderLen]
				fmt.Printf("Edited key %v", key)

			}

			subStr := fmt.Sprintf(
				fmt.Sprintf("%%-%ds", placeholderLen), // %-5s
				fmt.Sprintf(
					fmt.Sprintf("%%%ds", placeholderLen/2), // %5s
					key),
			)

			currentLayer = strings.Replace(
				currentLayer,
				placeholder,
				subStr,
				1,
			)
		}
		fmt.Println(currentLayer)
	}
}
