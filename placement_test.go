package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_unplaceLastHouse(t *testing.T) {
	var houseA, houseB House

	// houseA is in 0th pos and houseB in 3rd pos
	spots = [5]*House{
		&houseA, nil, nil, &houseB, nil,
	}
	// houseA was placed and then houseB was placed
	placedHouses = []*House{
		&houseA, &houseB,
	}

	// remove B
	unplacedHouse := unplaceLastHouse()
	require.True(t, &houseB == unplacedHouse)

	expectedSpots := [5]*House{
		&houseA, nil, nil, nil, nil, // houseB removed
	}
	expectedPlacedHouses := []*House{
		&houseA, // houseB removed
	}
	require.Equal(t, expectedSpots, spots)
	require.Equal(t, expectedPlacedHouses, placedHouses)

	// now remove A
	unplacedHouse = unplaceLastHouse()
	require.True(t, &houseA == unplacedHouse)

	expectedSpots = [5]*House{
		nil, nil, nil, nil, nil, // houseA removed
	}
	expectedPlacedHouses = []*House{}

	require.Equal(t, expectedSpots, spots)
	require.Equal(t, expectedPlacedHouses, placedHouses)
}

func Test_placeHouseAtPosition(t *testing.T) {
	var houseA House

	spots = [5]*House{}
	placedHouses = []*House{}

	placeHouseAtPosition(&houseA, 1)

	expectedSpots := [5]*House{
		nil, &houseA, nil, nil, nil,
	}
	assert.Equal(t, expectedSpots, spots)

	expectedPlacedHouses := []*House{
		&houseA,
	}
	assert.Equal(t, expectedPlacedHouses, placedHouses)
}

func Test_placeHouseIfCan_simplest(t *testing.T) {
	spots = [5]*House{}
	placedHouses = []*House{}

	hint1 := Hint{
		Color: "red",
	}
	houseA := House{
		Color: "red",
		Hints: []*Hint{
			&hint1,
		},
	}

	hintsInHouses = map[*Hint]int{
		&hint1: 0,
	}

	wasPlaced := placeHouseIfCan(&houseA, nil)
	assert.True(t, wasPlaced)

	// should be placed in 0th position
	expectedSpots := [5]*House{
		&houseA, nil, nil, nil, nil,
	}
	assert.Equal(t, expectedSpots, spots)

	expectedPlacedHouses := []*House{
		&houseA,
	}
	assert.Equal(t, expectedPlacedHouses, placedHouses)
}

func Test_placeHouseIfCan_simple_toLeft(t *testing.T) {
	hint1 := Hint{
		Color: "red",
	}
	hint2 := Hint{
		Beverage: "tea",
	}

	hint1.ToLeft = &hint2
	hint1.NextTo = []*Hint{&hint2}

	hint2.ToRight = &hint1
	hint2.NextTo = []*Hint{&hint1}

	houseWithHint1 := House{
		Color: "red",
		Hints: []*Hint{
			&hint1,
		},
	}

	houseWithHint2 := House{
		Beverage: "tea",
		Hints: []*Hint{
			&hint2,
		},
	}

	// house with hint 1 is already placed!
	// house with hint 2 should go to its left
	spots = [5]*House{
		nil, nil, &houseWithHint1, nil, nil,
	}
	placedHouses = []*House{
		&houseWithHint1,
	}

	// numbers don't matter here
	hintsInHouses = map[*Hint]int{
		&hint1: 0,
		&hint2: 1,
	}

	wasPlaced := placeHouseIfCan(&houseWithHint2, nil)
	assert.True(t, wasPlaced)

	// should be placed left of hint 1
	expectedSpots := [5]*House{
		nil, &houseWithHint2, &houseWithHint1, nil, nil,
	}
	assert.Equal(t, expectedSpots, spots)

	expectedPlacedHouses := []*House{
		&houseWithHint1, &houseWithHint2,
	}
	assert.Equal(t, expectedPlacedHouses, placedHouses)
}

func Test_placeHouseIfCan_adjacentSpotIsStillSatisfied(t *testing.T) {
	hint1 := Hint{
		Color: "red",
	}
	hint2 := Hint{
		Beverage: "tea",
	}

	hint1.ToLeft = &hint2
	hint1.NextTo = []*Hint{&hint2}

	hint2.ToRight = &hint1
	hint2.NextTo = []*Hint{&hint1}

	houseWithHint1 := House{
		Color: "red",
		Hints: []*Hint{
			&hint1,
		},
	}

	// any spot fulfills its spatial reqs
	var houseFakeEmpty House

	// scenario below:
	// houseWithHint1 is already placed
	// houseFakeEmpty should not be able to go to its left,
	// since hint1 has a ToLeft and houseFakeEmpty does not fulfill it
	spots = [5]*House{
		nil, &houseWithHint1, nil, nil, nil,
	}
	placedHouses = []*House{
		&houseWithHint1,
	}

	hintsInHouses = map[*Hint]int{
		&hint1: 0, // 0 doesn't matter... (wherever hint1 was in the original list of hints is irrelevant right now)
	}

	wasPlaced := placeHouseIfCan(&houseFakeEmpty, nil)
	assert.True(t, wasPlaced)

	// houseFakeEmpty should be placed at spot 2 (middle) since it's the least resistance
	// position 0 doesn't (immediately) work, since left of houseWithHint1 should remain open
	// position 1 doesn't (immediately) work, since there's something there
	// for simplicitly, we only move things if we can't place what we're trying to place without moving things
	expectedSpots := [5]*House{
		nil, &houseWithHint1, &houseFakeEmpty, nil, nil,
	}
	assert.Equal(t, expectedSpots, spots)

	expectedPlacedHouses := []*House{
		&houseWithHint1, &houseFakeEmpty,
	}
	assert.Equal(t, expectedPlacedHouses, placedHouses)
}

func Test_createProposedSpots(t *testing.T) {
	house := House{
		Color: "red",
	}
	spots = [5]*House{
		nil, &house, nil, nil, nil,
	}
	newHouse := House{
		Color: "blue",
	}
	proposedSpots := createProposedSpots(&newHouse, 2)
	expectedProposedSpots := [5]*House{
		nil, &house, &newHouse, nil, nil,
	}
	assert.Equal(t, expectedProposedSpots, proposedSpots)

	proposedSpots[1] = nil
	expectedProposedSpots = [5]*House{
		nil, nil, &newHouse, nil, nil,
	}
	assert.Equal(t, expectedProposedSpots, proposedSpots)

	expectedSpots := [5]*House{
		nil, &house, nil, nil, nil,
	}
	assert.Equal(t, expectedSpots, spots)
}
