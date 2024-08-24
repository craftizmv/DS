package two_pointers

type TP struct {
}

func NewTP() *TP {
	//input := []int16{1, 2, 3, 4, 5}
	//input := []int16{1, 2, 3, 4}
	//var input []int16
	return &TP{}
}

// odd : 1, 2, 3, 4, 5
// even : 1, 2, 3, 4
// bound case : empty, 1
func (input *TP) ReverseArray(ip []int16) []int16 {
	sIdx := 0
	eIdx := len(ip) - 1
	for sIdx < eIdx {
		if sIdx == eIdx || eIdx < sIdx {
			return ip
		}
		swapInPlace(ip, sIdx, eIdx)
		sIdx++
		eIdx--
	}
	return ip
}

func swapInPlace(ip []int16, sIdx int, endIdx int) {
	temp := ip[endIdx]
	ip[endIdx] = ip[sIdx]
	ip[sIdx] = temp
}
