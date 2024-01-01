package main


import (
    "fmt"
    "bufio"
    "os"
)
func ReadLines(path string) ([]string, error){
    var result []string
    buf, err := os.Open(path)
    if err != nil{
	return result, err
    }

    defer func (){
	if err = buf.Close(); err != nil {
	    return
	}
    }()

    scanner := bufio.NewScanner(buf)
    for scanner.Scan(){
	result = append(result, scanner.Text())
    }
    err = scanner.Err()
    return result,err
}
type coordinate struct {
    x int
    y int
}

func main(){
    lines, _ := ReadLines("input")
    //p := fmt.Println
    cnt := 1
    vert := 0
    horz := 0
    visited := make(map[coordinate]bool)

    for i:= 0; i < len(lines); i++ {
	line := lines[i]
	for j:= 0; j < len(line); j++ {
	    char := string(line[j])
	    switch char {
	    case "^":
		vert += 1
	    case "v":
		vert -= 1
	    case ">":
		horz += 1
	    case "<":
		horz -= 1
	    }
	    newPair := coordinate{horz, vert}
	    visited[newPair] = true
	}
    }
    cnt += len(visited) - 1
    part2()

}
func part2(){
    lines, _ := ReadLines("input")
    p := fmt.Println
    cnt := 0
    vertSanta := 0
    horzSanta := 0
    vertRobo := 0
    horzRobo := 0
    santa := false
    visited := make(map[coordinate]bool)
    visitedRobo := make(map[coordinate]bool)

    for i:= 0; i < len(lines); i++ {
	line := lines[i]
	for j:= 0; j < len(line); j++ {
	    if j % 2 == 0 {
		print("GENAP")
		santa = true
	    } else {
		print("GANJIL")
		santa = false
	    }
	    char := string(line[j])
	    if santa == true{
		switch char {
		case "^":
		    vertSanta += 1
		case "v":
		    vertSanta -= 1
		case ">":
		    horzSanta += 1
		case "<":
		    horzSanta -= 1
		}
		if i == 0 && j == 0 {
		    horzRobo, vertRobo = horzSanta, vertSanta
		    visitedRobo[coordinate{horzSanta, vertSanta}] = true
		}
		santaPair := coordinate{horzSanta, vertSanta}
		visited[santaPair] = true
		fmt.Printf("santa : %d %d \n", horzSanta, vertSanta)
	    } else if santa == false{
		switch char {
		case "^":
		    vertRobo += 1
		case "v":
		    vertRobo -= 1
		case ">":
		    horzRobo += 1
		case "<":
		    horzRobo -= 1
		}
		roboPair := coordinate{horzRobo, vertRobo}
		visitedRobo[roboPair] = true

		fmt.Printf("robo : %d %d \n", horzRobo, vertRobo)
	    }
	}
    }
    p(len(visited))
    p(len(visitedRobo))
    p(cnt)


}
