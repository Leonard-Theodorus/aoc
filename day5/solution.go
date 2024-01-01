package main

import(
    "fmt"
    "bufio"
    "os"
)
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

func isVowel(char string) bool{
    if char == "a" || char == "i" || char == "u" || char == "e" || char == "o"{
	return true
    } else {
	return false
    }
}

func doesContains(str string) bool {
    if str == "ab" || str == "cd" || str == "pq" || str == "xy"{
	return true
    } else {
	return false
    }
}

func checkPalindrome(index int, arr string) bool{
    if index - 1 < 0 || index + 1 >= len(arr) {return false}
    start := arr[index - 1]
    end := arr[index + 1]
    if start == end {
	return true
    } else {
	return false
    }
}


func main(){
    data, _ := Readlines("input")
    vowel := 0
    repeating := false
    contains := false
    prev := ""
    cnt := 0
    p := fmt.Println

    for i := 0; i < len(data); i++ {
	line := data[i]
	prev = string(line[0])
	if isVowel(prev) {
	    vowel += 1
	}
	for j:= 1; j < len(line); j++ {
	    content := string(line[j])
	    if doesContains(prev + content){
		contains = true
	    }
	    if content == prev{
		repeating = true
	    } else {
		prev = content 
	    }
	    if isVowel(content){
		vowel += 1
	    }
	}
	if vowel >= 3 && !contains && repeating{
	    cnt += 1
	}
	vowel = 0
	repeating = false
	contains = false
    }

    p(cnt)
    part2()

}

func part2(){
    data, _ := Readlines("input")
    cnt := 0
    p := fmt.Println
    prev := ""
    hasPair := false
    hasPalindrome := false

    for i := 0; i < len(data); i++ {
	line := string(data[i])
	prev = string(line[0])
	for j := 1; j < len(line); j++{
	    content := string(line[j])
	    pair := prev + content

	    for k := j + 1; k < len(line); k++ {
		if k + 1 < len(line){
		    tempPair := string(line[k]) + string(line[k + 1])
		    if tempPair == pair{
			hasPair = true
			break
		    }
		}
	    }
	    prev = content
	    palindrome := checkPalindrome(j, line)
	    if !hasPalindrome{
		hasPalindrome = palindrome
	    }
	}
	if hasPalindrome && hasPair{
	    cnt += 1
	}
	p(hasPalindrome, hasPair)
	hasPalindrome = false
	hasPair = false
    }
    p(cnt)
	
}
