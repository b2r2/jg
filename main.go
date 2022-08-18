package main

import (
	"fmt"
	"log"
	"strconv"
	_ "strconv"
	"strings"
	_ "strings"
)

func main() {
	var s = "a5bc2d4e"
	var sb strings.Builder
	var prev string

	for i := 0; i < len(s); i++ {
		var str = string(s[i])

		if _, err := strconv.Atoi(str); err != nil {
			sb.WriteString(str)
			prev = string(s[i])
		} else {
			ch, err := strconv.Atoi(str)
			if err != nil {
				log.Fatal(err)
			} else {
				sb.WriteString(strings.Repeat(prev, ch))
			}
		}
	}
	fmt.Println(sb.String())
}
