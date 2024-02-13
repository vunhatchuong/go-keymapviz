package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/vunhatchuong/go-keymapviz/internal/cmd"
	"github.com/vunhatchuong/go-keymapviz/internal/tui"
	"github.com/vunhatchuong/go-keymapviz/pkg/keymapviz"
)

func main() {
	cmdFlags := cmd.NewCmdFlags()
	if cmdFlags.Keyboard == "" && cmdFlags.KeymapPath == "" {
		model := tui.NewModel()
		model.InitList()
		p := tea.NewProgram(model, tea.WithAltScreen(), tea.WithMouseCellMotion())

		if _, err := p.Run(); err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(0)
		}

		os.Exit(0)
	}

	if cmdFlags.Keyboard == "" {
		fmt.Println("Missing keyboard argument -kb")
		os.Exit(0)
	}

	if cmdFlags.KeymapPath == "" {
		fmt.Println("Missing keymap path, example: keymapviz -kb sofle ./sofle.c")
		os.Exit(0)
	}

	keymapviz, err := keymapviz.NewKeymapviz(
		cmdFlags.KeymapPath,
		cmdFlags.Keyboard,
		cmdFlags.Layout,
		cmdFlags.LegendsPath,
		cmdFlags.WrapperPath,
	)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(0)
	}

	keymaps, err := keymapviz.ExtractKeymaps()
	if err != nil {
		fmt.Printf("Can't load keymaps: %v", err)
		os.Exit(0)
	}

	keymapviz.OutputStdout(keymaps)
}
