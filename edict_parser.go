package edict_parser

import (
	"bufio"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func ParseEdict(filePath string) (Word, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return Word{}, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			return Word{}, nil
		} else if err != nil {
			return Word{}, err
		}

		text, err := ioutil.ReadAll(transform.NewReader(strings.NewReader(line), japanese.EUCJP.NewDecoder()))

		if err != nil {
			return Word{}, err
		}

		println(string(text))
	}
}
