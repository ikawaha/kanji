package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	// https://www.bunka.go.jp/kokugo_nihongo/sisaku/joho/joho/kijun/naikaku/kanji/joyokanjisakuin/index.html
	jyouyouHTML = "../../../testdata/jyouyou_H22-11-30.html"
)

func OpenGoldenSrc(path string) (io.ReadCloser, error) {
	if !strings.HasPrefix(path, "https://") {
		return os.Open(path)
	}
	res, err := http.Get(path)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

func MakeGolden(w io.Writer, werr io.Writer) error {
	r, err := OpenGoldenSrc(jyouyouHTML)
	if err != nil {
		return fmt.Errorf("cannot open golden src: %w", err)
	}
	defer r.Close()
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return fmt.Errorf("invalid document src: %w", err)
	}
	var records [][]string
	doc.Find("table#urlist.display").Each(func(_ int, s *goquery.Selection) {
		s.Find("tr").Each(func(_ int, s *goquery.Selection) {
			var record []string
			s.Find("td").Each(func(_ int, s *goquery.Selection) {
				record = append(record, s.Text())
			})
			records = append(records, record)
		})
	})
	for i, record := range records {
		if _, err := io.WriteString(w, strings.Join(record, "\t")); err != nil {
			return fmt.Errorf("cannot write, record=%d, %w", i, err)
		}
		if _, err := io.WriteString(w, "\n"); err != nil {
			return fmt.Errorf("cannot write, record=%d, %w", i, err)
		}
	}
	return nil
}
