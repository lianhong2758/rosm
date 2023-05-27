package draw

import (
	"image"
	"image/color"
	"math"

	"github.com/FloatTech/gg"
	"github.com/disintegration/imaging"
)

// Shadow绘制阴影,圆角矩形
func Shadow(x int, y int, r float64, c color.Color) image.Image {
	ctx := gg.NewContext(x, y)
	ctx.SetColor(c)
	ctx.DrawRoundedRectangle(0, 0, float64(x), float64(y), r)
	ctx.Fill()
	return ctx.Image()
}

// DrawPolygon 画多边形
func DrawPolygon(n int) []gg.Point {
	result := make([]gg.Point, n)
	for i := 0; i < n; i++ {
		a := float64(i)*2*math.Pi/float64(n) - math.Pi/2
		result[i] = gg.Point{X: math.Cos(a), Y: math.Sin(a)}
	}
	return result
}

// DrawStars 画星星,最多6颗,左对齐
func DrawStars(side, all string, num int) image.Image {
	dc := gg.NewContext(500, 80)
	n := 5
	points := DrawPolygon(n)
	for x, i := 40, 0; i < num; x += 80 {
		dc.Push()
		dc.Translate(float64(x), 45)
		dc.Scale(30, 30) //大小
		for i := 0; i < n+1; i++ {
			index := (i * 2) % n
			p := points[index]
			dc.LineTo(p.X, p.Y)
		}
		dc.SetLineWidth(10)
		dc.SetHexColor(side) //线
		dc.StrokePreserve()
		dc.SetHexColor(all)
		dc.Fill()
		dc.Pop()
		i++
	}
	return dc.Image()
}

// ChangeLucency 更改透明度
func ChangeLucency(m image.Image, percentage float64) image.Image {
	bounds := m.Bounds()
	dx, dy := bounds.Dx(), bounds.Dy()
	nimg := image.NewRGBA64(bounds)
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			r, g, b, a := m.At(i, j).RGBA()
			opacity := uint16(float64(a) * percentage)
			r, g, b, a = nimg.ColorModel().Convert(color.NRGBA64{R: uint16(r), G: uint16(g), B: uint16(b), A: opacity}).RGBA()
			nimg.SetRGBA64(i, j, color.RGBA64{R: uint16(r), G: uint16(g), B: uint16(b), A: uint16(a)})
		}
	}
	return nimg
}

// Size 改变大小
func Size(im image.Image, w, h int) image.Image {
	return imaging.Resize(im, w, h, imaging.Lanczos)
}
