package utils

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"
	"strings"

	"github.com/golang/freetype"
)

const fontFile = "source/d.ttf"

// DrawText 生成并保存到指定文件
func DrawText(text rune, path string, name string) {

	fontBytes, err := ioutil.ReadFile(fontFile) // 读取字体
	if err != nil {
		fmt.Println(err)
		return
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		fmt.Println(err)
		return
	}
	blue := color.RGBA{30, 144, 255, 255}
	bg := image.NewUniform(blue)
	fg := image.White

	rgba := image.NewRGBA(image.Rect(0, 0, 128, 128))
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(f)
	c.SetFontSize(96)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)

	pt := freetype.Pt(36, int(c.PointToFixed(96)>>6))
	c.DrawString(strings.ToUpper(string(text)), pt)
	os.MkdirAll(path, os.ModePerm)
	outFile, err := os.Create(path + name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outFile.Close()
	b := bufio.NewWriter(outFile)
	png.Encode(b, rgba)
	b.Flush()
}
