package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/vunhatchuong/go-keymapviz/pkg/keymapviz"
)

func main() {
	keyboard := flag.String("kb", "", "Keyboard name")
	layout := flag.String("t", "ascii", "ASCII layout output")
	legendsPath := flag.String("l", "", "Custom legends for layout")
	wrapperPath := flag.String("w", "", "Wrapper path")

	flag.Usage = func() {
		fmt.Println("Usage: keymapviz [-h] [-l LEGENDS] [-t {ascii,fancy}] [-w WRAPPERS] keymap_c")
		fmt.Println("\nExample: keymapviz -kb sofle -t fancy ./sofle.c")
		fmt.Println("\noptions:")
		flag.PrintDefaults()
	}
	flag.Parse()
	keymapPath := flag.Arg(0)

	if *keyboard == "" {
		log.Fatal("Missing keyboard argument -kb")
	}

	if keymapPath == "" {
		log.Fatal("Missing keymap path, example: keymapviz -kb sofle ./sofle.c")
	}

	keymapviz, err := keymapviz.NewKeymapviz(
		keymapPath,
		*keyboard,
		*layout,
		*legendsPath,
		*wrapperPath,
	)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	keymaps, err := keymapviz.ExtractKeymaps()
	if err != nil {
		log.Fatalf("Can't load keymaps: %v", err)
	}
	keymapviz.OutputStdout(keymaps)
}
