package main

import (
	"fmt"
)

var (
	spots         [5]*House             // are the ordered locations with completed houses placed in them
	placedHouses  []*House              // are the completed houses in order of least to most recently placed
	hintsInHouses = make(map[*Hint]int) // hints that are already in houses... int is the pos of the hint in the Hints slice!
)

func main() {
	who, found := whoOwnsTheFish(&House{}, 0)
	if !found {
		panic("couldn't solve the puzzle!")
	}

	fmt.Printf("\n\n%s owns the fish! :-)\n\n\n", who)

	for i, house := range spots {
		fmt.Printf("Spot %d\n", i)
		fmt.Printf("%+v\n", *house)
	}
}

func whoOwnsTheFish(house *House, startAt int) (string, bool) {
	if spotsAreFull(spots) { // wooo, done!
		fishsHouse := findHouseMissingAPet()
		return fishsHouse.Who, true
	}

	for i := startAt; i < len(hints); i++ {
		hint := hints[i]
		if _, hintIsAlreadyUsed := hintsInHouses[hint]; hintIsAlreadyUsed {
			continue
		}

		if !addHintToHouseIfCan(hint, i, house) {
			continue
		}

		// did the hint complete the house?
		if house.isComplete() {
			return tryPlaceHouseAndProceed(house)
		}
	}

	// tried all hints but could not produce a complete house
	// if complete enough (minus only a pet), and no other placed houses are also missing a pet, then proceed
	// else, remove last hint
	if house.isCompleteMinusPet() && spotsHaveOnlyCompleteHouses() {
		return tryPlaceHouseAndProceed(house)
	}

	// alas -- need to backtrack
	houseWithoutHint, hintsPos := removeLastHint(house)
	return whoOwnsTheFish(houseWithoutHint, hintsPos+1)
}
