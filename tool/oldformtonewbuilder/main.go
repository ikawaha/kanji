package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run() error {
	f, err := os.Open("../../testdata/golden_old-new.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	var (
		line int
		b    bytes.Buffer
	)
	b.WriteString("var oldnew = []string{\n")
	for s.Scan() {
		l := s.Text()
		line++
		if strings.HasPrefix(l, "!") {
			continue
		}
		a := strings.Split(l, " ")
		if len(a) != 4 {
			return fmt.Errorf("invalid golden file format, line=%d, %s", line, l)
		}
		b.WriteString(`"\u` + a[0] + `", "\u` + a[2] + `", //` + a[1] + ", " + a[3] + "\n")
	}
	if err := s.Err(); err != nil {
		return err
	}
	b.WriteString("}")
	fmt.Println(b.String())
	return nil
}
