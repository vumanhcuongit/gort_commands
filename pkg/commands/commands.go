package commands

import "fmt"

func Echo(input string) string {
	return fmt.Sprintf("%s - sent from Cuong with love", input)
}
