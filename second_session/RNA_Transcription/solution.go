package rnatranscription

func ToRNA(dna string) string {
	retVal := ""
	for _, ch := range dna {
		switch ch {
		case 'G':
			retVal += "C"
		case 'C':
			retVal += "G"
		case 'T':
			retVal += "A"
		case 'A':
			retVal += "U"
		}
	}
	return retVal
}
