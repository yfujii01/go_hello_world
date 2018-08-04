package src

import "go_hello_world/src/utils"

func bar(s1 string, s2 string) string {
	return utils.Join(s1, s2)
}
