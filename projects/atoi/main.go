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
    s2 := ""
	a := strings.Split(s, "")
    for i:=0; i < len(a); i++{
        var num, _ = strconv.Atoi(a[i])
        s2 += strconv.Itoa(num * num)
    }
    fmt.Println(s2)
}
