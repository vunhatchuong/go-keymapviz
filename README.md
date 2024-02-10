# Go-Keymapviz

Go version of [keymapviz](https://github.com/yskoht/keymapviz) with different approach to extracting keymaps.

Command-line tool to convert [qmk_firmware](https://github.com/qmk/qmk_firmware/) keymap.c to ascii art.

## Installation

```bash
go install github.com/vunhatchuong/go-keymapviz/main
```

## Usage

Basic usage:

```bash
go-keymapviz [-h] [-l LEGENDS] [-t {ascii,fancy}] [-w WRAPPERS] keymap_c
```

Example:

```bash
go-keymapviz -kb sofle -t fancy ./sofle.c
```

You can checkout available keyboards and their ascii templates when run the program without any arguments.

```bash
go-keymapviz
```

## Convention

The program looks for specific pattern in keymap.c to extract correctly.

### Layout Zone

Layout zone is the entry point of your keymap, everything around it will be discarded.

It specifically has to ends with `};`

```c
const uint16_t PROGMEM keymaps[][MATRIX_ROWS][MATRIX_COLS] = {
                        // ----
};
```

### Keymap Layers

Keymap layers are thing inside `()` and the closing `)` **must** be on a newline.

It doesn't matter what the name of the layer or what kind of layout it is.
```c
[_QWERTY] = LAYOUT_ortho_4x12(
                        // ----
),
```

### Keymaps

The final step is to extract keycodes inside layers, keycodes are defined as any **word**.

Ex: `AU_OFF` is a word, `AU OFF` is 2 words.
