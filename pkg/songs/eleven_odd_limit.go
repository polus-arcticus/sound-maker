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
	eighth := 0.125 // 1/8 note
	quarter := 0.25 // 1/4 note
	half := 0.5     // 1/2 note
	whole := 1.0    // 1/1 note

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
	/*
		song.AddNote(unison, eighth)
		song.AddNote(secondTone, eighth)
		song.AddNote(thirdTone, eighth)
		song.AddNote(forthTone, eighth)
		song.AddNote(fifthTone, eighth)
		song.AddNote(sixthTone, eighth)
		song.AddNote(seventhTone, eighth)
		song.AddNote(eighthTone, eighth)
		song.AddNote(ninthTone, eighth)
		song.AddNote(tenthTone, eighth)
		song.AddNote(eleventhTone, eighth)
		song.AddNote(twelfthTone, eighth)
		song.AddNote(thirteenthTone, eighth)
		song.AddNote(fourteenthTone, eighth)
		song.AddNote(fifteenthTone, eighth)
		song.AddNote(sixteenthTone, eighth)
		song.AddNote(seventeenthTone, eighth)
		song.AddNote(eighteenthTone, eighth)
		song.AddNote(nineteenthTone, eighth)
		song.AddNote(twentiethTone, eighth)
		song.AddNote(twentyfirstTone, eighth)
		song.AddNote(twentysecondTone, eighth)
		song.AddNote(twentythirdTone, eighth)
		song.AddNote(twentyfourthTone, eighth)
		song.AddNote(twentyfifthTone, eighth)
		song.AddNote(twentysixthTone, eighth)
		song.AddNote(twentyseventhTone, eighth)
		song.AddNote(twentyeighthTone, eighth)
		song.AddNote(twentyninthTone, eighth)
		song.AddNote(thirtiethTone, eighth)

		song.AddNote(thirtiethTone, eighth)
		song.AddNote(twentyninthTone, eighth)
		song.AddNote(twentyeighthTone, eighth)
		song.AddNote(twentyseventhTone, eighth)
		song.AddNote(twentysixthTone, eighth)
		song.AddNote(twentyfifthTone, eighth)
		song.AddNote(twentyfourthTone, eighth)
		song.AddNote(twentythirdTone, eighth)
		song.AddNote(twentysecondTone, eighth)
		song.AddNote(twentyfirstTone, eighth)
		song.AddNote(twentiethTone, eighth)
		song.AddNote(nineteenthTone, eighth)
		song.AddNote(eighteenthTone, eighth)
		song.AddNote(seventeenthTone, eighth)
		song.AddNote(sixteenthTone, eighth)
		song.AddNote(fifteenthTone, eighth)
		song.AddNote(fourteenthTone, eighth)
		song.AddNote(thirteenthTone, eighth)
		song.AddNote(twelfthTone, eighth)
		song.AddNote(eleventhTone, eighth)
		song.AddNote(tenthTone, eighth)
		song.AddNote(ninthTone, eighth)
		song.AddNote(eighthTone, eighth)
		song.AddNote(seventhTone, eighth)
		song.AddNote(sixthTone, eighth)
		song.AddNote(fifthTone, eighth)
		song.AddNote(forthTone, eighth)
		song.AddNote(thirdTone, eighth)
		song.AddNote(secondTone, eighth)
		song.AddNote(unison, eighth)
	*/
	// 5-limit major triad (pure just intonation)
	// 4:5:6 = 1/1 : 5/4 : 3/2
	song.AddChord([]int{unison, fifthTone, sixthTone}, eighth)
	song.AddChord([]int{unison, tenthTone, eighteenthTone}, quarter)
	song.AddChord([]int{unison, tenthTone, eighteenthTone, twentyfifthTone}, quarter)
	song.AddChord([]int{unison, tenthTone, eighteenthTone, twentyfifthTone}, half)
	song.AddChord([]int{unison, tenthTone, eighteenthTone}, whole)

	//Otonality
	// 11-odd limit otonality scale: ratios 8:9:10:11:12 (harmonic series with same denominator)
	// 8/8 (1/1), 9/8, 10/8 (5/4), 11/8, 12/8 (3/2)
	fmt.Println("Playing 11-odd limit otonality scale...")
	song.AddNote(unison, quarter)         // 8/8 = 1/1
	song.AddNote(fifthTone, quarter)      // 9/8
	song.AddNote(tenthTone, quarter)      // 10/8 = 5/4
	song.AddNote(fourteenthTone, quarter) // 11/8
	song.AddNote(eighteenthTone, quarter) // 12/8 = 3/2

	//Utonality
	// 11-odd limit utonality scale: ratios with 11 in numerator (subharmonic series)
	// 11/11 (1/1), 11/10, 11/9, 11/8, 11/7, 11/6
	fmt.Println("Playing 11-odd limit utonality scale...")
	song.AddNote(unison, quarter)          // 11/11 = 1/1
	song.AddNote(thirdTone, quarter)       // 11/10
	song.AddNote(ninthTone, quarter)       // 11/9
	song.AddNote(fourteenthTone, quarter)  // 11/8
	song.AddNote(twentiethTone, quarter)   // 11/7
	song.AddNote(twentyninthTone, quarter) // 11/6

	// 7-limit dominant seventh (septimal)
	// 4:5:6:7 = 1/1 : 5/4 : 3/2 : 7/4

	/*
		// 5-limit minor triad
		// 10:12:15 = 1/1 : 6/5 : 3/2
		song.AddChord([]int{unison, eighthTone, eighteenthTone}, quarter)


		// 7-limit minor with septimal seventh
		// 10:12:15:18 approximation, but using 7/4 on top
		song.AddChord([]int{unison, eighthTone, eighteenthTone, twentyfifthTone}, quarter)
			// 11-limit "alpha chord" - features the characteristic 11/8
			// 8:10:11:12 = 1/1 : 5/4 : 11/8 : 3/2
			song.AddChord([]int{unison, tenthTone, fourteenthTone, eighteenthTone}, quarter)

			// Suspended 11 flavor
			// Using 1/1 : 4/3 : 11/8 : 3/2
			song.AddChord([]int{unison, thirteenthTone, fourteenthTone, eighteenthTone}, quarter)

			// 11-limit minor with undecimal color
			// 1/1 : 6/5 : 11/8 : 3/2
			song.AddChord([]int{unison, eighthTone, fourteenthTone, eighteenthTone}, quarter)

			// Septimal tritone chord (very tense!)
			// 1/1 : 7/5 : 10/7 (sqrt of 2 approximation area)
			song.AddChord([]int{unison, fifteenthTone, sixteenthTone}, quarter)

			// 11-limit extended harmony pentad
			// 8:10:12:14:11 rearranged = 1/1 : 5/4 : 11/8 : 3/2 : 7/4
			song.AddChord([]int{unison, tenthTone, fourteenthTone, eighteenthTone, twentyfifthTone}, quarter)

			// Quartal stack with 11-limit color
			// 1/1 : 4/3 : 16/11 (approximates quartal but with 11-limit flavor)
			song.AddChord([]int{unison, thirteenthTone, seventeenthTone}, quarter)
	*/

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
