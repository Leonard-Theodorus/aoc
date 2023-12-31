package main

import (
    "fmt"
    "os"
)

func check(e error){
    if e != nil{
	panic(e) 
    }
}

func main(){
    data, err := os.ReadFile("input")
    check(err)
    var level int = 0
    var basement int = 0
    for i := 0; i < len(data); i++{
	//Part 2
	if level == -1 {
	    basement = i + 1
	    break
	}
	//MARK: PART 2

	var s string = string(data[i])
	if s == "(" {
	    level += 1
	} else {
	    level -= 1
	}
    }
    fmt.Printf("%d", basement - 1)
    //fmt.Printf("%d", level + 1)
}
