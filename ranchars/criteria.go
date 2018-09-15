package ranchars

import (
	"fmt"

	"github.com/luhring/ranchars/cli"
)

// Criterion represents the evaluation whether or not a byte satisfies a criterion function
type Criterion func(uint8) bool

// GetCharacterCriteriaFromArguments figures out the correct collection of criteria checks based on specified arguments
func GetCharacterCriteriaFromArguments(arguments []string) []Criterion {
	var criteria []Criterion

	if cli.IsFlagPresentAmongArguments(flagForIncludeAllTypesOfCharacters, arguments) {
		criteria = append(criteria, isASCIINumberDigit)
		criteria = append(criteria, isASCIINumberLowercaseLetter)
		criteria = append(criteria, isASCIINumberUppercaseLetter)
		criteria = append(criteria, isASCIINumberSpecialCharacter)
		criteria = append(criteria, isASCIINumberSpace)

		return criteria
	}

	if cli.IsFlagPresentAmongArguments(flagForIncludeDigits, arguments) {
		criteria = append(criteria, isASCIINumberDigit)
	}

	if cli.IsFlagPresentAmongArguments(flagForIncludeBothLowercaseAndUppercaseLetters, arguments) {
		criteria = append(criteria, isASCIINumberLowercaseLetter)
		criteria = append(criteria, isASCIINumberUppercaseLetter)
	} else {
		if cli.IsFlagPresentAmongArguments(flagForIncludeLowercaseLetters, arguments) {
			criteria = append(criteria, isASCIINumberLowercaseLetter)
		}

		if cli.IsFlagPresentAmongArguments(flagForIncludeUppercaseLetters, arguments) {
			criteria = append(criteria, isASCIINumberUppercaseLetter)
		}
	}

	if cli.IsFlagPresentAmongArguments(flagForIncludeSpecialCharacters, arguments) {
		criteria = append(criteria, isASCIINumberSpecialCharacter)
	}

	if cli.IsFlagPresentAmongArguments(flagForIncludeSpaces, arguments) {
		criteria = append(criteria, isASCIINumberSpace)
	}

	if len(criteria) == 0 {
		criteria = getDefaultCriteria()
	}

	return criteria
}

// VerifyThatCriteriaSatisfactionIsPossibleGivenLength ...
func VerifyThatCriteriaSatisfactionIsPossibleGivenLength(criteria []Criterion, length uint32) error {
	if length < uint32(len(criteria)) {
		return fmt.Errorf("it is impossible to generate a string with length '%v' that includes '%v' character types", length, len(criteria))
	}

	return nil
}

const flagForIncludeDigits = "-d"
const flagForIncludeLowercaseLetters = "-l"
const flagForIncludeUppercaseLetters = "-u"
const flagForIncludeBothLowercaseAndUppercaseLetters = "-L"
const flagForIncludeSpecialCharacters = "-c"
const flagForIncludeSpaces = "-s"
const flagForIncludeAllTypesOfCharacters = "-a"

func getDefaultCriteria() []Criterion {
	criteria := []Criterion{
		isASCIINumberDigit,
		isASCIINumberLowercaseLetter,
		isASCIINumberUppercaseLetter,
	}

	return criteria
}

func doesCharacterSatisfyAtLeastOneOfCriteria(character uint8, criteria []Criterion) bool {
	for _, isCriterionSatisifed := range criteria {
		if isCriterionSatisifed(character) {
			return true
		}
	}

	return false
}

func doesStringSatisfyAllCriteria(testString string, criteria []Criterion) bool {
	for _, criterion := range criteria {
		if false == doesStringContainCharactersThatMeetCriterion(testString, criterion) {
			return false
		}
	}

	return true
}

func doesStringContainCharactersThatMeetCriterion(testString string, criterion func(uint8) bool) bool {
	for _, character := range testString {
		if criterion(uint8(character)) {
			return true
		}
	}

	return false
}
