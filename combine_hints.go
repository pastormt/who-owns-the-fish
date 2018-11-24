package main

// don't need to check that != since none of the hint values were listed > once
func canCombine2Hints(h1, h2 *Hint) bool {
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
	return true
}

// otherHints are already combined properly!
func canCombineManyHints(hint *Hint, otherHints []*Hint) bool {
	for _, otherHint := range otherHints {
		if !canCombine2Hints(hint, otherHint) {
			return false
		}
	}
	return true
}

func addHintDataToHouse(hint *Hint, house *House) {
	if isKnown(hint.Who) {
		house.Who = hint.Who
	}
	if isKnown(hint.Pet) {
		house.Pet = hint.Pet
	}
	if isKnown(hint.Sport) {
		house.Sport = hint.Sport
	}
	if isKnown(hint.Color) {
		house.Color = hint.Color
	}
	if isKnown(hint.Beverage) {
		house.Beverage = hint.Beverage
	}
}

func removeHintDataFromHouse(hint *Hint, house *House) {
	if isKnown(hint.Who) {
		house.Who = ""
	}
	if isKnown(hint.Pet) {
		house.Pet = ""
	}
	if isKnown(hint.Sport) {
		house.Sport = ""
	}
	if isKnown(hint.Color) {
		house.Color = ""
	}
	if isKnown(hint.Beverage) {
		house.Beverage = ""
	}
}

func addHintToHouseIfCan(hint *Hint, hintPos int, house *House) bool {
	if canCombineManyHints(hint, house.Hints) {
		house.Hints = append(house.Hints, hint)
		addHintDataToHouse(hint, house)
		// mark that hint was added
		hintsInHouses[hint] = hintPos
		return true
	}
	return false
}

// if the house is empty, unplace the last house that was completed and remove the last hint from it
// else, remove the last hint from house
func removeLastHint(house *House) (*House, int) {
	removeFrom := house
	if house.isEmpty() {
		removeFrom = unplaceLastHouse()
	}

	// the hint to remove
	lastHintPos := len(removeFrom.Hints) - 1
	removedHint := removeFrom.Hints[lastHintPos]

	// remove it!
	removeFrom.Hints = removeFrom.Hints[:lastHintPos]
	removeHintDataFromHouse(removedHint, removeFrom)

	hintsPos := hintsInHouses[removedHint]
	// mark that hint is available
	delete(hintsInHouses, removedHint)

	return removeFrom, hintsPos
}
