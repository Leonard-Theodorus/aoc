package main;

import (
    "fmt"
    "log"
    "os"
    "bufio"
    "regexp"
    "strings"
    "strconv"
)

func check(color string, number string) bool {
    num, err:= strconv.Atoi(number);
    res := true;
    if (err != nil){

    }
    if (color == "blue" && num > 14) {
        res = false;
    } else if (color == "red" && num > 12){
        res = false;
    } else if (color == "green" && num > 13) {
        res = false;
    }
    return res;
}


func main () {
    file, err := os.Open("input.txt");
    total := 0;
    if (err != nil) {
        log.Fatal("file broken");
    }
    scanner := bufio.NewScanner(file);
    var lines []string;

    for scanner.Scan() {
        lines = append(lines, scanner.Text());
    }
    gameRe := regexp.MustCompile(`Game .*:`);
    restOfStringRe := regexp.MustCompile(`:.*`);

    gamePattern := regexp.MustCompile(`([0-9]+) (blue|red|green)`);
    for i := 0; i < len(lines); i++ {
        //lineValid := true;
        gameStr := gameRe.FindAllString(lines[i], -1)[0];
        ballStr := restOfStringRe.FindAllString(lines[i], -1)[0];
        gameStr = strings.Trim(gameStr, ":");
        //currId := strings.Trim(gameStr, "Game ");

        ballStr = strings.TrimLeft(ballStr, ": ");
        games := strings.Split(ballStr, ";");
        maxRed := 0;
        maxBlue := 0;
        maxGreen := 0;
        for j := range games {
            currGame := games[j];
            fmt.Println(currGame);
            matches := gamePattern.FindAllStringSubmatch(currGame, -1);
            for k := range matches {
                currRound := matches[k];
                currNumber := currRound[1];
                currColor := currRound[2];

                if (currColor == "red") {
                    num, err := strconv.Atoi(currNumber);
                    if (err != nil ){

                    }
                    maxRed = max(maxRed, num);
                } else if (currColor == "green") {
                    num, err := strconv.Atoi(currNumber);
                    if (err != nil ){

                    }
                    maxGreen = max(maxGreen, num);
                } else if (currColor == "blue") {
                    num, err := strconv.Atoi(currNumber);
                    if (err != nil ){

                    }
                    maxBlue = max(maxBlue, num);
                }

                //if (!check(currColor, currNumber)) {
                    //lineValid = false;
                    //break;
                //}
            }
        }
        cube := maxRed * maxGreen * maxBlue;
        total += cube;
        //if (lineValid) {
            //fmt.Println(i);
            //id, err := strconv.Atoi(currId);
            //if (err != nil){

            //}
            //total += id;
        //}
    }

    fmt.Println(total);
}
