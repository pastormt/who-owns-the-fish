package main

func placeHouseIfCan(house *House, excludedPos map[int]bool) bool {

	var spotsThatSatisfy []int

	for i := range spots {
		if _, isExcluded := excludedPos[i]; isExcluded {
			continue
		}
		if !spotSatisfies(house, i, spots) {
			continue
		}
		spotsThatSatisfy = append(spotsThatSatisfy, i)
	}

	// trying good spots first
	// good ->
	// spot satisfies, is empty, placing there doesnt interfere with other placed houses
	for _, i := range spotsThatSatisfy {
		spot := spots[i]

		// putting a new house next to an existing house could unmeet the needs of the existing house, if say,
		// the new house doesn't need to be next to anything particular, but the existing house does.
		// in this scenario, if only the new house's spatial reqs were checked when it was placed, then they would be met,
		// but by placing the new house there, it could eliminate a spot that the existing house may need for its spatial reqs to be met.
		houseToRightStillSatisfied := adjacentHouseStillSatified(house, i, true)
		houseToLeftStillSatisfied := adjacentHouseStillSatified(house, i, false)

		if !houseToLeftStillSatisfied || !houseToRightStillSatisfied {
			continue
		}

		if spot != nil { // there's a house there already
			continue
		}

		// great, place!
		placeHouseAtPosition(house, i)
		return true
	}

	// trying not so good, but valid, spots next
	// not so good ->
	// spot satisfies, but either isn't empty or interferes with something else
	// in this case, if you tried all the good spots already, then if you
	// can place whatever other house is interfering here elsewhere, then do it!

	// if you have to move something A for the house B to be placed at spot 1 ->
	// in moving A, you will check that A's spatial reqs are still met when placeIfCan.
	// since the spatial relationships are put in both houses (eg if house D is left of house E,
	// then house E is right of house D), then if the adjacent house's spatial reqs are met when it is moved,
	// if there is a house C that was adjacent to this just-moved adjacent house A, if C needed something from A,
	// then A would also need that from C, so if A's spatial reqs are met in the new position, then so are C's! ;-)
	for _, i := range spotsThatSatisfy {
		spot := spots[i]

		if !adjacentHouseStillSatified(house, i, true) && !placeHouseElsewhereIfCan(i+1, excludedPos) {
			continue
		}

		if !adjacentHouseStillSatified(house, i, false) && !placeHouseElsewhereIfCan(i-1, excludedPos) {
			continue
		}

		if spot != nil && !placeHouseElsewhereIfCan(i, excludedPos) {
			continue
		}

		// great, place!
		placeHouseAtPosition(house, i)
		return true
	}

	// none of the good, or not so good, spots worked, so...
	return false
}

// adjacentHint is the hint that is adjacent to some root house, A. A is at rootHousePos.
func requiredAdjacentHintIsAdjacent(adjacentHint *Hint, rootHousePos int, isToRight bool, spots [5]*House) bool {
	adjacentPos := rootHousePos
	if isToRight {
		adjacentPos++
	} else {
		adjacentPos--
	}
	if adjacentPos < 0 || adjacentPos > 4 {
		// the house that should be to the right or left of root house can't be, since nothing can be
		return false
	}

	// is the hint that should be in the house to the right or left placed yet?
	if _, inHouse := hintsInHouses[adjacentHint]; !inHouse {
		// that hint is not placed -- there should be no house to the right / left
		return spots[adjacentPos] == nil
	}

	// okay, the hint that should be adjacent is placed -- let's see if it's placed correctly
	if spots[adjacentPos] == nil {
		// there's no house where the hint should be, so no, the hint isn't in the right place
		return false
	}

	// there is a house there -- does it contain the hint?
	hintFound := false
	for _, actualHintInAdjacentHouse := range spots[adjacentPos].Hints {
		if actualHintInAdjacentHouse == adjacentHint {
			hintFound = true
			break
		}
	}
	return hintFound
}

