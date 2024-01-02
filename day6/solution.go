package main

import (
    "fmt"
    "os"
    "bufio"
    s "strings"
    sc "strconv"
    "math"
)

func ReadLines(path string) ([]string, error){
    var lines []string
    file, err := os.Open(path)
    if err != nil {
	return lines, err
    }

    defer func (){
	fileCloseError := file.Close()
	if fileCloseError != nil {
	    return
	}
    }()

    scanner := bufio.NewScanner(file)
    for scanner.Scan(){
	lines = append(lines, scanner.Text())
    }
    scanError := scanner.Err()
    return lines,scanError
}

func processTurn(grid [1000][1000]bool, turnOn bool, from []string, to []string) ([1000][1000]bool){
    rowFrom, colFrom := 0, 0
    rowTo, colTo := 0, 0
    if len(from) != 0 {
	rowFrom, _ = sc.Atoi(from[0])
	colFrom, _ = sc.Atoi(from[1])
    }
    if len(to) != 0 {
	rowTo, _ = sc.Atoi(to[0])
	colTo, _ = sc.Atoi(to[1])
    }
    temp := colFrom
    if turnOn {
	for rowFrom <= rowTo{
	    for colFrom <= colTo{
		grid[rowFrom][colFrom] = true
		colFrom += 1
	    }
	    rowFrom += 1
	    colFrom = temp
	}

    } else {
	for rowFrom <= rowTo{
	    for colFrom <= colTo{
		grid[rowFrom][colFrom] = false
		colFrom += 1
	    }
	    rowFrom += 1
	    colFrom = temp
	}

    }
    return grid
}

func processToggle(grid [1000][1000]bool, from []string, to []string) ([1000][1000]bool){
    rowFrom, colFrom := 0, 0
    rowTo, colTo := 0, 0
    if len(from) != 0 {
	rowFrom, _= sc.Atoi(from[0])
	colFrom, _= sc.Atoi(from[1])
    }
    if len(to) != 0 {
	rowTo, _ = sc.Atoi(to[0])
	colTo, _ = sc.Atoi(to[1])
    }
    temp := colFrom
    for rowFrom <= rowTo{
	for colFrom <= colTo{
	    grid[rowFrom][colFrom] = !grid[rowFrom][colFrom]
	    colFrom += 1
	}
	rowFrom += 1
	colFrom = temp
    }
    return grid
}

func processToggleBrightness(grid [1000][1000]int, from []string, to []string) ([1000][1000]int){
    rowFrom, colFrom := 0, 0
    rowTo, colTo := 0, 0
    if len(from) != 0 {
	rowFrom, _= sc.Atoi(from[0])
	colFrom, _= sc.Atoi(from[1])
    }
    if len(to) != 0 {
	rowTo, _ = sc.Atoi(to[0])
	colTo, _ = sc.Atoi(to[1])
    }
    temp := colFrom
    for rowFrom <= rowTo{
	for colFrom <= colTo{
	    grid[rowFrom][colFrom] += 2 
	    colFrom += 1
	}
	rowFrom += 1
	colFrom = temp
    }
    return grid
}

func processTurnBrightness(grid [1000][1000]int, turnOn bool, from []string, to []string) ([1000][1000]int){
    rowFrom, colFrom := 0, 0
    rowTo, colTo := 0, 0
    if len(from) != 0 {
	rowFrom, _ = sc.Atoi(from[0])
	colFrom, _ = sc.Atoi(from[1])
    }
    if len(to) != 0 {
	rowTo, _ = sc.Atoi(to[0])
	colTo, _ = sc.Atoi(to[1])
    }
    temp := colFrom
    if turnOn {
	for rowFrom <= rowTo{
	    for colFrom <= colTo{
		grid[rowFrom][colFrom] += 1
		colFrom += 1
	    }
	    rowFrom += 1
	    colFrom = temp
	}

    } else {
	for rowFrom <= rowTo{
	    for colFrom <= colTo{
		grid[rowFrom][colFrom] = int(math.Max(float64(grid[rowFrom][colFrom] - 1), float64(0)))
		colFrom += 1
	    }
	    rowFrom += 1
	    colFrom = temp
	}

    }
    return grid
    
}

func main(){
    data, err := ReadLines("input")
    if err != nil {
	panic(err)
    }
    p := fmt.Println
    cnt := 0
    var lights [1000][1000]bool
    for i := 0; i < 1000; i++ {
	for j := 0; j < 1000; j++{
	    lights[i][j] = false
	}
    }
    for i := 0; i < len(data); i++{
	line := string(data[i])
	substr := s.Split(line, " ")
	command := substr[0]
	if command == "toggle"{
	    from := substr[1]
	    fromC := s.Split(from, ",")
	    to := substr[3]
	    toC := s.Split(to, ",")
	    lights = processToggle(lights, fromC, toC)
	} else {
	    inst := substr[1]
	    turnOn := false
	    from := substr[2]
	    to := substr[4]
	    fromC := s.Split(from, ",")
	    toC := s.Split(to, ",")
	    if inst == "on"{
		turnOn = true
	    }
	    lights = processTurn(lights, turnOn, fromC, toC)
	}
    }

    for i := 0; i < 1000; i++ {
	for j := 0; j < 1000; j++{
	    if lights[i][j] {cnt += 1}
	}
    }
    p(cnt)

    part2()
}

func part2(){
    data, err := ReadLines("input")
    if err != nil {
	panic(err)
    }
    p := fmt.Println
    cnt := 0
    var lights [1000][1000]int
    for i := 0; i < 1000; i++ {
	for j := 0; j < 1000; j++{
	    lights[i][j] = 0
	}
    }
    for i := 0; i < len(data); i++{
	line := string(data[i])
	substr := s.Split(line, " ")
	command := substr[0]
	if command == "toggle"{
	    from := substr[1]
	    fromC := s.Split(from, ",")
	    to := substr[3]
	    toC := s.Split(to, ",")
	    lights = processToggleBrightness(lights, fromC, toC)
	} else {
	    inst := substr[1]
	    turnOn := false
	    from := substr[2]
	    to := substr[4]
	    fromC := s.Split(from, ",")
	    toC := s.Split(to, ",")
	    if inst == "on"{
		turnOn = true
	    }
	    lights = processTurnBrightness(lights, turnOn, fromC, toC)
	}
    }

    for i:= 0; i < 1000; i++{
	for j := 0; j < 1000; j++ {
	    cnt += lights[i][j]
	}
    }
    p(cnt)

}
