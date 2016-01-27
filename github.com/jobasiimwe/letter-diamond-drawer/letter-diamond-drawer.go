package main

import (
	"github.com/jobasiimwe/letter-diamond"
    "fmt"
    "os"
)

func main() {
    fmt.Println("Enter an Upper Case alphabetic letter to draw a diamond or 7 to exit!")
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
                fmt.Println("Goodbye, Diamond Drawer is going to sleep!!!")
                break
            } else {
                fmt.Println(string(letter) + " is invalid input!!!")
            }
        }

    }
}
