package main

import (
	"flag"
	"fmt"
	"os"
	"sound/pkg/songs"
	"strings"
)

func main() {
	// Define command line flags
	listFlag := flag.Bool("list", false, "List all available compositions")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] [composition_name]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nAvailable compositions:\n")
		for _, name := range songs.ListCompositions() {
			comp, _ := songs.GetComposition(name)
			fmt.Fprintf(os.Stderr, "  %s - %s\n", name, comp.GetDescription())
		}
	}

	// Parse command line flags
	flag.Parse()

	// If list flag is set, list all compositions and exit
	if *listFlag {
		fmt.Println("Available compositions:")
		for _, name := range songs.ListCompositions() {
			comp, _ := songs.GetComposition(name)
			fmt.Printf("  %s - %s\n", name, comp.GetDescription())
		}
		return
	}

	// Get the composition name from command line arguments
	var compositionName string
	if flag.NArg() > 0 {
		compositionName = flag.Arg(0)
	} else {
		// Default to just_intonation if no composition name is provided
		compositionName = "just_intonation"
	}

	// Get the composition
	composition, err := songs.GetComposition(compositionName)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		fmt.Println("\nAvailable compositions:")
		for _, name := range songs.ListCompositions() {
			comp, _ := songs.GetComposition(name)
			fmt.Printf("  %s - %s\n", name, comp.GetDescription())
		}
		os.Exit(1)
	}

	// Print a header
	fmt.Printf("Playing composition: %s\n", composition.GetName())
	fmt.Printf("Description: %s\n", composition.GetDescription())
	fmt.Println(strings.Repeat("-", 50))

	// Play the composition
	err = composition.Play()
	if err != nil {
		fmt.Printf("Error playing composition: %v\n", err)
		os.Exit(1)
	}
}
