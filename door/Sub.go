package door

import "fmt"

// Amplifier 音响系统
type Amplifier struct {
	volume int
}

func (a *Amplifier) On() {
	fmt.Println("Amplifier is on")
}

func (a *Amplifier) SetVolume(volume int) {
	a.volume = volume
	fmt.Printf("Amplifier volume set to %d\n", volume)
}

func (a *Amplifier) Off() {
	fmt.Println("Amplifier is off")
}

// Projector 投影仪
type Projector struct{}

func (p *Projector) On() {
	fmt.Println("Projector is on")
}

func (p *Projector) SetInput(input string) {
	fmt.Printf("Projector input set to %s\n", input)
}

func (p *Projector) Off() {
	fmt.Println("Projector is off")
}

// DVDPlayer DVD播放器
type DVDPlayer struct {
	movie string
}

func (d *DVDPlayer) On() {
	fmt.Println("DVD Player is on")
}

func (d *DVDPlayer) Play(movie string) {
	d.movie = movie
	fmt.Printf("DVD Player playing: %s\n", movie)
}

func (d *DVDPlayer) Stop() {
	fmt.Printf("DVD Player stopped: %s\n", d.movie)
}

func (d *DVDPlayer) Off() {
	fmt.Println("DVD Player is off")
}

// Lights 灯光系统
type Lights struct {
	intensity int
}

func (l *Lights) Dim(level int) {
	l.intensity = level
	fmt.Printf("Lights dimmed to %d%%\n", level)
}

func (l *Lights) On() {
	fmt.Println("Lights are on")
}

// Screen 投影屏幕
type Screen struct{}

func (s *Screen) Down() {
	fmt.Println("Screen is down")
}

func (s *Screen) Up() {
	fmt.Println("Screen is up")
}
