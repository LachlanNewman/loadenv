package bigl_loadenv

import (
	"fmt"
	"os"
)

func readFile(fileName string) {
	file, err := os.Open(fileName)
	// defer runs after function has finished
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(file.Stat())
}
