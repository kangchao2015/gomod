package level3

import "fmt"

func Gretting(name string) (string, error) {
	return fmt.Sprintf("%s ahaha !", name), nil
}