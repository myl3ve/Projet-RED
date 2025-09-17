package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var in = bufio.NewReader(os.Stdin)

func Input(prompt string) string {
	fmt.Print(prompt)
	txt, _ := in.ReadString('\n')
	return strings.TrimSpace(strings.ToLower(txt))
}

func PressEnter() { _ = Input("\n(Entrée pour continuer) ") }
