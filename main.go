package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	t := Tetromino(scanner.Text())

	scanner.Scan()
	line2 := strings.Split(scanner.Text(), " ")
	commands := []Command{}
	for _, l := range line2 {
		commands = append(commands, Command(l))
	}
	fmt.Println(t)
	fmt.Println(commands)

	brick := &Brick{t, 5, 0, POINTUP}
	for _, command := range commands {
		RunCommand(brick, command)
	}
	fmt.Println(brick)
}
