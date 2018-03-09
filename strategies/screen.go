package strategies

import "fmt"

type OutputScreen struct{}

func (OutputScreen) Apply(container map[string]map[string]string) {
	for env, key := range container {
		fmt.Printf("%s: \n", env)
		for i, value := range key {
			fmt.Printf("\t%s : \"%s\"\n", i, value)
		}
	}
}
