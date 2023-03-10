package addresses

import "strings"

// AddressType verify if has a valid type and return it
func AddressType(address string) string {
	validTypes := []string{"rua", "avenida", "estrada", "rodovia"}

	firstWord := strings.ToLower(address)
	firstWord = strings.Split(firstWord, " ")[0]

	hasValidType := false

	for _, typeAddress := range validTypes {
		if typeAddress == firstWord {
			hasValidType = true
		}
	}

	if hasValidType {
		return strings.Title(firstWord)
	}

	return "Tipo inválido"
}
