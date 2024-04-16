package helloworld

const (
	french  = "French"
	spanish = "Spanish"

	helloEnglishPrefix = "Hello "
	helloSpanishPrefix = "Hola "
	helloFrenchPrefix  = "Bonjour "
)

func Hello(name, langauge string) string {
	if name == "" {
		name = "Abdoul"
	}

	return languagePrefixes(langauge) + name
}

func languagePrefixes(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = helloSpanishPrefix
	case french:
		prefix = helloFrenchPrefix
	default:
		prefix = helloEnglishPrefix
	}
	return
}
