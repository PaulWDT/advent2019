package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Hello Advent of code 2019")

	myFile, err := os.Open("input")
	if err != nil {
		fmt.Println("File open error !")
		return
	}
	defer myFile.Close()

	myScanner := bufio.NewScanner(myFile)
	myScanner.Split(bufio.ScanLines)

	var i, v int
	for myScanner.Scan() {
		i++
		zeile := myScanner.Text()
		fmt.Print(zeile)
		fmt.Printf(" %T\n", zeile)
		zahl, _ := strconv.Atoi(zeile)
		fmt.Print(zahl)
		fmt.Printf(" %T\n", zahl)
		v += (zahl / 3) - 2
	}
	fmt.Println("Count lines = ", i)
	fmt.Println("Accumulated fuel requirements = ", v)

}
