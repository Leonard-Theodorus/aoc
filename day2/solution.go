package main

import (
    "fmt"
    "os"
    "bufio"
)
func check(e error){
    if e != nil{
	panic(e)
    }
}

func minmin(a, b int) int {
    if a < b {
	return a
    } else {
	return b
    }
}

func Readlines(path string) ([]string, error){
    var result []string
    buf, err := os.Open(path)
    if err != nil {
	return result, err
    }
    defer func() {
	if err = buf.Close(); err != nil{
	    return
	}
    }()

    scanner := bufio.NewScanner(buf)
    for scanner.Scan(){
	result = append(result, scanner.Text())
    }
    err = scanner.Err()
    return result, err
}

func main(){
    lines, err := Readlines("input")
    if err != nil {
	panic(err)
    }
    var total int = 0

    for i:= 0; i < len(lines); i++{
	l := 0
	w := 0
	h := 0

	_, err := fmt.Sscanf(lines[i], "%dx%dx%d", &l, &w, &h)
	if err != nil {
	    message := fmt.Sprintf("Failed Scan: '%v' %v", lines[i], err)
	    panic(message)

	}
	side1 := l * w
	side2 := w * h
	side3 := h * l

	smallest := side1
	
	if side2 < smallest {
	    smallest = side2
	}
	if side3 < smallest{
	    smallest = side3
	}

	paper := 2 * side1 + 2 * side2 + 2 * side3 + smallest
	total += paper

    }
    fmt.Println(total)
    part2()
}

func part2(){
    lines, err := Readlines("input")
    if err != nil {
	panic(err)
    }
    var total int = 0

    for i:= 0; i < len(lines); i++{
	l := 0
	w := 0
	h := 0

	_, err := fmt.Sscanf(lines[i], "%dx%dx%d", &l, &w, &h)
	if err != nil {
	    message := fmt.Sprintf("Failed Scan: '%v' %v", lines[i], err)
	    panic(message)

	}
	side1 := 2 * l + 2 * w
	side2 := 2 * w + 2 * h
	side3 := 2 * h + 2 * l

	smallest := side1
	
	if side2 < smallest {
	    smallest = side2
	}
	if side3 < smallest{
	    smallest = side3
	}

	ribbon := smallest + (l * w * h)
	total += ribbon 

    }
    fmt.Println(total)
    
}
