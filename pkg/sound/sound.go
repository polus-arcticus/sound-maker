// Package sound provides functionality for generating and playing audio
package sound

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"
	"os/exec"
)

const Tau = 2 * math.Pi

// Song represents the internal state of a song
type Song struct {
	file          *os.File
	BaseFrequency int
	BaseNote      int
	SampleRate    float64
	TotalDuration float64
}

// SongImpl implements the Song interface
type SongImpl struct {
	Song
}

// ISong defines the interface for song operations
type ISong interface {
	Play() error
	Write(buf []byte) (int, error)
	Close() error
	AddNote(frequency int, duration float64)
	AddChord(notes []int, duration float64)
}

// Play plays the generated sound file
func (s *SongImpl) Play() error {
	fmt.Println("Playing song...")
	cmd := exec.Command("aplay", "-f", "FLOAT_LE", "-r", fmt.Sprintf("%.0f", s.SampleRate), "-c", "1", s.file.Name())
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error playing sound: %w", err)
	}
	fmt.Println("Sound played successfully")
	return nil
}

// Write writes raw bytes to the song file
func (s *SongImpl) Write(buf []byte) (int, error) {
	return s.file.Write(buf)
}

// Close closes the song file
func (s *SongImpl) Close() error {
	return s.file.Close()
}

// AddNote adds a single note to the song
func (s *SongImpl) AddNote(note int, duration float64) {
	s.TotalDuration += duration
	var numberSamples float64 = duration * s.SampleRate
	var angle float64 = Tau / numberSamples

	for i := 0; i < int(numberSamples); i++ {
		sample := math.Sin(angle * float64(s.BaseFrequency*note/s.BaseNote) * float64(i))
		var buf [4]byte
		binary.LittleEndian.PutUint32(buf[:], math.Float32bits(float32(sample)))
		_, err := s.file.Write(buf[:])
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
}

// sum adds up all values in a float64 slice
func sum(values []float64) float64 {
	var total float64
	for _, v := range values {
		total += v
	}
	// Normalize to avoid clipping
	if len(values) > 0 {
		total /= float64(len(values))
	}
	return total
}

// AddChord adds multiple notes played simultaneously to create a chord
func (s *SongImpl) AddChord(notes []int, duration float64) {
	s.TotalDuration += duration
	var numberSamples float64 = duration * s.SampleRate
	var angle float64 = Tau / numberSamples

	for i := 0; i < int(numberSamples); i++ {
		accumulator := make([]float64, 0, len(notes)) // Initialize with 0 length but capacity of len(notes)
		for _, note := range notes {
			accumulator = append(accumulator, math.Sin(angle*float64(s.BaseFrequency*note/s.BaseNote)*float64(i)))
		}
		var buf [4]byte
		binary.LittleEndian.PutUint32(buf[:], math.Float32bits(float32(sum(accumulator))))
		_, err := s.file.Write(buf[:])
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
}

// NewSong creates a new song with the specified parameters
func NewSong(fileName string, sampleRate float64, baseFrequency int, baseNote int) ISong {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return nil
	}
	return &SongImpl{Song: Song{
		file:          file,
		SampleRate:    sampleRate,
		TotalDuration: 0,
		BaseFrequency: baseFrequency,
		BaseNote:      baseNote,
	}}
}

/*
   fmt.Println("Hello, World!")

    song := NewSong("out.bin", 36000.0)
    defer song.Close()

    song.AddNote(360, 1.0) // C
    song.AddNote(405, 1.0) // D
    song.AddNote(450, 1.0) // E
    song.AddNote(480, 1.0) // F
    song.AddNote(540, 1.0) // G
    song.AddNote(600, 1.0) // A
    song.AddNote(675, 1.0) // B
    song.AddNote(720, 1.0) // C2

    song.AddChord([]int{360, 450, 540}, 1.0) // C E G
    song.AddChord([]int{480, 600, 360}, 1.0) // F A C
    song.AddChord([]int{540, 675, 405}, 1.0) // G B D
    //song.AddChord([]int{660, 440}, 1.0)
    //song.AddChord([]int{440, 220}, 1.0)
    song.Play()
*/
