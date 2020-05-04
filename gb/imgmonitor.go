package gb

import (
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

type ImgMonitor struct {
	BaseMonitor
}

type ImgRenderer struct {
	screen *[160][144][3] uint8
	inc int
	x int
}

func (m *ImgMonitor) Init(screen *[160][144][3] uint8, title string) error {
	m.renderer = &ImgRenderer{
		screen: screen,
		x: 4,
	}
	return m.BaseMonitor.Init(screen, title)
}

func (r *ImgRenderer) Render() {
	if r.inc < 60 {
		r.inc++
		return
	}
	r.inc = 0
	img := image.NewRGBA(image.Rect(0, 0, 160 * r.x, 144 * r.x))
	for x := 0;x < 160 * r.x; x++ {
		for y := 0; y < 144 * r.x; y++ {
			px := r.screen[x / r.x][y / r.x]
			img.SetRGBA(x, y, color.RGBA{
				R: px[0],
				G: px[1],
				B: px[2],
				A: 0,
			})
		}
	}

	imgFile, _ := os.OpenFile("imgMonitor.jpg", os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0666)
	defer imgFile.Close()
	jpeg.Encode(imgFile, img, &jpeg.Options{Quality:100})
}