package main

import (
	"fmt"
	"os"

	"github.com/luhring/ranchars/cli"
	"github.com/luhring/ranchars/ranchars"
)

func main() {
	if len(os.Args) == 1 {
		cli.ShowHelpAndExit()
	}

	arguments := os.Args[1:]

	if cli.ShouldHelpBeDisplayed(arguments) {
		cli.ShowHelpAndExit()
	}

	if cli.ShouldVersionBeDisplayed(arguments) {
		cli.ShowVersionAndExit()
	}

	length, err := cli.GetArgumentForLength(arguments)
	exitIfError(err)

	criteria := ranchars.GetCharacterCriteriaFromArguments(arguments)

	err = ranchars.VerifyThatCriteriaSatisfactionIsPossibleGivenLength(criteria, length)
	exitIfError(err)

	randomString, err := ranchars.GenerateRandomStringThatSatifiesAllCriteria(length, criteria)
	exitIfError(err)

	if false == cli.DoArgumentsSpecifyOmissionOfTrailingNewline(arguments) {
		randomString += "\n"
	}

	fmt.Print(randomString)
}

func exitIfError(err error) {
	if err != nil {
		fmt.Printf("error: %v\n", err.Error())
		os.Exit(1)
	}
}
