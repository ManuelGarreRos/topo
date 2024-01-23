package CLI

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AskForAuthorize(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	return strings.ToLower(strings.TrimSpace(scanner.Text()))
}
