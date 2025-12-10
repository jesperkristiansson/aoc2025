package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type button []int

type machine struct {
	diagram             []bool
	buttons             []button
	joltageRequirements []int
}

type input []machine

func parseDiagram(s string) []bool {
	l := len(s)
	if l < 3 {
		panic("Diagram too short")
	}
	if s[0] != '[' || s[l-1] != ']' {
		panic("Diagram has invalid format")
	}

	diagram := make([]bool, 0, l-2)
	for _, r := range s[1 : l-1] {
		switch r {
		case '.':
			diagram = append(diagram, false)
		case '#':
			diagram = append(diagram, true)
		default:
			panic("Diagram contains invalid rune")
		}
	}

	return diagram
}

func parseButton(s string) button {
	l := len(s)
	if l < 3 {
		panic("Button too short")
	}
	if s[0] != '(' || s[l-1] != ')' {
		panic("Diagram has invalid format")
	}

	spl := strings.Split(s[1:l-1], ",")

	button := make(button, 0, len(spl))
	for _, intStr := range spl {
		n, err := strconv.Atoi(intStr)
		if err != nil {
			panic(err)
		}

		button = append(button, n)
	}

	return button
}

func parseButtons(ss []string) []button {
	buttons := make([]button, 0, len(ss))
	for _, s := range ss {
		button := parseButton(s)
		buttons = append(buttons, button)
	}
	return buttons
}

func getInput() input {
	var input input

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		lineSpl := strings.Split(line, " ")
		// There needs to be 1 diagram, at least 1 button, and 1 joltage requirement
		l := len(lineSpl)
		if l < 3 {
			panic("Too few strings in line")
		}
		diagramStr := lineSpl[0]
		buttonStrs := lineSpl[1 : l-1]
		joltageRequirementsStr := lineSpl[l-1]
		// Not used in part 1
		_ = joltageRequirementsStr

		m := machine{diagram: parseDiagram(diagramStr), buttons: parseButtons(buttonStrs)}
		input = append(input, m)
	}

	return input
}

func slicesEqual[T comparable](s1, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func pressButton(lights []bool, button button) {
	//fmt.Println("Before pressing button", button, ":", lights)
	for _, lightsI := range button {
		lights[lightsI] = !lights[lightsI]
	}
	//fmt.Println("After:", lights)
}

func tryButtonsHelper(target, current []bool, buttons []button, n int) bool {
	if n <= 0 {
		//fmt.Println("final state:", current)
		return slicesEqual(current, target)
	}

	if len(buttons) == 0 {
		panic("No buttons left to try")
	}

	//fmt.Println("N:", n, "Buttons to try:", buttons)

	buttonsToTry := buttons[:len(buttons)-n+1]
	for i := range buttonsToTry {
		newState := make([]bool, len(current))
		nCopied := copy(newState, current)
		if nCopied != len(current) {
			panic("Could not copy all elements from current light state")
		}

		pressButton(newState, buttonsToTry[i])
		remainingButtons := buttons[i+1:]
		if tryButtonsHelper(target, newState, remainingButtons, n-1) {
			return true
		}
	}

	return false
}

// Checks if the machine can be solved by pressing exactly n buttons
func tryButtons(m machine, n int) bool {
	//fmt.Println("trying buttons with", n)
	initialState := make([]bool, len(m.diagram))
	return tryButtonsHelper(m.diagram, initialState, m.buttons, n)
}

func solveMachine(m machine) int {
	for i := 1; i < len(m.buttons); i++ {
		if tryButtons(m, i) {
			return i
		}
	}

	panic("Machine could not be solved")
}

func part1(machines input) int {
	sum := 0

	for _, m := range machines {
		sum += solveMachine(m)
	}

	return sum
}

func part2(input input) int {
	sum := 0

	return sum
}

func main() {
	input := getInput()
	fmt.Println("part 1:", part1(input))
	fmt.Println("part 2:", part2(input))
}
