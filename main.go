package main

import (
	"fmt"
	"sound/pkg/sound"
)

// Define 5-odd limit just intonation frequencies
// These are based on integer ratios rather than equal temperament
// For example, a perfect fifth is 3:2 ratio, not the 2^(7/12) of equal temperament

func main() {
	fmt.Println("Creating a 5-odd limit just intonation song")

	// Create a new song with 44.1kHz sample rate
	// BaseFrequency = 264 Hz (C4)
	// BaseNote = 24 (the denominator for our ratios)
	song := sound.NewSong("just_intonation.bin", 44100.0, 264, 24)
	defer song.Close()

	// Define our notes using integer ratios
	// With the new system, we just specify the ratio numerator
	// The denominator is fixed at 24 (BaseNote)

	// 5-odd limit just intonation ratios (integer form)
	// C:D:E:F:G:A:B:C = 24:27:30:32:36:40:45:48

	C4 := 24 // C4 = 24/24 * 264 Hz = 264 Hz
	D4 := 27 // D4 = 27/24 * 264 Hz = 297 Hz
	E4 := 30 // E4 = 30/24 * 264 Hz = 330 Hz
	F4 := 32 // F4 = 32/24 * 264 Hz = 352 Hz
	G4 := 36 // G4 = 36/24 * 264 Hz = 396 Hz
	A4 := 40 // A4 = 40/24 * 264 Hz = 440 Hz
	B4 := 45 // B4 = 45/24 * 264 Hz = 495 Hz
	C5 := 48 // C5 = 48/24 * 264 Hz = 528 Hz

	// Short duration for faster notes
	shortDur := 0.25
	medDur := 0.5
	longDur := 1.0

	fmt.Println("Generating melody...")

	// Simple melody in just intonation
	song.AddNote(C4, shortDur)
	song.AddNote(E4, shortDur)
	song.AddNote(G4, medDur)
	song.AddNote(C5, shortDur)
	song.AddNote(B4, shortDur)
	song.AddNote(A4, medDur)
	song.AddNote(G4, shortDur)
	song.AddNote(E4, shortDur)
	song.AddNote(C4, longDur)

	// Add some chords
	fmt.Println("Adding chords...")

	// Major triad: 4:5:6 ratio (24:30:36)
	song.AddChord([]int{C4, E4, G4}, longDur) // C major (24:30:36)

	// Minor triad: 10:12:15 ratio (40:48:60 → 40:48:E4)
	song.AddChord([]int{A4, C5, E4}, longDur) // A minor (40:48:30)

	// Dominant seventh: 4:5:6:7 ratio (24:30:36:42 → G:B:D:F)
	song.AddChord([]int{G4, B4, D4, F4}, longDur) // G7 (36:45:27:32)

	// Final chord
	song.AddChord([]int{C4, E4, G4, C5}, longDur*1.5) // C major with octave (24:30:36:48)

	// Play the song
	err := song.Play()
	if err != nil {
		fmt.Println("Error playing song:", err)
	}
}
