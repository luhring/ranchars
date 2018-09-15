package ranchars

const asciiNumberForFirstUppercaseLetter uint8 = 65
const asciiNumberForLastUppercaseLetter uint8 = 90
const asciiNumberForFirstLowercaseLetter uint8 = 97
const asciiNumberForLastLowercaseLetter uint8 = 122
const asciiNumberForFirstDigit uint8 = 48
const asciiNumberForLastDigit uint8 = 57

func isASCIINumberUppercaseLetter(number uint8) bool {
	return number >= asciiNumberForFirstUppercaseLetter && number <= asciiNumberForLastUppercaseLetter
}

func isASCIINumberLowercaseLetter(number uint8) bool {
	return number >= asciiNumberForFirstLowercaseLetter && number <= asciiNumberForLastLowercaseLetter

}

func isASCIINumberLetter(number uint8) bool {
	return isASCIINumberUppercaseLetter(number) || isASCIINumberLowercaseLetter(number)
}

func isASCIINumberDigit(number uint8) bool {
	return number >= asciiNumberForFirstDigit && number <= asciiNumberForLastDigit
}

func isASCIINumberSpecialCharacter(number uint8) bool {
	return (number >= 33 && number <= 47) || (number >= 58 && number <= 64) || (number == 91) || (number >= 93 && number <= 96) || (number >= 123 && number <= 126)
}

func isASCIINumberSpace(number uint8) bool {
	return number == 32
}
