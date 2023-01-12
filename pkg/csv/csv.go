package csv

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func CsvToStr() ([][]string, error) {
	data := make([][]string, 0)

	file, err := os.Open("ueba.csv")
	if err != nil {
		return nil, err
	}

	r := bufio.NewReader(file)

	// skip first line
	_, err = r.ReadString('\n')
	if err != nil {
		if err != io.EOF {
			return nil, err
		}
		return data, nil
	}

	for i := 0; ; i++ {
		data = append(data, make([]string, 0))
		line, err := r.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				return nil, err
			}
			str := strings.Split(line, ",")
			data[i] = append(data[i], str[1:]...)
			return data, nil
		}
		str := strings.Split(line[:len(line)-2], ",")
		data[i] = append(data[i], str[1:]...)
	}
}
