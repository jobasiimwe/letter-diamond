package main

import (
	"github.com/jobasiimwe/letter-diamond"
    "fmt"
    "os"
)

func main() {
    fmt.Println("Enter an Upper Case alphabetic letter to draw a diamond!!!")
    var letter []byte = make([]byte, 1)
    for {
        os.Stdin.Read(letter)
        if letterDiamond.IsValidInput(string(letter)) == true {
            letterDiamond.PrintDiamond(string(letter))
        } else {
            if string(letter) == "\n" {
                continue
            }

            if string(letter) == "7" {
                break
            } else {
                fmt.Println(string(letter) + " is invalid input!!!")
            }
        }

    }
}
