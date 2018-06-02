package main

type house struct {
	Who, Pet, Sport, Color, Beverage string
  Position uint
  ToRight, ToLeft *house
  NextTo []*house
}

var winners [][]*house
// set of sets
// each slice element in the outer slice is a winning set,
// meaning the houses at those positions combine to form a complete house

// yes, everything except the fish is mentioned in a hint ;)
// Brit, Swede, Dane, Norwegian, German
// red, white, green, yellow, blue
// dogs, birds, cats, horses, FISH
// tea, coffee, milk, water, beer
// polo, hockey, soccer, baseball, billiards

var originalHouses []*house = useHints()

func useHints() []*house {
	o1 := house{Who: "Brit", Color: "red"}
	o2 := house{Who: "Swede", Pet: "dogs"}
	o3 := house{Who: "Dane", Beverage: "tea"}
	o4 := house{Color: "white", ToRight: &o5, NextTo: []*house{&o5}}
	o5 := house{Color: "green", Beverage: "coffee": ToLeft: &o4, NextTo: []*house{&o4}}
	o6 := house{Pet: "birds", Sport: "polo"}
	o7 := house{Color: "yellow", Sport: "hockey", NextTo: []*house{&o10}}
	o8 := house{Sport: "baseball", NextTo: []*house{&o9, &o15}}
	o9 := house{Pet: "cats", NextTo: []*house{&o8}}
	o10 := house{Pet: "horses", NextTo: []*house{&o7}}
	o11 := house{Beverage: "beer", Sport: "billiards"}
	o12 := house{Sport: "soccer", Who: "German"}
	o13 := house{Color: "blue", NextTo: []*house{&o14}}
	o14 := house{Who: "Norwegian", Position: 1, NextTo: []*house{&o13}} // first house!
	o15 := house{Beverage: "water", NextTo: []*house{&o8}}
  o16 := house{Beverage: "milk", Position: 3} // middle, 12345, index from 1 so that unknown (default) doesnt confuse with the first pos (0)
  return []*house{&o1,&o2,&o3,&o4,&o5,&o6,&o7,&o8,&o9,&o10,&o11,&o12,&o13,&o14,&o15,&o16}
}

func isKnown(s string) bool {
  return s != ""
}

// if they're next to each other, can't be the same house
func isNextTo(h1, h2 *house) bool {
  // don't check both houses, since data is in each
  for _, house := range h1.NextTo {
		if house == h2 {
			return true
		}
  }
  return false
}

// don't think need to check && h1.Blah != h2.Blah,
// since they start as one copy of each value,
// and won't be copying a value
func canCombine2Houses(h1, h2 house) bool {
  if isKnown(h1.Who) && isKnown(h2.Who) {
    return false
  }
  if isKnown(h1.Pet) && isKnown(h2.Pet) {
    return false
  }
  if isKnown(h1.Sport) && isKnown(h2.Sport) {
    return false
  }
  if isKnown(h1.Color) && isKnown(h2.Color) {
    return false
  }
  if isKnown(h1.Beverage) && isKnown(h2.Beverage) {
    return false
  }
  if isKnown(h1.Position) && isKnown(h2.Position) {
    return false
  }
	if isNextTo(h1,h2) {
		return false
	}
	// if it's known what's to the right of both, and it's not the same thing
	// if it's a not-yet-complete house that each points to,
	// then they don't have to point to the same thing,
	// since multiple pieces of houses will be combined to make a full house
	if h1.ToRight != nil && h2.ToRight != nil {
		if isWinner(h1.ToRight) != isWinner(h2.ToRight) {
			return false
		}
	}
	if h1.ToLeft != nil && h2.ToLeft != nil {
		if isWinner(h1.ToLeft) != isWinner(h2.ToLeft) {
			return false
		}
	}
  return true
}

func isWinner(h *house) *[]*house {
	for i := range winners {
		for j := range winners[i] {
			if h == winners[i][j] {
				return &winners[i]
			}
		}
	}
	return nil
}

