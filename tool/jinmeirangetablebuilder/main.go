package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const goldenFilePath = "../../testdata/golden_jinmei.txt"

func main() {
	forName, err := loadRunesFromGolden(goldenFilePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(os.Stderr, "人名用漢字: %d 字\n", len(forName))

	forNameRangeTable := RangeTable(forName)
	DumpRangeTable(os.Stdout, "DesignatedForPersonalNamesRangeTable", forNameRangeTable)
}

// 人名漢字(651字：633字+異体字18字), 人名用異体字(212字)
func loadRunesFromGolden(path string) (forName []rune, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	s.Scan()
	if line := s.Text(); !strings.HasPrefix(line, "!!!") {
		return nil, fmt.Errorf("invalid file format, %s", line)
	}
	i := 1
	for s.Scan() {
		line := s.Text()
		if strings.HasPrefix(line, "!!!") {
			break
		}
		runes := []rune(line)
		if len(runes) < 1 {
			log.Println("empty line, line no:", i)
			continue
		}
		for _, v := range runes {
			if v == '‐' {
				continue
			}
			forName = append(forName, v)
		}
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	for s.Scan() {
		line := s.Text()
		runes := []rune(line)
		if len(runes) < 1 {
			log.Println("empty line, line no:", i)
			continue
		}
		forName = append(forName, runes[0])
	}
	return forName, s.Err()
}
