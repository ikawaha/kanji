package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	// https://ja.wikipedia.org/wiki/%E4%BA%BA%E5%90%8D%E7%94%A8%E6%BC%A2%E5%AD%97
	jinmeiHTML = "../../../testdata/jinmei-wikipedia_202108.html"
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

var (
	variantP = regexp.MustCompile(`.\s*?（.）`)
)

func MakeGolden(w io.Writer, werr io.Writer) error {
	r, err := OpenGoldenSrc(jinmeiHTML)
	if err != nil {
		return fmt.Errorf("cannot open golden src: %w", err)
	}
	defer r.Close()
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return fmt.Errorf("invalid document src: %w", err)
	}
	var forName []string
	doc.Find(".mw-parser-output > dl:nth-child(39) > dd:nth-child(1) > span:nth-child(2) > span:nth-child(1)").Each(func(_ int, s *goquery.Selection) {
		list := strings.ReplaceAll(s.Text(), "\n", "")
		forName = strings.Fields(list)
	})
	var variants []string
	doc.Find(".mw-parser-output > dl:nth-child(42) > dd:nth-child(1) > span:nth-child(2) > span:nth-child(1)").Each(func(_ int, s *goquery.Selection) {
		list := strings.ReplaceAll(s.Text(), "\n", "")
		variants = variantP.FindAllString(list, -1)
	})

	io.WriteString(w, "!!!人名用漢字\n")
	for _, v := range forName {
		io.WriteString(w, v+"\n")
	}
	io.WriteString(w, "!!!人名用異体字\n")
	for _, v := range variants {
		io.WriteString(w, v+"\n")
	}

	return nil
}
