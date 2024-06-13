package facade

import "fmt"

// HomeTheaterFacade simplifies the interface for using a home theater system
type HomeTheaterFacade struct {
	amplifier     *Amplifier
	dvdPlayer     *DvdPlayer
	projector     *Projector
	theaterLights *TheaterLights
	screen        *Screen
}

// NewHomeTheaterFacade creates a new instance of HomeTheaterFacade
func NewHomeTheaterFacade(amplifier *Amplifier, dvdPlayer *DvdPlayer, projector *Projector, theaterLights *TheaterLights, screen *Screen) *HomeTheaterFacade {
	return &HomeTheaterFacade{amplifier, dvdPlayer, projector, theaterLights, screen}
}

// WatchMovie simplifies the process of watching a movie
func (h *HomeTheaterFacade) WatchMovie(movie string) {
	h.theaterLights.Dim(10)
	h.screen.Down()
	h.projector.On()
	h.amplifier.On()
	h.dvdPlayer.On()
	h.dvdPlayer.Play(movie)
}

// EndMovie simplifies the process of ending the movie
func (h *HomeTheaterFacade) EndMovie() {
	h.dvdPlayer.Stop()
	h.dvdPlayer.Eject()
	h.dvdPlayer.Off()
	h.amplifier.Off()
	h.projector.Off()
	h.screen.Up()
	h.theaterLights.On()
}

// Amplifier, DvdPlayer, Projector, TheaterLights, and Screen are placeholders for complex subsystems
type Amplifier struct{}
type DvdPlayer struct{}
type Projector struct{}
type TheaterLights struct{}
type Screen struct{}

// Placeholder methods for the subsystems
func (a *Amplifier) On() {
	fmt.Println("Amplifier is turned on")
}
func (a *Amplifier) Off() {
	fmt.Println("Amplifier is turned off")
}
func (d *DvdPlayer) On() {
	fmt.Println("DvdPlayer is turned on")
}
func (d *DvdPlayer) Off() {
	fmt.Println("DvdPlayer is turned off")
}
func (d *DvdPlayer) Play(movie string) {
	fmt.Printf("DvdPlayer is playing movie: %s\n", movie)
}
func (d *DvdPlayer) Stop() {
	fmt.Println("DvdPlayer has stopped")
}
func (d *DvdPlayer) Eject() {
	fmt.Println("DvdPlayer has ejected the disk")
}
func (p *Projector) On() {
	fmt.Println("Projector is turned on")
}
func (p *Projector) Off() {
	fmt.Println("Projector is turned off")
}
func (t *TheaterLights) On() {
	fmt.Println("TheaterLights are turned on")
}
func (t *TheaterLights) Dim(intensity int) {
	fmt.Printf("TheaterLights are dimmed to intensity: %d\n", intensity)
}
func (t *TheaterLights) Off() {
	fmt.Println("TheaterLights are turned off")
}
func (s *Screen) Up() {
	fmt.Println("Screen is up")
}
func (s *Screen) Down() {
	fmt.Println("Screen is down")
}
