package main

import (
	//"bufio"
	"fmt"
	"math/rand"
	//"os"
	//"strconv"
	//"strings"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	// fmt.Println("The secret number is ", secretNumber)

	fmt.Println("Please input your guess")
	//reader := bufio.NewReader(os.Stdin)
	for {
		var input int
		 
		cnt, err := fmt.Scanf("%d", &input)
		fmt.Println("the number of input : ", cnt)
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			continue
		}
		//input = strings.TrimSuffix(input, "\n")

		//guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer value")
			continue
		}
		fmt.Println("You guess is", input)
		if input > secretNumber {
			fmt.Println("Your guess is bigger than the secret number. Please try again")
		} else if input < secretNumber {
			fmt.Println("Your guess is smaller than the secret number. Please try again")
		} else {
			fmt.Println("Correct, you Legend!")
			break
		}
	}
}
