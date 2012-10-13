package chop

func chop(sought int, data []int) int {

	length := len(data)

	if length == 0 {
		// We didn't find it...
		return -1
	}

	cut := length / 2

	if data[cut] > sought {
		return chop(sought, data[0:cut])
	}
	if data[cut] < sought {
		return chop(sought, data[cut+1:length])
	}

	return cut
}
