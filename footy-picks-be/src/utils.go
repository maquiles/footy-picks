package main

func contains(slice []int, value int) bool {
	for _, val := range slice {
		if val == value {
			return true
		}
	}

	return false
}
