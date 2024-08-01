package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "regexp"
)

func isNumber(input string) bool {
    _, err := strconv.Atoi(input);
    if (err != nil) {
        return false;
    }
    return true;
}

func main () {
    total := 0;
    var lines[]string;
    file, err := os.Open("test.txt");
    if (err != nil) {
        // file error
    }
    scanner := bufio.NewScanner(file);
    for scanner.Scan() {
        lines = append(lines, scanner.Text());
    }
    tmpIdx := make(map[[3]int]bool);
    directions := [][]int {
        {-1, 0},
        {0, -1},
        {-1, -1},
        {-1, 1},
        {0, 1},
        {1, 1},
        {1, 0},
        {1, -1},
    }
    for i := 0; i < len(lines); i++ {
        curr := lines[i];
        for j := 0; j < len(curr); j++ {
            if !isNumber(string(curr[j])) && string(curr[j]) != "."{
                // if this is symbol
                for k := range directions {
                    dx := directions[k][0] + j;
                    dy := directions[k][1] + i;

                    if (dx >= 0 && dx < len(curr) && dy >= 0 && 
                    dy < len(lines) ) {
                        inCheck := string(lines[dy][dx]);
                        number := isNumber(inCheck);
                        if (number) {
                            pattern := fmt.Sprintf(`([0-9]*)(%s)([0-9]*)`, inCheck);
                            checkRe := regexp.MustCompile(pattern);
                            idxs := checkRe.FindAllStringIndex(string(lines[dy]), -1);
                            for z := range idxs {
                                start := idxs[z][0];
                                end := idxs[z][1];
                                key := [3]int{dy, start, end};
                                if (tmpIdx[key] == false) {
                                    // coordinate never seen before
                                    tmpIdx[key] = true;
                                    fmt.Println(lines[dy][start:end], start, end);
                                    partNumber, err := strconv.Atoi(lines[dy][start:end]);
                                    fmt.Println(partNumber);
                                    if (err != nil) {

                                    }
                                    total += partNumber;
                                }
                            }
                        }
                    }
                }
            }
        }
    }
    fmt.Println(total);
}
