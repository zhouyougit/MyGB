package gb

type Monitor interface {
	InitMonitor(*[160][144][3] uint8, string) error

	RenderScreen()

	Run(chan struct{})

	AddExitListener(func())
}

type MonitorRenderer interface {
	Render()
}

type BaseMonitor struct {
	screen *[160][144][3] uint8
	title string

	renderSignal chan struct{}

	stoped bool

	renderer MonitorRenderer
}

func (m *BaseMonitor) InitMonitor(screen *[160][144][3] uint8, title string) error {
	m.screen = screen
	m.title = title
	m.renderSignal = make(chan struct{}, 1)
	return nil
}

func (m *BaseMonitor) RenderScreen() {
	m.renderSignal <- struct{}{}
}

func (m *BaseMonitor) Run(stop chan struct{}) {
	for !m.stoped{
		select {
		case <-m.renderSignal:
			m.renderer.Render()
		case <-stop:
			m.stoped = true
		}
	}
}

func (m *BaseMonitor) AddExitListener(f func()) {

}