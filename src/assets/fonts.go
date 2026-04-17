package assets

import (
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/quasilyte/bitsweetfont"
)

var (
	font1 = bitsweetfont.New1_3()

	FontTinyX = text.NewGoXFace(bitsweetfont.New1())
	Font1X    = text.NewGoXFace(font1)
	Font2X    = text.NewGoXFace(bitsweetfont.Scale(font1, 2))
	Font3X    = text.NewGoXFace(bitsweetfont.Scale(font1, 3))

	fontTinyIface text.Face = FontTinyX
	font1iface    text.Face = Font1X
	font2iface    text.Face = Font2X
	font3iface    text.Face = Font3X

	FontTiny = &fontTinyIface
	Font1    = &font1iface
	Font2    = &font2iface
	Font3    = &font3iface
)
