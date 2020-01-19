package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	truth := true
	negate(&truth)
	fmt.Println(truth)
	lies := false
	negate(&lies)
	fmt.Println(lies)
}

func negate(myBoolean *bool) {
	*myBoolean = !*myBoolean
}

func createPointer(myFloatValue float64) *float64 {
	return &myFloatValue
}

func printPointer(myFloatValue *float64) {
	print(*myFloatValue)
}

func double(number *int) {
	*number *= 2
}

func computeArea(height float64, width float64) (float64, error) {
	if width <= 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height <= 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}
	return (width * height), nil
}

func generateRandomNumber() int {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	return target
}

func guessingGame(attempts int) {
	target := generateRandomNumber()

	fmt.Println("I've chosen a random number between 1 and 100")
	fmt.Println("Can you guess it?")

	success := false

	for guesses := 0; guesses < attempts; guesses++ {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Make a Guess: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Oops. Your guess was Low")
		} else if guess > target {
			fmt.Println("Oops. Your guess was High")
		} else {
			success = true
			fmt.Println("Good Job! You Guessed it!")
			break
		}
	}

	if !success {
		fmt.Println("You Failed to Guess the number. It was ", target)
	}
}

func getFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}
