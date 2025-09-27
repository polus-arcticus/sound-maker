package songs

import (
	"fmt"
	"sound/pkg/sound"
)

type HarmonicStudyComposition struct {
	Name        string
	Description string
}

func (c *HarmonicStudyComposition) GetName() string {
	return c.Name
}

func (c *HarmonicStudyComposition) GetDescription() string {
	return c.Description
}

func (c *HarmonicStudyComposition) Play() error {
	fmt.Println("Creating a harmonic study composition")

	// Create a new song with 44.1kHz sample rate
	song := sound.NewSong("harmonic_study.bin", 44100.0, 264, 24)
	defer song.Close()

	// Define our notes using integer ratios
	C4 := 24 // C4 = 1:1 = 24:24 (unison)
	/*
		D4 := 27 // D4 = 9:8 = 27:24 (major whole tone)
		E4 := 30 // E4 = 5:4 = 30:24 (major third)
		F4 := 32 // F4 = 4:3 = 32:24 (perfect fourth)
		G4 := 36 // G4 = 3:2 = 36:24 (perfect fifth)
		A4 := 40 // A4 = 5:3 = 40:24 (major sixth)
		B4 := 45 // B4 = 15:8 = 45:24 (major seventh)
		C5 := 48 // C5 = 2:1 = 48:24 (octave)
	*/

	// Note duration
	noteDur := 1.0 // half note

	// Play harmonic study
	fmt.Println("Playing harmonic study...")
	song.AddOvertoneNotes(C4, noteDur, 1)
	song.AddOvertoneNotes(C4, noteDur, 3)
	song.AddOvertoneNotes(C4, noteDur, 5)
	song.AddOvertoneNotes(C4, noteDur, 7)
	song.AddOvertoneNotes(C4, noteDur, 9)
	song.AddOvertoneNotes(C4, noteDur, 11)
	song.AddOvertoneNotes(C4, noteDur, 13)
	song.AddOvertoneNotes(C4, noteDur, 15)

	// Play the song
	return song.Play()
}

func NewHarmonicStudyComposition() Composition {
	return &HarmonicStudyComposition{
		Name:        "harmonic_study",
		Description: "A harmonic study composition",
	}
}

func init() {
	RegisterComposition("harmonic_study", NewHarmonicStudyComposition)
}
