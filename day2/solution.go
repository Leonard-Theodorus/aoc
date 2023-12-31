package main

import (
    "fmt"
    "os"
    sc "strconv"
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
	if err == buf.Close(); err != nil{
	    return
	}
    }

    scanner := bufio.NewScanner(buf)
    for scanner.Scan(){
	result = append(result, scanner.Text())
    }
    err = scanner.Err()
    return result, err
}

func main(){
    data, err := os.ReadFile("input")
    check(err)
    var p = fmt.Println
    var total int = 0
    var cnt int = 0
    var l int = 0
    var w int = 0
    var temp string = ""
    for i :=0; i < len(data); i++{
	var s string = string(data[i])
	if s == "x"{
	    if cnt == 0 {
		length, err := sc.Atoi(temp)
		l = length
		check(err) 
	    } else if cnt == 1 { 
		width,err := sc.Atoi(temp) 
		check(err) 
		w = width 
	    } 
	    temp = ""
	    cnt += 1
	    continue
	} else if s == "\n" || s == ""{
	    cnt = 0
	    temp = ""
	    continue
	} else {
	    temp += s
	}

	if cnt >= 2{
	    h,err := sc.Atoi(temp)
	    check(err)
	    min1 := l * w
	    if w * h < min1 {
		min1 = w * h
	    }
	    if h * l < min1 {
		min1 = h * l
	    }
	    var surface int = (2 * l * w) + (2 * w * h) + (2 * h * l) + min1 
	    total += surface 
	    temp = ""
	    cnt = 0
	}
    }
    p(total)
    
}
