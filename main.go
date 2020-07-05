package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var maze []string

func loadMaze(file string) error {
	f, err := os.Open()
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		maze = append(maze, line)
	}
	return nil
}

func printScreen() {
	for _, line := range maze {
		log.Println(line)
	}
}

func main() {
	//initialise game
	err := loadMaze("maze01.txt")
	if err != nil {
		fmt.Println("failed to load maze: ", err)
		return
	}
	//load resources

	for {
		//update screen
		printScreen()
		//process input
		//process movement
		//process collisions
		//check game over

		fmt.Println("Hello, Pac Go!")
		break // breaks the infinite loop
	}
}
