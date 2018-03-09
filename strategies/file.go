package strategies

import (
	"bufio"
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type OutputFile struct{}

func (OutputFile) Apply(container map[string]map[string]string) {
	f, err := os.Create(".secret")
	check(err)
	defer f.Close()

	w := bufio.NewWriter(f)
	for env, key := range container {
		fmt.Fprintf(w, "%s: \n", env)
		for i, value := range key {
			fmt.Fprintf(w, "\t%s : \"%s\"\n", i, value)
		}
	}
	w.Flush()
}
