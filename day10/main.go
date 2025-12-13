package main

import (
	"bufio"
	"fmt"
	"math"
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

func parseIntList(s string) []int {
	spl := strings.Split(s, ",")

	l := make([]int, 0, len(spl))
	for _, intStr := range spl {
		n, err := strconv.Atoi(intStr)
		if err != nil {
			panic(err)
		}

		l = append(l, n)
	}

	return l
}

func parseButton(s string) button {
	l := len(s)
	if l < 3 {
		panic("Button too short")
	}
	if s[0] != '(' || s[l-1] != ')' {
		panic("Button has invalid format")
	}

	intListStr := s[1 : l-1]
	return parseIntList(intListStr)
}

func parseButtons(ss []string) []button {
	buttons := make([]button, 0, len(ss))
	for _, s := range ss {
		button := parseButton(s)
		buttons = append(buttons, button)
	}
	return buttons
}

func parseJoltage(s string) []int {
	l := len(s)
	if l < 3 {
		panic("Joltage list too short")
	}
	if s[0] != '{' || s[l-1] != '}' {
		panic("Joltage list has invalid format")
	}

	intListStr := s[1 : l-1]
	return parseIntList(intListStr)
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

		m := machine{diagram: parseDiagram(diagramStr), buttons: parseButtons(buttonStrs), joltageRequirements: parseJoltage(joltageRequirementsStr)}
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

func solveLights(m machine) int {
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
		sum += solveLights(m)
	}

	return sum
}

type equation struct {
	coeffs []int
	result int
}

func copyEquation(eq equation) equation {
	eq2 := equation{result: eq.result, coeffs: make([]int, len(eq.coeffs))}
	copy(eq2.coeffs, eq.coeffs)
	return eq2
}

func copyEquations(eqs []equation) []equation {
	eqs2 := make([]equation, len(eqs))
	for i := range eqs {
		eqs2[i] = copyEquation(eqs[i])
	}
	return eqs2
}

func printEquations(equations []equation) {
	for _, eq := range equations {
		for i, coeff := range eq.coeffs {
			_ = i
			// fmt.Printf("%d*x%d ", coeff, i)
			fmt.Printf("%d ", coeff)
		}
		fmt.Printf("| %d\n", eq.result)
	}
}

func toEquations(m machine) []equation {
	equations := make([]equation, len(m.joltageRequirements))

	for i, joltage := range m.joltageRequirements {
		equations[i].result = joltage
		equations[i].coeffs = make([]int, len(m.buttons))
	}

	for buttonI, button := range m.buttons {
		for _, eqI := range button {
			equations[eqI].coeffs[buttonI] = 1
		}
	}

	return equations
}

func gcd_internal(n1, n2 int) int {
	switch {
	case n1 == n2:
		return n1
	case n1 > n2:
		return gcd_internal(n1-n2, n2)
	default:
		return gcd_internal(n1, n2-n1)
	}
}

func Gcd(n1, n2 int) int {
	if n1 < 0 || n2 < 0 {
		panic("Cannot compute gcd with negative numbers")
	}
	return gcd_internal(n1, n2)
}

func Lcm(n1, n2 int) int {
	if n1 < 0 {
		n1 = -n1
	}
	if n2 < 0 {
		n2 = -n2
	}

	return n1 * n2 * Gcd(n1, n2)
}

func rowReduce(equations []equation) {
	// for i := range equations[0].coeffs {
	// }

	elimI := 0
	for i := range equations {
		// fmt.Println("----------------")
		// printEquations(equations)
		moveI := -1
		for moveI == -1 {
			if elimI >= len(equations[0].coeffs) {
				return
				// panic("ran out of variables to eliminate")
			}
			for j := i; j < len(equations); j++ {
				if equations[j].coeffs[elimI] != 0 {
					moveI = j
					break
				}
			}
			if moveI == -1 {
				elimI++
			}
		}

		if moveI != i {
			//swap equations
			tmp := equations[i]
			equations[i] = equations[moveI]
			equations[moveI] = tmp
		}

		mainCoeff := equations[i].coeffs[elimI]
		if mainCoeff == 0 {
			panic("pivot variable == 0")
		}

		if mainCoeff < 0 {
			for k := range equations[i].coeffs {
				equations[i].coeffs[k] *= -1
			}
			equations[i].result *= -1
			mainCoeff = equations[i].coeffs[elimI]
		}

		for j := i + 1; j < len(equations); j++ {
			mainCoeff2 := equations[j].coeffs[elimI]
			if mainCoeff2 != 0 {
				if mainCoeff2%mainCoeff != 0 {
					// make divisible
					for k := range equations[j].coeffs {
						equations[j].coeffs[k] *= mainCoeff
					}
					equations[j].result *= mainCoeff
					mainCoeff2 = equations[j].coeffs[elimI]
				}

				factor := mainCoeff2 / mainCoeff
				for k := range equations[j].coeffs {
					equations[j].coeffs[k] -= factor * equations[i].coeffs[k]
				}
				equations[j].result -= factor * equations[i].result
			}
		}

		elimI++
	}
}

func findDeterminingEquations(rowReducedSystem []equation) []int {
	l := len(rowReducedSystem[0].coeffs)
	determiningEquations := make([]int, l)

	varI := 0
	for eqI, row := range rowReducedSystem {
		stop := false
		for !stop {
			if varI == l {
				return determiningEquations
			}
			if row.coeffs[varI] == 0 {
				determiningEquations[varI] = -1
			} else {
				determiningEquations[varI] = eqI
				stop = true
			}
			varI++
		}
	}

	for varI < l {
		determiningEquations[varI] = -1
		varI++
	}

	return determiningEquations
}

var maxJoltage int = 0
var minG int = math.MaxInt
var bestSolution []int = nil

func h(equations []equation, determiningEquations []int, solution []int, varI int) int {
	if varI < 0 {
		sum := 0
		for _, val := range solution {
			sum += val
		}
		if sum < minG {
			minG = sum
			bestSolution = make([]int, len(solution))
			copy(bestSolution, solution)
		}
		return sum
	}

	isFreeVariable := determiningEquations[varI] == -1
	if !isFreeVariable {
		determiningEquation := equations[determiningEquations[varI]]
		val := determiningEquation.result
		for i := varI + 1; i < len(determiningEquations); i++ {
			val -= determiningEquation.coeffs[i] * solution[i]
		}

		denom := determiningEquation.coeffs[varI]
		if val%denom != 0 {
			// fmt.Printf("%d %% %d = %d\n", val, denom, val%denom)
			return -1
		}

		val /= denom
		if val < 0 {
			return -1
		}

		solution[varI] = val
		return h(equations, determiningEquations, solution, varI-1)
	} else {
		min := math.MaxInt
		// foundPossibleSolution := false

		val := 0
		for {
			if val > maxJoltage {
				if min == math.MaxInt {
					return -1
				} else {
					return min
				}
			}

			solution[varI] = val
			// fmt.Println("testing solution:", solution)
			res := h(equations, determiningEquations, solution, varI-1)
			// fmt.Println("res:", res)
			if res == -1 {
				//decide whether to stop or not
				// if foundPossibleSolution {
				// 	return min
				// }
			} else {
				// foundPossibleSolution = true
				if res < min {
					min = res
				}
			}

			val++
		}
	}
}

func findMinSolution(equations []equation, determiningEquations []int) int {
	numVars := len(determiningEquations)

	// Trim all-zero equations
	for i := len(equations) - 1; i >= 0; i-- {
		allZero := true
		for _, c := range equations[i].coeffs {
			if c != 0 {
				allZero = false
				break
			}
		}
		if !allZero {
			equations = equations[:i+1]
			break
		}
	}

	// fmt.Println("After trimming")
	// printEquations(equations)

	minG = math.MaxInt
	solution := make([]int, numVars)
	return h(equations, determiningEquations, solution, numVars-1)
}

func testSolution(equations []equation, solution []int) bool {
	for _, eq := range equations {
		res := 0
		for i := range solution {
			res += eq.coeffs[i] * solution[i]
		}
		if res != eq.result {
			fmt.Println(eq)
			return false
		}
	}

	return true
}

func solveEquations(equations []equation) int {
	orig := copyEquations(equations)
	// fmt.Println("before reduction")
	// printEquations(equations)
	rowReduce(equations)
	// fmt.Println("after reduction")
	// printEquations(equations)
	// fmt.Println("original")
	// printEquations(orig)
	determiningEquations := findDeterminingEquations(equations)
	// fmt.Println("determining equations", determiningEquations)

	min := findMinSolution(equations, determiningEquations)
	// fmt.Println("min:", min)
	// fmt.Println("global min:", minG, "solution:", bestSolution)

	solves1 := testSolution(equations, bestSolution)
	solves2 := testSolution(orig, bestSolution)
	if !solves1 {
		panic("does not solve reduced equations")
	}
	if !solves2 {
		panic("does not solve original equations")
	}
	return min
}

func getMaxJoltage(m machine) int {
	max := 0
	for _, j := range m.joltageRequirements {
		if j > max {
			max = j
		}
	}
	return max
}

func solveJoltage(m machine) int {
	maxJoltage = getMaxJoltage(m)
	equations := toEquations(m)
	return solveEquations(equations)
}

func part2(machines input) int {
	sum := 0

	for _, m := range machines {
		sum += solveJoltage(m)
	}

	return sum
}

func main() {
	input := getInput()
	fmt.Println("part 1:", part1(input))
	fmt.Println("part 2:", part2(input))
}
