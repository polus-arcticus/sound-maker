package songs

import (
	"fmt"
	"sound/pkg/sound"
)

// JustIntonationComposition represents our 5-odd limit just intonation composition
type JustIntonationComposition struct {
	Name        string
	Description string
}

// GetName returns the name of the composition
func (c *JustIntonationComposition) GetName() string {
	return c.Name
}

// GetDescription returns the description of the composition
func (c *JustIntonationComposition) GetDescription() string {
	return c.Description
}

// Play performs the composition
func (c *JustIntonationComposition) Play() error {
	fmt.Println("Creating a 5-odd limit just intonation composition")

	// Create a new song with 44.1kHz sample rate
	// BaseFrequency = 264 Hz (C4)
	// BaseNote = 24 (the denominator for our ratios)
	song := sound.NewSong("just_intonation.bin", 44100.0, 264, 24)
	defer song.Close()

	// Define our notes using integer ratios
	// With the new system, we just specify the ratio numerator
	// The denominator is fixed at 24 (BaseNote)

	// 5-odd limit just intonation ratios (integer form)
	// All ratios are derived from powers of 2, 3, and 5

	// Base octave - using only 5-odd limit ratios
	C4 := 24 // C4 = 1:1 = 24:24 (unison)
	D4 := 27 // D4 = 9:8 = 27:24 (major whole tone)
	E4 := 30 // E4 = 5:4 = 30:24 (major third)
	F4 := 32 // F4 = 4:3 = 32:24 (perfect fourth)
	G4 := 36 // G4 = 3:2 = 36:24 (perfect fifth)
	A4 := 40 // A4 = 5:3 = 40:24 (major sixth)
	B4 := 45 // B4 = 15:8 = 45:24 (major seventh)
	C5 := 48 // C5 = 2:1 = 48:24 (octave)

	// Lower octave
	C3 := 12 // C3 = 1:2 = 12:24 (octave below)
	G3 := 18 // G3 = 3:4 = 18:24 (fifth below C4)

	// Note durations
	veryShortDur := 0.125 // 1/8 note
	shortDur := 0.25      // 1/4 note
	medDur := 0.5         // 1/2 note
	longDur := 1.0        // whole note

	fmt.Println("Generating intro...")

	// Arpeggiated intro - C major chord
	song.AddNote(C4, veryShortDur)
	song.AddNote(E4, veryShortDur)
	song.AddNote(G4, veryShortDur)
	song.AddNote(C5, veryShortDur)
	song.AddNote(G4, veryShortDur)
	song.AddNote(E4, veryShortDur)
	song.AddNote(C4, veryShortDur)
	song.AddNote(E4, veryShortDur)

	fmt.Println("Generating main melody...")

	// Main melody - staying within 5-odd limit
	song.AddNote(C4, shortDur)
	song.AddNote(D4, shortDur)
	song.AddNote(E4, medDur)
	song.AddNote(C4, shortDur)
	song.AddNote(E4, shortDur)
	song.AddNote(G4, medDur)
	song.AddNote(E4, shortDur) // Using E4 instead of F#4
	song.AddNote(G4, shortDur)
	song.AddNote(A4, medDur)
	song.AddNote(G4, shortDur)
	song.AddNote(E4, shortDur)
	song.AddNote(C4, medDur)

	// Descending line with tension and release
	song.AddNote(A4, shortDur)
	song.AddNote(G4, shortDur)
	song.AddNote(F4, shortDur)
	song.AddNote(E4, shortDur)
	song.AddNote(D4, shortDur)
	song.AddNote(C4, medDur)
	song.AddNote(G3, shortDur)
	song.AddNote(C4, longDur)

	fmt.Println("Adding harmonic progression...")

	// Chord progression with bass movement
	// I chord - C major (4:5:6 ratio)
	song.AddChord([]int{C4, E4, G4, C3}, medDur) // C major (24:30:36:12)

	// IV chord - F major (4:5:6 ratio transposed)
	song.AddChord([]int{F4, A4, C5}, medDur) // F major (32:40:48)

	// V chord - G major with 7th (4:5:6:7 ratio)
	song.AddChord([]int{G4, B4, D4, F4}, medDur) // G7 (36:45:27:32)

	// vi chord - A minor (10:12:15 ratio)
	song.AddChord([]int{A4, C5, E4}, medDur) // A minor (40:48:30)

	// ii-V-I cadence - all within 5-odd limit
	song.AddChord([]int{D4, F4, A4}, shortDur)    // D minor (27:32:40)
	song.AddChord([]int{G4, B4, D4}, shortDur)    // G major (36:45:27)
	song.AddChord([]int{C4, E4, G4, C5}, longDur) // C major (24:30:36:48)

	fmt.Println("Adding final section...")

	// Final arpeggiated section - descending C major
	song.AddNote(C5, veryShortDur)
	song.AddNote(G4, veryShortDur)
	song.AddNote(E4, veryShortDur)
	song.AddNote(C4, veryShortDur)
	song.AddNote(G3, veryShortDur)
	song.AddNote(C4, veryShortDur)
	song.AddNote(E4, veryShortDur)
	song.AddNote(G4, veryShortDur)

	// Final chord - C major with octave extensions
	song.AddChord([]int{C3, G3, C4, E4, G4, C5}, longDur*1.5) // C major (12:18:24:30:36:48)

	// Play the song
	return song.Play()
}

// NewJustIntonationComposition creates a new just intonation composition
func NewJustIntonationComposition() Composition {
	return &JustIntonationComposition{
		Name:        "just_intonation",
		Description: "A composition using 5-odd limit just intonation",
	}
}

func init() {
	// Register this composition with the registry
	RegisterComposition("just_intonation", NewJustIntonationComposition)
}
