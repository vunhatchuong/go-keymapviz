package keyboards

import (
	"github.com/vunhatchuong/go-keymapviz/keyboards/crkbd"
	"github.com/vunhatchuong/go-keymapviz/keyboards/dactyl_manuform5x6"
	"github.com/vunhatchuong/go-keymapviz/keyboards/ferris"
	"github.com/vunhatchuong/go-keymapviz/keyboards/sofle"
	"github.com/vunhatchuong/go-keymapviz/keyboards/sweet16"
)

var Keyboards = map[string]map[string]string{
	"sofle":              sofle.Layout,
	"crkbd":              crkbd.Layout,
	"dactyl_manuform5x6": dactyl_manuform5x6.Layout,
	"ferris":             ferris.Layout,
	"sweet16":            sweet16.Layout,
}
