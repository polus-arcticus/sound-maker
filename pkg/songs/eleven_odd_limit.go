package songs

import (
	"fmt"
	"sound/pkg/sound"
)

type ElevenOddLimitComposition struct {
	Name        string
	Description string
}

func (c *ElevenOddLimitComposition) GetName() string {
	return c.Name
}

func (c *ElevenOddLimitComposition) GetDescription() string {
	return c.Description
}

func (c *ElevenOddLimitComposition) Play() error {
	fmt.Println("Creating an eleven odd limit composition")

	// Create a new song with 44.1kHz sample rate
	song := sound.NewSong("eleven_odd_limit.bin", 44100.0, 264, 2310)
	defer song.Close()

	// the eleven odd limit
	/*
		11-limit tonality diamond
		12/11
		11/10
		10/9
		9/8
		8/7
		7/6
		6/5
		11/9
		5/4
		14/11
		9/7
		4/3
		11/8
		7/5
		10/7
		16/11
		3/2
		14/9
		11/7
		8/5
		18/11
		5/3
		12/7
		7/4
		16/9
		9/5
		20/11
		11/6
		2
	*/

	shortDur := 0.25 // 1/4 note

	unison := 2310 // unison = 1:1 = 2310:2310 (unison)
	fmt.Printf("unison: %d\n", unison)
	secondTone := unison * 12 / 11 //  second = 12/11
	fmt.Printf("secondTone: %d\n", secondTone)
	thirdTone := unison * 11 / 10 //  third = 11/10
	fmt.Printf("thirdTone: %d\n", thirdTone)
	forthTone := unison * 10 / 9 //  fourth = 10/9
	fmt.Printf("forthTone: %d\n", forthTone)
	fifthTone := unison * 9 / 8 //  fifth = 9/8
	fmt.Printf("fifthTone: %d\n", fifthTone)
	sixthTone := unison * 8 / 7 //  sixth = 8/7
	fmt.Printf("sixthTone: %d\n", sixthTone)
	seventhTone := unison * 7 / 6 //  seventh = 7/6
	fmt.Printf("seventhTone: %d\n", seventhTone)
	eighthTone := unison * 6 / 5 //  eighth = 6/5
	fmt.Printf("eighthTone: %d\n", eighthTone)
	ninthTone := unison * 11 / 9 //  ninth = 11/9
	fmt.Printf("ninthTone: %d\n", ninthTone)
	tenthTone := unison * 5 / 4 //  tenth = 5/4
	fmt.Printf("tenthTone: %d\n", tenthTone)
	eleventhTone := unison * 14 / 11 //  eleventh = 14/11
	fmt.Printf("eleventhTone: %d\n", eleventhTone)
	twelfthTone := unison * 9 / 7 //  twelfth = 9/7
	fmt.Printf("twelfthTone: %d\n", twelfthTone)
	thirteenthTone := unison * 4 / 3 //  thirteenth = 4/3
	fmt.Printf("thirteenthTone: %d\n", thirteenthTone)
	fourteenthTone := unison * 11 / 8 //  fourteenth = 11/8
	fmt.Printf("fourteenthTone: %d\n", fourteenthTone)
	fifteenthTone := unison * 7 / 5 //  fifteenth = 7/5
	fmt.Printf("fifteenthTone: %d\n", fifteenthTone)
	sixteenthTone := unison * 10 / 7 //  sixteenth = 10/7
	fmt.Printf("sixteenthTone: %d\n", sixteenthTone)
	seventeenthTone := unison * 16 / 11 //  seventeenth = 16/11
	fmt.Printf("seventeenthTone: %d\n", seventeenthTone)
	eighteenthTone := unison * 3 / 2 //  eighteenth = 3/2
	fmt.Printf("eighteenthTone: %d\n", eighteenthTone)
	nineteenthTone := unison * 14 / 9 //  nineteenth = 14/9
	fmt.Printf("nineteenthTone: %d\n", nineteenthTone)
	twentiethTone := unison * 11 / 7 //  twentieth = 11/7
	fmt.Printf("twentiethTone: %d\n", twentiethTone)
	twentyfirstTone := unison * 8 / 5 //  twentyfirst = 8/5
	fmt.Printf("twentyfirstTone: %d\n", twentyfirstTone)
	twentysecondTone := unison * 18 / 11 //  twentysecond = 18/11
	fmt.Printf("twentysecondTone: %d\n", twentysecondTone)
	twentythirdTone := unison * 5 / 3 //  twentythird = 5/3
	fmt.Printf("twentythirdTone: %d\n", twentythirdTone)
	twentyfourthTone := unison * 12 / 7 //  twentyfourth = 12/7
	fmt.Printf("twentyfourthTone: %d\n", twentyfourthTone)
	twentyfifthTone := unison * 7 / 4 //  twentyfifth = 7/4
	fmt.Printf("twentyfifthTone: %d\n", twentyfifthTone)
	twentysixthTone := unison * 16 / 9 //  twentysixth = 16/9
	fmt.Printf("twentysixthTone: %d\n", twentysixthTone)
	twentyseventhTone := unison * 9 / 5 //  twentyseventh = 9/5
	fmt.Printf("twentyseventhTone: %d\n", twentyseventhTone)
	twentyeighthTone := unison * 20 / 11 //  twentyeighth = 20/11
	fmt.Printf("twentyeighthTone: %d\n", twentyeighthTone)
	twentyninthTone := unison * 11 / 6 //  twentyninth = 11/6
	fmt.Printf("twentyninthTone: %d\n", twentyninthTone)
	thirtiethTone := unison * 2 //  thirtieth = 2 = 2310 * 2 = 4620
	fmt.Printf("thirtiethTone: %d\n", thirtiethTone)

	song.AddNote(unison, shortDur)
	song.AddNote(secondTone, shortDur)
	song.AddNote(thirdTone, shortDur)
	song.AddNote(forthTone, shortDur)
	song.AddNote(fifthTone, shortDur)
	song.AddNote(sixthTone, shortDur)
	song.AddNote(seventhTone, shortDur)
	song.AddNote(eighthTone, shortDur)
	song.AddNote(ninthTone, shortDur)
	song.AddNote(tenthTone, shortDur)
	song.AddNote(eleventhTone, shortDur)
	song.AddNote(twelfthTone, shortDur)
	song.AddNote(thirteenthTone, shortDur)
	song.AddNote(fourteenthTone, shortDur)
	song.AddNote(fifteenthTone, shortDur)
	song.AddNote(sixteenthTone, shortDur)
	song.AddNote(seventeenthTone, shortDur)
	song.AddNote(eighteenthTone, shortDur)
	song.AddNote(nineteenthTone, shortDur)
	song.AddNote(twentiethTone, shortDur)
	song.AddNote(twentyfirstTone, shortDur)
	song.AddNote(twentysecondTone, shortDur)
	song.AddNote(twentythirdTone, shortDur)
	song.AddNote(twentyfourthTone, shortDur)
	song.AddNote(twentyfifthTone, shortDur)
	song.AddNote(twentysixthTone, shortDur)
	song.AddNote(twentyseventhTone, shortDur)
	song.AddNote(twentyeighthTone, shortDur)
	song.AddNote(twentyninthTone, shortDur)
	song.AddNote(thirtiethTone, shortDur)

	return song.Play()

}

func NewElevenOddLimitComposition() Composition {
	return &ElevenOddLimitComposition{
		Name:        "eleven_odd_limit",
		Description: "A composition using 11-odd limit",
	}
}

func init() {
	RegisterComposition("eleven_odd_limit", NewElevenOddLimitComposition)
}
