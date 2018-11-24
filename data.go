package main

import "fmt"

type Hint struct {
	Who, Pet, Sport, Color, Beverage string
	Position                         *int
	ToRight, ToLeft                  *Hint
	NextTo                           []*Hint
}

type House struct {
	Who, Pet, Sport, Color, Beverage string
	Hints                            []*Hint
}

func (h House) isComplete() bool {
	return h.isCompleteMinusPet() && isKnown(h.Pet)
}

func (h House) isCompleteMinusPet() bool {
	return isKnown(h.Who) && isKnown(h.Sport) && isKnown(h.Color) && isKnown(h.Beverage)
}

func (h House) String() string {
	return fmt.Sprintf("Who: %s\nPet: %s\nSport: %s\nColor: %s\nBeverage: %s\n\n",
		h.Who, h.Pet, h.Sport, h.Color, h.Beverage)
}

func (h House) isEmpty() bool {
	return len(h.Hints) == 0
}

func isKnown(s string) bool {
	return s != ""
}

func ptrToInt(i int) *int {
	return &i
}

// yes, everything except the fish is mentioned in a hint ;)
// Brit, Swede, Dane, Norwegian, German
// red, white, green, yellow, blue
// dogs, birds, cats, horses, FISH
// tea, coffee, milk, water, beer
// polo, hockey, soccer, baseball, billiards

var (
	hint1  = Hint{Who: "Brit", Color: "red"}
	hint2  = Hint{Who: "Swede", Pet: "dogs"}
	hint3  = Hint{Who: "Dane", Beverage: "tea"}
	hint4  = Hint{Color: "white"}
	hint5  = Hint{Color: "green", Beverage: "coffee"}
	hint6  = Hint{Pet: "birds", Sport: "polo"}
	hint7  = Hint{Color: "yellow", Sport: "hockey"}
	hint8  = Hint{Sport: "baseball"}
	hint9  = Hint{Pet: "cats"}
	hint10 = Hint{Pet: "horses"}
	hint11 = Hint{Beverage: "beer", Sport: "billiards"}
	hint12 = Hint{Sport: "soccer", Who: "German"}
	hint13 = Hint{Color: "blue", Position: ptrToInt(1)}
	hint14 = Hint{Who: "Norwegian", Position: ptrToInt(0)} // first house!
	hint15 = Hint{Beverage: "water"}
	hint16 = Hint{Beverage: "milk", Position: ptrToInt(2)} // middle, 01234
)

func init() {
	// define relationships here to avoid type-checking loop
	hint4.ToLeft = &hint5
	hint4.NextTo = []*Hint{&hint5}

	hint5.ToRight = &hint4
	hint5.NextTo = []*Hint{&hint4}

	hint7.NextTo = []*Hint{&hint10}
	hint10.NextTo = []*Hint{&hint7}

	hint8.NextTo = []*Hint{&hint9, &hint15}
	hint9.NextTo = []*Hint{&hint8}
	hint15.NextTo = []*Hint{&hint8}

	hint13.ToRight = &hint16
	hint13.ToLeft = &hint14
	hint13.NextTo = []*Hint{&hint14, &hint16}

	hint14.ToRight = &hint13
	hint14.NextTo = []*Hint{&hint13}

	hint16.ToLeft = &hint13
	hint16.NextTo = []*Hint{&hint13}
}

var hints = []*Hint{
	&hint1, &hint2, &hint3, &hint4, &hint5, &hint6, &hint7, &hint8, &hint9,
	&hint10, &hint11, &hint12, &hint13, &hint14, &hint15, &hint16,
}
