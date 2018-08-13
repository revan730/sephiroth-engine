package resources

import (
	"os"
	"path"
)

const assetsDir = "assets"
const stringsDir = "strings"
const spritesheetsDir = "spritesheets"
const fontsDir = "fonts"

func GetFontPath(name string) string {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return path.Join(pwd, assetsDir, fontsDir, name+".ttf")
}
