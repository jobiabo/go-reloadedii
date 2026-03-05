package main

import (
	"fmt"
	"os"
)

func NewSfolder() {
	newfile, _ := os.Create("My sample.txt")
	defer newfile.Close()
	fmt.Println("file created succesfully")

	byte, _ := newfile.WriteString("Go is good")
	fmt.Println(byte, "has been added successfully")
}

func Read() {
	openfile, _ := os.ReadFile("My sample.txt")

	fmt.Print(string(openfile))
}

func main() {
	Read()
}
