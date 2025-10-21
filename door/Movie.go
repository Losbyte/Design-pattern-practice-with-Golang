package door

import "fmt"

// HomeTheaterFacade 家庭影院外观类
type HomeTheaterFacade struct {
	amp       *Amplifier
	projector *Projector
	dvd       *DVDPlayer
	lights    *Lights
	screen    *Screen
}

func NewHomeTheaterFacade() *HomeTheaterFacade {
	return &HomeTheaterFacade{
		amp:       &Amplifier{},
		projector: &Projector{},
		dvd:       &DVDPlayer{},
		lights:    &Lights{},
		screen:    &Screen{},
	}
}

// WatchMovie 观看电影 - 简化复杂操作
func (h *HomeTheaterFacade) WatchMovie(movie string) {
	fmt.Println("Get ready to watch a movie...")

	h.lights.Dim(10)           // 调暗灯光
	h.screen.Down()            // 放下屏幕
	h.projector.On()           // 打开投影仪
	h.projector.SetInput("DVD") // 设置输入源
	h.amp.On()                 // 打开音响
	h.amp.SetVolume(5)        // 设置音量
	h.dvd.On()                 // 打开DVD
	h.dvd.Play(movie)         // 播放电影

	fmt.Println("Movie is ready!")
}

// EndMovie 结束电影
func (h *HomeTheaterFacade) EndMovie() {
	fmt.Println("Shutting down the home theater...")

	h.lights.On()             // 打开灯光
	h.screen.Up()             // 收起屏幕
	h.dvd.Stop()              // 停止DVD
	h.dvd.Off()               // 关闭DVD
	h.amp.Off()               // 关闭音响
	h.projector.Off()         // 关闭投影仪

	fmt.Println("Home theater is off")
}

// ListenToMusic 听音乐（简化版）
func (h *HomeTheaterFacade) ListenToMusic() {
	fmt.Println("Getting ready for music...")

	h.lights.On()             // 正常灯光
	h.amp.On()                // 打开音响
	h.amp.SetVolume(3)        // 设置合适音量

	fmt.Println("Ready for music!")
}
