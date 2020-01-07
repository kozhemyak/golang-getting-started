package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// 9 lesson: basic switch
	fmt.Println("\n----------\nLesson 9: switch")
	fmt.Print("GO runs go on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s\n", os)
	}

	// 10 lesson: more switch
	fmt.Println("\n----------\nLesson 10: switch")
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Horray! It's Today.")
	case today + 1:
		fmt.Println("It will be Tomorrow.")
	default:
		difference := int(time.Saturday) - int(today)
		fmt.Printf("It will be in %d days\n", difference)
	}

	// 11 lesson: switch without expression
	fmt.Println("\n----------\nLesson 11: switch")
	now := time.Now()
	switch {
	case now.Hour() < 12:
		fmt.Println("Good morning!")
	case now.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
