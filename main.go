package main

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const minute int = 60

type block struct {
	minutes int
	seconds int
	number  int
}

func (m model) Init() tea.Cmd {
	return tea.SetWindowTitle("Grocery List")
}

func timer(workMin int) {
	var minuteTimer int
	var secondTimer int
	for minuteTimer = workMin - 1; minuteTimer > -1; minuteTimer-- {
		for secondTimer = minute - 1; secondTimer > -1; secondTimer-- {
			fmt.Println(minuteTimer, "m", secondTimer, "s")
			time.Sleep(time.Second)
		}
	}
}

func runBlocks(blocks int, workMin int, pauseMin int) {
	var blockTimer int
	for blockTimer = blocks; blockTimer > 0; blockTimer-- {
		timer(workMin)
		timer(pauseMin)
	}
}

func main() {
	var workMin int
	var pauseMin int
	var blocks int

	fmt.Println("How many minutes you will work?")
	fmt.Scanln(&workMin)
	fmt.Println("How many minutes you will pause?")
	fmt.Scanln(&pauseMin)
	fmt.Println("How many block will you do?")
	fmt.Scanln(&blocks)

	runBlocks(blocks, workMin, pauseMin)
}
