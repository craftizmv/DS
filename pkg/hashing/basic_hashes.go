package hashing

func basicHash(s string) uint32 {
	var hash uint32 = 0
	const prime uint32 = 31

	for i := 0; i < len(s); i++ {
		hash = hash*prime + uint32(s[i])
	}

	return hash
}
