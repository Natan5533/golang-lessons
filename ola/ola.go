package ola

const prefixInEnglish = "Hello "
const prefixInPortuguese = "Eae, "
const prefixInFrench = "Bonjour "
const portuguese = "portuguese"
const french = "french"

// functions
func Ola(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return prefixResponse(language) + name
}

func prefixResponse(language string) (prefix string) {
	switch language {
	case french:
		prefix = prefixInFrench
	case portuguese:
		prefix = prefixInPortuguese
	default:
		prefix = prefixInEnglish
	}
	return
}