func canCombineManyHouses(hs []*house, h2 house) bool {
	for _, house := range hs {
		if !canCombine2Houses(house, h2) {
			return false
		}
	}

	// TODO check that between all of hs and h2, only 2 distinct COMBINED houses are pointed to by the next to slices
	// eg if those in hs already point to 2 different winners, if h2 has a next to that doesnt point to one of those winners,
	// then cant combine
	// but if hs points to 1 winner, and another in hs points to a different house and h2 points to yet a different house,
	// and long as neither of the latter 2 are winners, no problem, since dont yet know that they wont be part of a combined house in the future
	// check how many winners point to. can be <= 2 at most, and then point to no others things.
	for _, house := range append(hs, h2) {

	}
	return true
}

func isComplete(hs []house) bool {
	// assumes no duplicates (ie only houses that could be combined, were so)
	var known int
	var pet string
	for _, house := range hs {
		if isKnown(house.Who) {
			known++
		}
		if isKnown(house.Beverage) {
			known++
		}
		if isKnown(house.Pet) {
			pet = house.Pet
			known++
		}
		if isKnown(house.Sport) {
			house++
		}
		if isKnown(house.Color) {
			house++
		}
	}
	return known == 5 || (known == 4 && pet == "") // room for fish
}

// since one set of houses will be complete without a pet, and that pet is implied to be the fish
func whoOwnsTheFish(hs [][]*house) (who string) {
	var thisCombinedHouseHasTheFish []house
	for _, combinedHouse := range hs {
		var pet string
		for _, origHouse := range combinedHouse {
			if isKnown(origHouse.Pet) {
				pet = origHouse.Pet
			}
		}
		if pet == "" {
			// this combined house owns the fish!!
			for _, origHouse := range combinedHouse {
				if isKnown(origHouse.Who) {
					return origHouse.Who
				}
			}
		}
		return ""
}

// dont necessarily have all 5 winners here
for canYouOrderWinnersOK(winners [][]*house, newWinner []*house) bool {
	// the Norwegian has to be in a first house
	// all with a position have to be in the position
	// left, right, next to have to be fulfilled ;-)

	var potentialWinners = append(winners, newWinner)
	var orderedWinners = make([][]*house, 5)

	var placedOK = make(map[int]bool)
	for i := len(potentialWinners) {
		placedOK[i] = false
	}

	// those with a position first, because that's hard requirement
	for i, winner := range potentialWinners {
		var pos int // do any of the houses in this winning combo have a position set?
		for _, h := range winner {
			if h.Position != 0 {
				pos = h.Position
				placedOK[i] = true // definitely placed
				break
			}
		}
		if pos != 0 {
			// put it in that pos - 1
			orderedWinners[pos-1] = winner
		}
	}

	// left, right, next to
	for i, winner := range potentialWinners {
		if _, alreadyPlaced := placedOK[i]; !alreadyPlaced {

		}
	}
}

// TODO make sure you dont try the same wrong combinations over and over
// mark them as already tried or by order they are tried, eg if i and you only try to combine with things to the right of 1
func solvePuzzle(current, remaining []*house) (who string) {
	if len(current) == 0 && len(winners) > 0 {
		redistributeHouses := winners[len(winners)-1)
		winners = winners[:len(winners)-1]
		return solvePuzzle([]*house{}, append(remaining, redistributeHouses...))
	}

	if isComplete(current) {
		winners = append(winners, current)
		// see if you can arrange winners in a satisfying order
		if canYouOrderWinnersOK(winners, current) {
			winners = append(winners, current)
			current = []*house{}
			if len(remaining) > 0 {
				return solvePuzzle(current, remaining)
			}
		} else {
			return ""
		}
		// if not, remove the winner you just added, and redistribute houses, and (don't get the same winner again!)

		// else, you've found the right answer!
		// do something to celebrate
		return whoOwnsTheFish(winners)
	}

	// current is incomplete. can you complete it?
	for i, h := range remaining {
		if canCombineManyHouses(current, h) {
			// add h to current, and subtract it from remaining
			newCurrent := append(current, h)
			newRemaining := append(remaining[:i], remaining[i+1:]...)
			sol := solvePuzzle(newCurrent, newRemaining)
			if sol != "" {
				return sol
			}
		}
	}
	// if here, then couldn't complete current
	// returning will get you back to where the last thing added to current is removed
	// or if current is empty, then to the last "winner" redistributed to remaining
	return ""
}

func main() {
  solvePuzzle([]*house{}, useHints())
}
