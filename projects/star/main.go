package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	a := strings.Split(s, "")
	s2 := strings.Join(a, "*")
	fmt.Println(s2)
}
