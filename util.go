package fabric

func contains(s []int, i int) bool {
	for _, v := range s {
		if i == v {
			return true
		}
	}
	return false
}
