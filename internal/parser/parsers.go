package parser

func Parseable(value string) bool {
	if isParsableJSON(value) {
		return true
	} else if isParsableXML(value) {
		return true
	}
	return false
}
