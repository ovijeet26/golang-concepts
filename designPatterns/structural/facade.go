package structural

import "fmt"

/*
The Facade Pattern is a structural design pattern that provides a simplified interface to a complex subsystem. It helps to hide the complexities of the system and provides an easy-to-use interface to the client.
*/

// DVDPlayer represents a DVD player.
type DVDPlayer struct{}

func (d *DVDPlayer) On() {
	fmt.Println("DVD Player is on.")
}

func (d *DVDPlayer) Play(movie string) {
	fmt.Printf("Playing movie: %s\n", movie)
}

func (d *DVDPlayer) Off() {
	fmt.Println("DVD Player is off.")
}

// Projector represents a projector.
type Projector struct{}

func (p *Projector) On() {
	fmt.Println("Projector is on.")
}

func (p *Projector) SetInput(input string) {
	fmt.Printf("Projector input set to %s.\n", input)
}

func (p *Projector) Off() {
	fmt.Println("Projector is off.")
}

// SoundSystem represents a sound system.
type SoundSystem struct{}

func (s *SoundSystem) On() {
	fmt.Println("Sound system is on.")
}

func (s *SoundSystem) SetVolume(level int) {
	fmt.Printf("Sound system volume set to %d.\n", level)
}

func (s *SoundSystem) Off() {
	fmt.Println("Sound system is off.")
}

// Lights represents the room lights.
type Lights struct{}

func (l *Lights) Dim(level int) {
	fmt.Printf("Lights dimmed to %d%%.\n", level)
}

func (l *Lights) On() {
	fmt.Println("Lights are on.")
}

// HomeTheaterFacade is the facade that provides a simple interface to control the home theater system.
type HomeTheaterFacade struct {
	dvdPlayer   *DVDPlayer
	projector   *Projector
	soundSystem *SoundSystem
	lights      *Lights
}

func NewHomeTheaterFacade(dvdPlayer *DVDPlayer, projector *Projector, soundSystem *SoundSystem, lights *Lights) *HomeTheaterFacade {
	return &HomeTheaterFacade{
		dvdPlayer:   dvdPlayer,
		projector:   projector,
		soundSystem: soundSystem,
		lights:      lights,
	}
}

func (h *HomeTheaterFacade) WatchMovie(movie string) {
	fmt.Println("Get ready to watch a movie...")
	h.lights.Dim(10)
	h.projector.On()
	h.projector.SetInput("DVD")
	h.soundSystem.On()
	h.soundSystem.SetVolume(5)
	h.dvdPlayer.On()
	h.dvdPlayer.Play(movie)
}

func (h *HomeTheaterFacade) EndMovie() {
	fmt.Println("Shutting down the home theater system...")
	h.dvdPlayer.Off()
	h.soundSystem.Off()
	h.projector.Off()
	h.lights.On()
}

// Facade caller
func FacadeCaller() {
	// Create instances of the subsystem components.
	dvdPlayer := &DVDPlayer{}
	projector := &Projector{}
	soundSystem := &SoundSystem{}
	lights := &Lights{}

	// Create the facade.
	homeTheater := NewHomeTheaterFacade(dvdPlayer, projector, soundSystem, lights)

	// Use the facade to watch a movie.
	homeTheater.WatchMovie("Inception")

	// Use the facade to end the movie.
	homeTheater.EndMovie()
}
