package tools

import (
	"regexp"
)

func FilterURL(origin string) (dest string) {

	re, _ := regexp.Compile("[a-zA-Z0-9/-/_/:/.//]*")
	one := re.Find([]byte(origin))

	return (string(one))

}
