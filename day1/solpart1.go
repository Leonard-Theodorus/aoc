package main

import (
	"fmt"
	"log"
	"os"
    "bufio"
    "regexp"
    "strconv"
    "strings"
    "sort"
)
func isDigit(input string) (bool, int){
    var check = regexp.MustCompile(`^[0-9]+$`);
    res, err := strconv.Atoi(input);
    if (err != nil){
        // number invalid
    }
    return check.MatchString(input), res;
}

func searchFirst(input string) ([]int, []int) {
    tries := map[string]int {"one" : 1, 
    "two" : 2, "three" : 3, "four" : 4, 
    "five" : 5, "six" : 6, "seven" : 7, 
    "eight" : 8, "nine" : 9};

    idxs := []int {-1, -1, -1, -1, -1, -1, -1, -1, -1};

    for k, v := range tries {
        trying := k;
        idx := strings.Index(input, trying);
        if ( idx != - 1) {
            // The word exists
            idxs[v - 1] = idx;
        }
    }
    mins := len(input);
    minNumber := 0;
    minNumberIdx := []int {0, 0};
    for i := 0 ; i < len(idxs); i++ {
        if (idxs[i] != -1) {
            if (idxs[i] < mins) {
                mins = idxs[i];
                minNumber = i + 1;
            }
        }
    }
    minNumberIdx[0] = minNumber;
    minNumberIdx[1] = mins;

    maxNumber := 0;
    maxs := -1; 
    maxNumberIdx := []int {0, 0};
    for i := 0 ; i < len(idxs); i++ {
        if (idxs[i] != -1) {
            if (idxs[i] > maxs) {
                maxs = idxs[i];
                maxNumber = i + 1;
            }
        }
    }
    maxNumberIdx[0] = maxNumber;
    maxNumberIdx[1] = maxs;

    return minNumberIdx, maxNumberIdx;
}

func main() {
    var total = 0;
    file , err := os.Open("input.txt");
    var lines []string;
    if (err != nil) {
            log.Println("Error reading file");
    }
    scanner := bufio.NewScanner(file);
    for scanner.Scan() {
        lines = append(lines,scanner.Text());
    }
    var strLen = len(lines);
    for i:= 0; i < strLen; i++ {
        curr := lines[i];
        currLen := len(curr);
        first := 0;
        last := 0;
        minDigit := -1;
        maxDigit := currLen;
        for j:= 0; j < currLen; j ++ {
            // Find the first digit by looping from the front
            isNumber, num:= isDigit(string(curr[j]));
            if (isNumber) {
                first = num;
                minDigit = j;
                break;
            }
        }
        for j:= currLen - 1; j >= 0 ; j --{
            // Last digit by looping from the back
            isNumber, num:= isDigit(string(curr[j]));
            if (isNumber) {
                last = num;
                maxDigit = j;
                break;
            }
        }
        mins, maxs := searchFirst(curr);
        // [0] -> angkanya, [1] -> indexny
        minWordIdx := mins[1];
        maxWordIdx := maxs[1];
        fmt.Println(minWordIdx, maxWordIdx, minDigit, maxDigit);

        dict := map[int]int {minWordIdx : mins[0], 
        maxWordIdx: maxs[0], minDigit : first,
        maxDigit : last};

        delete(dict, -1);
        delete(dict, currLen);

        keys := make([]int, 0, len(dict));
        for k := range dict {
            keys = append(keys, k);
        }
        sort.Ints(keys);
        res := []int{};
        for _, k := range keys {
            res = append(res, dict[k]);
        }

        first = res[0];
        last = res[len(res) - 1];
        tmpFirst := strconv.Itoa(first);
        tmpLast := strconv.Itoa(last);
        twoDigitNumber := tmpFirst + tmpLast;
        fmt.Println(twoDigitNumber);
        ans , err := strconv.Atoi(twoDigitNumber);
        if (err != nil) {
            // Not a number
        }
        total += ans;
    }
    fmt.Println(total);

}
