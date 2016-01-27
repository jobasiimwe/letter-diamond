package letterDiamond

import (
        "strings"
        "fmt"
        )

const alphabetLetters string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func IsValidInput(letter string) bool {
    return len(letter) ==  1 && strings.Index(alphabetLetters, letter) != -1
}

func LetterString(letter string) string {
    if IsValidInput(letter) == true {
        return strings.SplitAfter(alphabetLetters, letter)[0]
    } else {
        return ""
    }
}

func middleSpaces(index int) string {
    if index > 0 {
        return strings.Repeat(" ", (2 * index) - 1)
    } else {
        return ""
    }
}

func sideSpaces(index, length int) string {
    if (index + 1) < length {
        return strings.Repeat(" ", length - (index + 1))
    } else {
        return ""
    }
}

func StringForLetter(letter string, index, length int) string {
    sideSpaces := sideSpaces(index, length)
    middleSpaces := middleSpaces(index)
    if len(middleSpaces) > 0 {
        return sideSpaces + letter + middleSpaces + letter + sideSpaces
    } else {
        return sideSpaces + letter + sideSpaces
    }
}

func LetterStrings(lastLetter string) []string {
    letters := LetterString(lastLetter)
    var letterSlice = make([]string, len(letters))
    for index, letter := range letters {
        letterSlice[index] = StringForLetter(string(letter), index, len(letters))
    }
    return letterSlice
}

func PrintDiamond(lastLetter string) {
    fmt.Println("Diamond for: " + lastLetter)
    var letterSlice = LetterStrings(lastLetter)

    i := 0
    ascending := true
    for i >= 0 {
        fmt.Println(letterSlice[i])
        if i + 1 == len(letterSlice) {
            ascending = false
        }

        if ascending == true {
            i++
        }else{
            i--
        }
    }
}