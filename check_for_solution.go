package main

// assumes <= 1 house is missing a pet
func findHouseMissingAPet() *House {
	for _, house := range spots {
		if house == nil {
			continue
		}
		if house.Pet == "" {
			return house
		}
	}
	return nil
}

func spotsAreFull(spots [5]*House) bool {
	for _, spot := range spots {
		if spot == nil {
			return false
		}
	}
	return true
}

// not necessarily that all houses are placed
// just that those placed do not include a partially complete house (ie missing a pet!)
func spotsHaveOnlyCompleteHouses() bool {
	for _, house := range spots {
		if house != nil && !house.isComplete() {
			return false
		}
	}
	return true
}
