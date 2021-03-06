package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Welcome to the game!")

	reader := bufio.NewReader(os.Stdin)

	rand.Seed(time.Now().Unix())
	var hidden = rand.Intn(20)
	//fmt.Println(hidden)
	var move = false
	for {
		fmt.Print("Your guess: ")
		var input string
		var err error
		input, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(input)
		input = strings.TrimSuffix(input, "\n")
		if input == "x" {
			fmt.Println("Good bye!")
			break
		}
		if input == "p" {
			fmt.Printf("The hidden number is %d\n", hidden)
			continue
		}
		if input == "m" {
			move = !move
			continue
		}

		var guess int
		guess, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
			fmt.Print("Try again!")
			continue
		}
		fmt.Println(guess)
		if guess < hidden {
			fmt.Println("Too small")
		} else if guess > hidden {
			fmt.Println("Too big")
		} else {
			fmt.Println("Direct hit!")
			break
		}
		hidden += rand.Intn(4) - 2
	}
}
