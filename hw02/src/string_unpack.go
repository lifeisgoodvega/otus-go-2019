package stringunpack

import (
	"errors"

	"strconv"
)

const (
	getSymbolOrNumber = iota
	getEscapedSymbol
)

func partialUnpack(currentSymbol rune, currentRepetitionNumber string, acc string) string {
	if currentSymbol != 0 {
		if currentRepetitionNumber != "" {
			n, _ := strconv.Atoi(currentRepetitionNumber)
			for i := 0; i < n; i++ {
				acc += string(currentSymbol)
			}
		} else {
			acc += string(currentSymbol)
		}
	}
	return acc
}

// Unpack - convert input string unpacking by rule - symbol[number] to symbol-symbol...number...symbol
// symbols can be escaped
func Unpack(inputString string) (string, error) {
	var currentSymbol rune = 0
	var resultString string = ""
	var currentRepetitionNumber string = ""
	mode := getSymbolOrNumber
	for _, symbol := range inputString {
		switch mode {
		case getSymbolOrNumber:
			if symbol >= '0' && symbol <= '9' {
				if currentSymbol == 0 {
					return "", errors.New("Numbers must be preceeded by symbols")
				}
				currentRepetitionNumber += string(symbol)
			} else {
				resultString = partialUnpack(currentSymbol, currentRepetitionNumber, resultString)
				currentRepetitionNumber = ""
				if symbol == '\\' {
					mode = getEscapedSymbol
				} else {
					currentSymbol = symbol
				}
			}
		case getEscapedSymbol:
			currentSymbol = symbol
			mode = getSymbolOrNumber
		}
	}
	resultString = partialUnpack(currentSymbol, currentRepetitionNumber, resultString)
	return resultString, nil
}
