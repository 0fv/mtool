package mtool

import (
	"os"
	"strings"
)

func ReadCSVTab(file string, process func(rows []string) bool) error {
	return ReadCSVSetSplit(file, "\t", process)
}

func ReadCSVComma(file string, process func(rows []string) bool) error {
	return ReadCSVSetSplit(file, ",", process)
}
func ReadCSVSetSplit(file string, colsep string, process func(rows []string) bool) error {
	filedata, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	for _, row := range strings.Split(string(filedata), "\n") {
		cols := strings.Split(row, colsep)
		if !process(cols) {
			return nil
		}
	}
	return nil
}
