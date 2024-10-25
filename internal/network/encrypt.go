package network

import (
	"fmt"
)

func entrypt(s string) string {
	// ASCII XOR 0x77
	result := ""
	for i := 0; i < len(s); i++ {
		result += fmt.Sprintf("%02x", s[i]^0x77)
	}
	fmt.Println(result)
	return result
}