// if position satifies the spatial requirements of house
func spotSatisfies(house *House, position int, spots [5]*House) bool {
	// for all hints in the house, are the spatial reqs satisfied?
	for _, hint := range house.Hints {
		if hint.Position != nil && *hint.Position != position {
			// hint needs to be exactly at one position and isn't
			return false
		}
		// this hint needs to have a specific hint to its right -- does it, in this case?
		if hint.ToRight != nil && !requiredAdjacentHintIsAdjacent(hint.ToRight, position, true, spots) {
			return false
		}
		if hint.ToLeft != nil && !requiredAdjacentHintIsAdjacent(hint.ToLeft, position, false, spots) {
			return false
		}
		for _, nextTo := range hint.NextTo {
			// those in ToRight / ToLeft are already tested
			if nextTo != hint.ToRight && nextTo != hint.ToLeft {
				if !requiredAdjacentHintIsAdjacent(nextTo, position, true, spots) && !requiredAdjacentHintIsAdjacent(nextTo, position, false, spots) {
					return false
				}
			}
		}
	}
	return true
}

func unplaceLastHouse() *House {
	lastHousePos := len(placedHouses) - 1
	removeFrom := placedHouses[lastHousePos]

	// clear the spot
	for i, spot := range spots {
		if spot == removeFrom {
			spots[i] = nil
		}
	}
	// and remove from placed houses
	placedHouses = placedHouses[:lastHousePos]

	// return the unplaced house
	return removeFrom
}

func tryPlaceHouseAndProceed(house *House) (string, bool) {
	if placeHouseIfCan(house, nil) {
		// success! onward
		return whoOwnsTheFish(&House{}, 0)
	}
	// house's spatial reqs could not be fulfilled.
	// need to remove the last hint added, since it was no good :-(
	houseWithoutHint, hintsPos := removeLastHint(house)
	return whoOwnsTheFish(houseWithoutHint, hintsPos+1)
}

func placeHouseAtPosition(house *House, pos int) {
	spots[pos] = house

	// house may already be in placedHouses, if are placing house B and trying to move
	// house in order to accomodate this new house B.
	// if house is already in placedHouses, remove that copy.
	// either way, place house at the end of placedHouses, to indicate that it is the
	// most recently placed house
	foundAt := -1
	for i, existingPlaced := range placedHouses {
		if existingPlaced == house {
			foundAt = i
			break
		}
	}
	if foundAt != -1 {
		if foundAt == len(placedHouses)-1 {
			placedHouses = placedHouses[:foundAt]
		} else {
			placedHouses = append(placedHouses[:foundAt], placedHouses[foundAt+1:]...)
		}
	}
	placedHouses = append(placedHouses, house)
}

func placeHouseElsewhereIfCan(pos int, excludedPos map[int]bool) bool {
	if excludedPos == nil {
		excludedPos = make(map[int]bool)
	}
	// since want to place elsewhere, exclude pos (where it is currently placed)
	excludedPos[pos] = true
	house := spots[pos]
	spots[pos] = nil

	canMove := placeHouseIfCan(house, excludedPos)
	if !canMove {
		// can't go elsewhere, so put it back!
		spots[pos] = house
	}
	return canMove
}

// if you put the house in this spot, does the spot the adjacent house is in still work for that house?
func adjacentHouseStillSatified(house *House, position int, isToRight bool) bool {
	adjacentPos := position
	if isToRight {
		adjacentPos++
	} else {
		adjacentPos--
	}
	if adjacentPos < 0 || adjacentPos > 4 || spots[adjacentPos] == nil {
		// since there's nothing there, it is not unsatisfied
		return true
	}
	proposedSpots := createProposedSpots(house, position)
	adjacentHouse := proposedSpots[adjacentPos]
	// does the adjacent pos still satisfy the adjacent house?
	return spotSatisfies(adjacentHouse, adjacentPos, proposedSpots)
}

func createProposedSpots(house *House, pos int) [5]*House {
	var proposedSpots [5]*House
	for i := range spots {
		proposedSpots[i] = spots[i]
	}
	proposedSpots[pos] = house
	return proposedSpots
}
