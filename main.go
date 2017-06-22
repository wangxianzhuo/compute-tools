package main

import (
	"flag"
	"fmt"

	"github.com/wangxianzhuo/tools/tools"
)

func main() {
	toolName := flag.String("tool", "help", "[help | htof | ftoh]\n\thelp: how to use\n\thtof: hex to ieee754 float32\n\thtof: ieee754 float32 to hex")
	param := flag.String("param", "", "param to use")
	config := flag.String("config", "", "configs")
	flag.Parse()

	output := ""
	switch *toolName {
	case "htof":
		tools.HexToIEEE754(*param, &output, *config)
		fmt.Printf("float : %s\n", output)
	case "ftoh":
		tools.IEEE754ToHex(*param, &output, *config)
		fmt.Printf("hex : 0x%X\n", output)
	}
}
