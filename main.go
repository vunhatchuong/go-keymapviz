package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/gookit/ini/v2"
)

func loadLegends(legendFile string) (map[string]string, error) {
	err := ini.LoadExists(legendFile)
	if err != nil {
		return nil, err
	}
	legends := ini.StringMap("legends")
	return legends, nil
}

func loadArtTemplate(templateFile string) ([]byte, error) {
	template, err := os.ReadFile(templateFile)
	if err != nil {
		log.Fatalf("Can't load template: %v", err)
	}
	return template, nil
}

func wrapperExtractor(wrapperFile string) map[string]string {
	wrapper, err := os.ReadFile(wrapperFile)
	if err != nil {
	}
	wrapperStr := string(wrapper)
	wrapperExtractor := regexp.MustCompile(`(?m)#define (_[\S]+_)(.+)`)
	fmt.Printf("Wrapper: %v\n", wrapperExtractor.String())
	wrapperDefinition := wrapperExtractor.FindAllStringSubmatch(wrapperStr, -1)
	wrapperMap := make(map[string]string, len(wrapperDefinition))
	for i := range wrapperDefinition {
		k := wrapperDefinition[i][1]
		v := wrapperDefinition[i][2]
		wrapperMap[k] = v
	}
	return wrapperMap
}

func extractKeymaps(keymapfile string) ([][]string, error) {
	keymapFile, err := os.ReadFile(keymapfile)
	if err != nil {
		log.Fatal(err)
	}

	keymapFileStr := string(keymapFile)
	keymapZone := regexp.MustCompile(
		`const uint16_t PROGMEM keymaps\[\]\[MATRIX_ROWS\]\[MATRIX_COLS\] = \{([\s\S]*?)\};`,
	)

	fmt.Printf("Keymapzone: %v\n", keymapZone)

	removeBlockComment := regexp.MustCompile(`/\*[\s\S]*?\*/`)
	fmt.Printf("removeBlockComment: %v\n", removeBlockComment)

	captureKeymaps := regexp.MustCompile(`\(([\s\S]*?)\n\s*\)`)
	fmt.Printf("captureKeymaps: %v\n", captureKeymaps)

	getKeys := regexp.MustCompile(`([^\s,]+)`)
	fmt.Printf("getKeys: %v\n", getKeys)

	found := keymapZone.FindStringSubmatch(keymapFileStr)[1]
	found = removeBlockComment.ReplaceAllString(found, "")
	wrapper := wrapperExtractor("./wrappers.h")
	for k, v := range wrapper {
		found = strings.ReplaceAll(found, k, v)
	}

	fmt.Printf("Found: %v", found)
	keymaps := captureKeymaps.FindAllStringSubmatch(found, -1)

	keymap := make([][]string, len(keymaps))

	legends, err := loadLegends("legends.ini")
	if err != nil {
		log.Fatalf("Can't load legends.ini: %v", err)
	}

	for i := range keymaps {
		keys := getKeys.FindAllStringSubmatch(keymaps[i][1], -1)
		keymap[i] = make([]string, len(keys))

		for j := range keys {
			if elem, ok := legends[keys[j][1]]; ok == true {
				keymap[i][j] = elem
			} else {
				keymap[i][j] = keys[j][1]
			}
		}
	}
	return keymap, nil
}

func main() {
	keymap, err := extractKeymaps("./replace_lets_split_keymap_fancy.c")
	if err != nil {
		log.Fatalf("Can't load keymaps: %v", err)
	}
	for _, key := range keymap {
		fmt.Println(len(key))
	}

	template, err := loadArtTemplate("./keyboards/lets_split/ascii.tmpl")
	if err != nil {
		log.Fatalf("Can't load template: %v", err)
	}
	templateStr := string(template)

	sub := regexp.MustCompile(`\{.*?\}`)
	fmt.Println(sub)

	for i, layer := range keymap {
		currentLayer := templateStr
		for j := range layer {
			subString := sub.FindString(currentLayer)
			currentLayer = strings.Replace(
				currentLayer,
				subString,
				fmt.Sprintf(
					fmt.Sprintf("%%-%ds", len(subString)),
					fmt.Sprintf(
						fmt.Sprintf("%%%ds", (len(subString)+len(subString))/2),
						keymap[i][j],
					),
				),
				1,
			)
		}
		fmt.Println(currentLayer)
	}
}
