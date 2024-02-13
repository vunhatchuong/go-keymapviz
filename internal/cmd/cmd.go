package cmd

import (
	"flag"
	"fmt"
)

type cmdFlags struct {
	Keyboard    string
	Layout      string
	LegendsPath string
	WrapperPath string
	KeymapPath  string
}

func NewCmdFlags() *cmdFlags {
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

	return &cmdFlags{
		Keyboard:    *keyboard,
		Layout:      *layout,
		LegendsPath: *legendsPath,
		WrapperPath: *wrapperPath,
		KeymapPath:  keymapPath,
	}
}
