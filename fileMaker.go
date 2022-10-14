package main

import (
	"fmt"
	"os"
)

func CreateFile() {

  var Name string
  var Num int
  var Amount int 

    fmt.Println("filename?")
    fmt.Scanln(&Name)

    fmt.Println("Amount of files?")
    fmt.Scanln(&Amount)

    for i := 0; i <= Amount; i++{
	    Num++
    FileName := fmt.Sprint(Name, Num,  ".txt")

    os.Create(FileName)
    }
return
}

func main() {

    CreateFile()
    fmt.Println("file made!")

}
