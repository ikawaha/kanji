package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const csvFilePath = "../../testdata/golden_jyouyou_H22-11-30.csv"

func main() {
	standard, oldform, tolerable, err := loadRunesFromCSV(csvFilePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(os.Stderr, "標準字体: %d 字, 旧字体: %d 字, 許容字体: %d 字\n", len(standard), len(oldform), len(tolerable))

	standardTable := RangeTable(standard)
	DumpRangeTable(os.Stdout, "StandardRegularUseRangeTable", standardTable)

	oldFormTable := RangeTable(oldform)
	DumpRangeTable(os.Stdout, "OldFormRegularUseRangeTable", oldFormTable)

	tolerableTable := RangeTable(tolerable)
	DumpRangeTable(os.Stdout, "TolerableRegularUseRangeTable", tolerableTable)

}

// 標準字体, 旧字体, 許容字体
func loadRunesFromCSV(path string) (standard []rune, oldform []rune, tolerable []rune, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, nil, err
	}
	defer f.Close()
	s := bufio.NewScanner(f)

	i := 0
	for s.Scan() {
		row := strings.Split(s.Text(), "\t")
		i++
		if len(row) < 1 {
			log.Println("empty line, line no:", i)
			continue
		}
		col0 := []rune(row[0])
		for i, v := range col0 {
			if v == '［' || v == '］' || v == '（' || v == '）' { // 餅［餅］（餠）
				continue
			}
			if i == 0 {
				standard = append(standard, v)
			} else if col0[i-1] == '（' {
				oldform = append(oldform, v)
			} else if col0[i-1] == '［' {
				tolerable = append(tolerable, v)
			}
		}
	}
	return standard, oldform, tolerable, s.Err()
}
