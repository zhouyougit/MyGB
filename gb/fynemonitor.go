package gb

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"image"
	"image/color"
)

type FyneMonitor struct {
	screen *[160][144][3]uint8
	title  string
	x      int

	app fyne.App
	img *image.RGBA

	renderSignal chan struct{}

	stopped bool

	exitListener []func()
}

type FyneLayout struct {
	raster *canvas.Raster
	x int
}

func (f FyneLayout) Layout(_ []fyne.CanvasObject, size fyne.Size) {
	f.raster.Resize(size)
}

func (f FyneLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(160 * f.x, 144 * f.x)
}

func (m *FyneMonitor) Init(screen *[160][144][3] uint8, title string) error {
	m.screen = screen
	m.title = title
	m.x = 2
	m.img = image.NewRGBA(image.Rect(0, 0, 160 * m.x, 144 * m.x))
	m.stopped = false
	m.renderSignal = make(chan struct{})
	return nil
}

func (m *FyneMonitor) RenderScreen() {
	m.renderSignal <- struct{}{}
}

func (m *FyneMonitor) AddExitListener(f func()) {
	m.exitListener = append(m.exitListener, f)
}

func (m *FyneMonitor) generateImage(w, h int) image.Image {
	for x := 0;x < 160 * m.x; x++ {
		for y := 0; y < 144 * m.x; y++ {
			px := m.screen[x / m.x][y / m.x]
			m.img.SetRGBA(x, y, color.RGBA{
				R: px[0],
				G: px[1],
				B: px[2],
				A: 0,
			})
		}
	}
	return m.img
}

func (m *FyneMonitor) Run(stop chan struct{}) {
	m.app = app.New()
	window := m.app.NewWindow(m.title)
	window.SetFixedSize(true)
	window.Resize(fyne.NewSize(160 * m.x, 144 * m.x))

	raster := canvas.NewRaster(m.generateImage)

	layout := &FyneLayout{
		raster: raster,
		x: m.x,
	}

	window.SetContent(fyne.NewContainerWithLayout(layout, raster))


	go func() {
		for !m.stopped {
			select {
			case <-m.renderSignal:
				canvas.Refresh(raster)
			case <-stop:
				fmt.Println("receive stop signal")
				m.stopped = true
				m.app.Quit()
			}
		}
	}()

	window.SetOnClosed(func() {
		for _, el := range m.exitListener {
			el()
		}
		fmt.Println("finish execute exitListener")
		m.app.Quit()
	})
	window.ShowAndRun()
}