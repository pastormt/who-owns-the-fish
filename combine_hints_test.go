package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_canCombine2Hints(t *testing.T) {
	cases := []struct {
		name     string
		h1, h2   *Hint
		expected bool
	}{
		{
			name: "yes",
			h1: &Hint{
				Color: "blue",
			},
			h2: &Hint{
				Beverage: "beer",
			},
			expected: true,
		},
		{
			name: "no",
			h1: &Hint{
				Color: "blue",
			},
			h2: &Hint{
				Color: "green",
			},
			expected: false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expected, canCombine2Hints(c.h1, c.h2))
		})
	}
}

func Test_removeLastHint(t *testing.T) {
	hint1 := Hint{
		Beverage: "coffee",
	}
	hint2 := Hint{
		Beverage: "tea",
	}
	hint3 := Hint{
		Color: "red",
	}

	house1 := House{
		Beverage: "coffee",
		Hints: []*Hint{
			&hint1,
		},
	}

	house2 := House{
		Beverage: "tea",
		Color:    "red",
		Hints: []*Hint{
			&hint2,
			&hint3,
		},
	}

	cases := []struct {
		name                                      string
		inputHouse, expectedHouse                 House
		inputPlacedHouses, expectedPlacedHouses   []*House
		inputSpots, expectedSpots                 [5]*House
		inputHintsInHouses, expectedHintsInHouses map[*Hint]int
		expectedRemovedHintPos                    int
	}{
		{
			name:       "houseIsEmpty",
			inputHouse: House{},
			expectedHouse: House{ // house 2 minus hint 3
				Beverage: "tea",
				Hints: []*Hint{
					&hint2,
				},
			},
			inputPlacedHouses: []*House{
				&house1,
				&house2,
			},
			expectedPlacedHouses: []*House{
				&house1,
				// house 2 unplaced since placed more recently
			},
			inputSpots: [5]*House{
				nil, &house1, &house2, nil, nil,
			},
			expectedSpots: [5]*House{
				nil, &house1, nil, nil, nil, // house 2 -> nil
			},
			inputHintsInHouses: map[*Hint]int{
				&hint1: 1, // the numbers here dont matter
				&hint2: 2,
				&hint3: 3, // will be removed, as house 2 was unplaced, and its most recently added hint (this) was removed
			},
			expectedHintsInHouses: map[*Hint]int{
				&hint1: 1, // the numbers here dont matter
				&hint2: 2,
			},
			expectedRemovedHintPos: 3,
		},
		{
			name: "houseIsNotEmpty",
			inputHouse: House{
				Beverage: "coffee",
				Hints: []*Hint{
					&hint1,
				},
			},
			expectedHouse:        House{Hints: []*Hint{}},
			inputPlacedHouses:    []*House{},
			expectedPlacedHouses: []*House{},
			inputHintsInHouses: map[*Hint]int{
				&hint1: 1, // the numbers here dont matter
			},
			expectedHintsInHouses:  map[*Hint]int{},
			expectedRemovedHintPos: 1,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			spots = c.inputSpots
			placedHouses = c.inputPlacedHouses
			hintsInHouses = c.inputHintsInHouses

			house, removedHintPos := removeLastHint(&c.inputHouse)

			assert.Equal(t, &c.expectedHouse, house)
			assert.Equal(t, c.expectedSpots, spots)
			assert.Equal(t, c.expectedRemovedHintPos, removedHintPos)
			assert.Equal(t, c.expectedPlacedHouses, placedHouses)
			assert.Equal(t, c.expectedHintsInHouses, hintsInHouses)
		})
	}
}

func Test_addHintToHouseIfCan(t *testing.T) {
	h1 := Hint{
		Beverage: "tea",
	}

	cases := []struct {
		name                                      string
		inputHint                                 *Hint
		inputHintPos                              int
		inputHouse, expectedHouse                 *House
		inputHintsInHouses, expectedHintsInHouses map[*Hint]int
		expectedAdded                             bool
	}{
		{
			name:               "can",
			inputHint:          &h1,
			inputHintPos:       1,
			inputHouse:         &House{},
			inputHintsInHouses: map[*Hint]int{},
			expectedHouse: &House{
				Beverage: "tea",
				Hints:    []*Hint{&h1},
			},
			expectedHintsInHouses: map[*Hint]int{
				&h1: 1,
			},
			expectedAdded: true,
		},
		{
			name:         "cant",
			inputHint:    &h1,
			inputHintPos: 1,
			inputHouse: &House{
				Beverage: "tea",
				Hints:    []*Hint{&h1},
			},
			inputHintsInHouses: map[*Hint]int{
				&h1: 1,
			},
			expectedHouse: &House{
				Beverage: "tea",
				Hints:    []*Hint{&h1},
			},
			expectedHintsInHouses: map[*Hint]int{
				&h1: 1,
			},
			expectedAdded: false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			hintsInHouses = c.inputHintsInHouses
			added := addHintToHouseIfCan(c.inputHint, c.inputHintPos, c.inputHouse)

			assert.Equal(t, c.expectedAdded, added)
			assert.Equal(t, c.expectedHouse, c.inputHouse)
			assert.Equal(t, c.expectedHintsInHouses, hintsInHouses)
		})
	}
}
