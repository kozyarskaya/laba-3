package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
    "strconv"
)

func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	a := strings.Split(s, "")
    maxi:=0
    for i:=0; i  < len(a); i++{
        var num, _ = strconv.Atoi(a[i])
        if num > maxi{
            maxi = num
        }
    }
    fmt.Println(maxi)
}
