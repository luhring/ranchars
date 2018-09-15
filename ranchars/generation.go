package ranchars

import "crypto/rand"

// GenerateRandomStringThatSatifiesAllCriteria ...
func GenerateRandomStringThatSatifiesAllCriteria(length uint32, criteria []Criterion) (string, error) {
	for {
		testRandomString, err := generateRandomStringUsingCharactersBasedOnCriteria(length, criteria)

		if err != nil {
			return "", err
		}

		if doesStringSatisfyAllCriteria(testRandomString, criteria) {
			return testRandomString, nil
		}
	}
}

func generateRandomCharacterThatSatisfiesCriterion(criterion func(uint8) bool) (string, error) {
	randomCharacter, err := generateRandomCharacterThatSatisfiesAtLeastOneOfCriteria([]Criterion{criterion})

	return randomCharacter, err
}

func generateRandomCharacterThatSatisfiesAtLeastOneOfCriteria(criteria []Criterion) (string, error) {
	for {
		randomByte, err := generateRandomByte()

		if err != nil {
			return "", err
		}

		if doesCharacterSatisfyAtLeastOneOfCriteria(randomByte, criteria) {
			return string(randomByte), nil
		}
	}
}

func generateRandomLetterCharacter() (string, error) {
	character, err := generateRandomCharacterThatSatisfiesCriterion(isASCIINumberLetter)

	if err != nil {
		return "", err
	}

	return character, nil
}

func generateRandomUppercaseLetterCharacter() (string, error) {
	character, err := generateRandomCharacterThatSatisfiesCriterion(isASCIINumberUppercaseLetter)

	if err != nil {
		return "", err
	}

	return character, nil
}

func generateRandomLowercaseLetterCharacter() (string, error) {
	character, err := generateRandomCharacterThatSatisfiesCriterion(isASCIINumberLowercaseLetter)

	if err != nil {
		return "", err
	}

	return character, nil
}

func generateRandomSpecialCharacter() (string, error) {
	character, err := generateRandomCharacterThatSatisfiesCriterion(isASCIINumberSpecialCharacter)

	if err != nil {
		return "", err
	}

	return character, nil
}

func generateRandomDigitCharacter() (string, error) {
	character, err := generateRandomCharacterThatSatisfiesCriterion(isASCIINumberDigit)

	if err != nil {
		return "", err
	}

	return character, nil
}

func generateRandomByte() (uint8, error) {
	buffer := make([]byte, 1)
	_, err := rand.Read(buffer)

	if err != nil {
		return 0, err
	}

	return buffer[0], nil
}

func generateRandomStringUsingCharactersBasedOnCriteria(length uint32, criteria []Criterion) (string, error) {
	randomString := ""

	currentCharacter := uint32(1)

	for currentCharacter <= length {
		newCharacter, err := generateRandomCharacterThatSatisfiesAtLeastOneOfCriteria(criteria)

		if err != nil {
			return "", err
		}

		randomString += newCharacter

		currentCharacter++
	}

	return randomString, nil
}
