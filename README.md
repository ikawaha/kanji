Kanji
===

日本語漢字に関するパッケージです。

## 常用漢字

常用漢字は `一般の社会生活において現代の国語を書き表すための漢字使用の目安` として示される漢字の集合で、このパッケージでは [平成22年内閣告示第2号 (2010年11月30日)](https://www.bunka.go.jp/kokugo_nihongo/sisaku/joho/joho/kijun/naikaku/kanji/index.html) として告示されているものを対象にしています。

このパッケージで「常用漢字」は、標準字体 2136字、旧字体 364字、許容字体 5字からなる集合として扱っています。標準字体、旧字体、許容字体のそれぞれを `unicode.RangeTable` として定義していますので直接利用可能です。また、これらを扱う関数も定義されています。詳しくは [ドキュメント](https://pkg.go.dev/github.com/ikawaha/encoding/kanji) を参照してください。

## 人名用漢字

人名用漢字は、常用漢字以外で子の名に使える漢字の集合のことです。法務省のページに [子の名に使える漢字](http://www.moj.go.jp/MINJI/minji86.html) として定義されています。 人名用漢字を `unicode.RangeTable` として定義していますので直接利用可能です。また、人名に使える漢字であるかどうか（常用漢字であるかまたは人名用漢字であるか）をチェックする関数を用意しています。詳しくは [ドキュメント](https://pkg.go.dev/github.com/ikawaha/encoding/kanji) を参照してください。

