package songs

import (
	"fmt"
	"sound/pkg/sound"
)

// SimpleScaleComposition represents a simple scale composition
type SimpleScaleComposition struct {
	Name        string
	Description string
}

// GetName returns the name of the composition
func (c *SimpleScaleComposition) GetName() string {
	return c.Name
}

// GetDescription returns the description of the composition
func (c *SimpleScaleComposition) GetDescription() string {
	return c.Description
}

// Play performs the composition
func (c *SimpleScaleComposition) Play() error {
	fmt.Println("Creating a simple scale in 5-odd limit just intonation")

	// Create a new song with 44.1kHz sample rate
	song := sound.NewSong("simple_scale.bin", 44100.0, 264, 24)
	defer song.Close()

	// Define our notes using integer ratios
	C4 := 24 // C4 = 1:1 = 24:24 (unison)
	D4 := 27 // D4 = 9:8 = 27:24 (major whole tone)
	E4 := 30 // E4 = 5:4 = 30:24 (major third)
	F4 := 32 // F4 = 4:3 = 32:24 (perfect fourth)
	G4 := 36 // G4 = 3:2 = 36:24 (perfect fifth)
	A4 := 40 // A4 = 5:3 = 40:24 (major sixth)
	B4 := 45 // B4 = 15:8 = 45:24 (major seventh)
	C5 := 48 // C5 = 2:1 = 48:24 (octave)

	// Note duration
	noteDur := 0.5 // half note

	// Play ascending scale
	fmt.Println("Playing ascending scale...")
	song.AddNote(C4, noteDur)
	song.AddNote(D4, noteDur)
	song.AddNote(E4, noteDur)
	song.AddNote(F4, noteDur)
	song.AddNote(G4, noteDur)
	song.AddNote(A4, noteDur)
	song.AddNote(B4, noteDur)
	song.AddNote(C5, noteDur)

	// Play descending scale
	fmt.Println("Playing descending scale...")
	song.AddNote(C5, noteDur)
	song.AddNote(B4, noteDur)
	song.AddNote(A4, noteDur)
	song.AddNote(G4, noteDur)
	song.AddNote(F4, noteDur)
	song.AddNote(E4, noteDur)
	song.AddNote(D4, noteDur)
	song.AddNote(C4, noteDur)

	// Final chord
	fmt.Println("Playing final chord...")
	song.AddChord([]int{C4, E4, G4, C5}, noteDur*2)

	// Play the song
	return song.Play()
}

// NewSimpleScaleComposition creates a new simple scale composition
func NewSimpleScaleComposition() Composition {
	return &SimpleScaleComposition{
		Name:        "simple_scale",
		Description: "A simple scale in 5-odd limit just intonation",
	}
}

func init() {
	// Register this composition with the registry
	RegisterComposition("simple_scale", NewSimpleScaleComposition)
}
