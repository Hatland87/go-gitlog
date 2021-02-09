package terminal

import (
	"fmt"
	"os/exec"
	"strings"
)

func Run(arg string) string {

	arr := strings.Split(arg, " ")
	arg1 := arr[0]
	args := arr[1:]
	out, err := exec.Command(arg1, args...).CombinedOutput()
	if err != nil {
		fmt.Println(string(out[:]), err)
	} else {
		fmt.Println("Success!")
	}

	return string(out[:])
}
