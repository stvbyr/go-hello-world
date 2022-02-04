package readfile

import (
	"io/ioutil"
)

func ReadMarkdownFile(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
