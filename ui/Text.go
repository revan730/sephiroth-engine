package ui

import (
	"io/ioutil"
	"os"
	"unicode"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"github.com/revan730/sephiroth-engine/resources"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
)

const fontsDir = "fonts"

var DefaultAtlas = text.NewAtlas(basicfont.Face7x13, text.ASCII, text.RangeTable(unicode.Latin))

func loadTTF(fontName string, size float64) (font.Face, error) {
	file, err := os.Open(resources.GetFontPath(fontName))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	font, err := truetype.Parse(bytes)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(font, &truetype.Options{
		Size:              size,
		GlyphCacheEntries: 1,
	}), nil
}

func NewText(fontName string, size float64, orig pixel.Vec) {
	if fontName == "" {
		// Use default
		return text.New(orig, DefaultAtlas)
	} else {
		// Load TTF font from game resources
		face, err := loadTTF(fontName, size)
		if err != nil {
			panic(err)
		}
		// TODO: Reuse atlases for same fonts
		atlas := text.NewAtlas(face, text.ASCII, text.RangeTable(unicode.Latin))
		return text.New(orig, atlas)
	}
}
