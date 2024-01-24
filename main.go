package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hedhyw/rex/pkg/rex"
)

func main() {
	keymapFile, err := os.ReadFile("./keymaps/sofle_no_wrapper.c")
	if err != nil {
		log.Fatal(err)
	}
	keymapFileStr := string(keymapFile)
	keymapZone := rex.New(
		rex.Common.Text("const uint16_t PROGMEM keymaps[][MATRIX_ROWS][MATRIX_COLS] = {"),
		rex.Group.Define(
			rex.Common.Class(
				rex.Chars.Whitespace(),
				rex.Common.Raw(`\S`),
			).Repeat().ZeroOrMorePreferFewer(),
		),
		rex.Common.Text(`};`),
	).MustCompile()
	fmt.Println(keymapZone)

	blockComment := rex.New(
		rex.Common.Text(`/*`),
		rex.Common.Class(
			rex.Chars.Whitespace(),
			rex.Common.Raw(`\S`),
		).Repeat().ZeroOrMorePreferFewer(),
		rex.Common.Text(`*/`),
	).MustCompile()
	fmt.Println(blockComment)
	/* normalize := rex.New(
		   rex.Common.Class(
		   // rex.Chars.Whitespace(),
	            rex.Chars.Blank(),
		       ),
		   ).MustCompile() */

	captureKeymaps := rex.New(
		rex.Common.Text(`(`),
		rex.Group.Define(
			rex.Common.Class(
				rex.Chars.Whitespace(),
				rex.Common.Raw(`\S`),
			).Repeat().ZeroOrMorePreferFewer(),
		),
		rex.Common.Text(`)`),
	).MustCompile()
	fmt.Println(captureKeymaps)

	found := keymapZone.FindStringSubmatch(keymapFileStr)[1]
	found = blockComment.ReplaceAllString(found, "")
	// stripedWhiteSpace := removeWhiteSpace.ReplaceAllString(stripedBlockComment, "")
	keymaps := captureKeymaps.FindAllStringSubmatch(found, -1)
	keymap := make([]string, 0, len(keymaps))
	for i := range keymaps {
		keymap = append(keymap, keymaps[i][1])
	}
	fmt.Println(keymap)
}
