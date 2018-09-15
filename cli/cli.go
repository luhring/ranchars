package cli

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const version = "ranchars version 0.1.0"

// ShowVersionAndExit outputs the program's current version and then exits the program
func ShowVersionAndExit() {
	fmt.Println(version)
	os.Exit(0)
}

// ShouldVersionBeDisplayed ...
func ShouldVersionBeDisplayed(arguments []string) bool {
	return len(arguments) == 1 && (arguments[0] == "-v" || arguments[0] == "--version")
}

// ShowHelpAndExit outputs command-line help to stdout and then exits the program
func ShowHelpAndExit() {
	fmt.Println(helpText)
	os.Exit(0)
}

// ShouldHelpBeDisplayed ...
func ShouldHelpBeDisplayed(arguments []string) bool {
	return len(arguments) == 1 && (arguments[0] == "-h" || arguments[0] == "--help")
}

// GetArgumentForLength from slice of arguments
func GetArgumentForLength(arguments []string) (uint32, error) {
	for _, argument := range arguments {
		if false == isArgumentFlag(argument) {
			numberOfCharactersAsInt64, err := strconv.ParseUint(argument, 10, 32)

			if err != nil {
				return 0, err
			}

			numberOfCharacters := uint32(numberOfCharactersAsInt64)

			return numberOfCharacters, nil
		}
	}

	return 0, errors.New("unable to find argument for number of characters")
}

// DoArgumentsSpecifyOmissionOfTrailingNewline ...
func DoArgumentsSpecifyOmissionOfTrailingNewline(arguments []string) bool {
	return IsFlagPresentAmongArguments(flagForOmitTrailingNewLine, arguments)
}

// IsFlagPresentAmongArguments checks for an individual flag name among a slice of arguments
func IsFlagPresentAmongArguments(flagName string, arguments []string) bool {
	if false == isArgumentFlag(flagName) {
		return false
	}

	for _, argument := range arguments {
		if isArgumentFlag(argument) && doesFlagBeginWithSingleDash(argument) {
			if strings.Contains(argument, flagName[1:2]) {
				return true
			}
		}
	}

	return false
}

const flagForOmitTrailingNewLine = "-n"

func isArgumentFlag(argument string) bool {
	if len(argument) < 2 {
		return false
	}

	return argument[0:1] == "-"
}

func doesFlagBeginWithSingleDash(argument string) bool {
	if len(argument) < 1 {
		return false
	}

	return argument[0:1] == "-" && argument[0:2] != "--"
}

const helpText = `
Usage:  ranchars [options] length

Generates a string of random characters that meets the specified criteria, of a specified length. With no options specified, the output string will contain lowercase letters, uppercase letters, and numeric digits.

'length' must be a positive integer. 'length' must also be greater than or equal to the specified number of character types. For example, it is impossible to generate a 3-character string that includes at least one numeric digit, special character, lowercase letter, and uppercase letter.

Options:
  -d	Include numeric digits (0-9)
  -l	Include lowercase letters (a-z)
  -u	Include uppercase letters (A-Z)
  -L	Include both lowercase and uppercase letters (equivalent to specifying both -l and -u)
  -c	Include special characters
  -s	Include the space character
  -a	Include all kinds of characters (equivalent to specifying -d, -l, -u, -c, and -s)
  -n	Omit trailing newline character from output

Examples:
  ranchars -d 6
  --> generates a random 6-digit number, e.g. 214645

  ranchars -L -c 40
  --> generates a random string, 40 characters long, that includes lowercase and uppercase letters and special characters, e.g. qn$Doq*s)S(Opxf?AJE?W)!=#MC{]voR#AjhT$,K
`
