package go_rls

import (
	"bufio"
	"strings"
	"os"
	"fmt"
)

func Get(message string) string {
	fmt.Print(message)
	reader := bufio.NewReader(os.Stdin)
	i, e := reader.ReadString('\n')
	if e != nil {
		return e.Error()
	}
	i = strings.TrimSuffix(i, "\n")
	return i
}

func main() {
	o := Get("Hello? ")
	fmt.Println(o)
}