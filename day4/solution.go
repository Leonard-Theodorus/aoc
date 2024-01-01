package main

import (
    "fmt"
    "crypto/md5"
    "encoding/hex"
    s "strings"
    sc "strconv"
)

func getMD5Hash(str string) string{
    hash := md5.Sum([]byte(str))
    return hex.EncodeToString(hash[:])
}

func main(){
    const INPUT string = "bgvyzdsv"
    p := fmt.Println
    var num int = 0
    for {
	temp := INPUT + sc.Itoa(num) 
	p(temp)
	hash := getMD5Hash(temp)
	p(hash)
	if s.HasPrefix(hash, "000000"){
	    break
	} else {
	    num += 1
	}
    }
    p(num)
    
}
