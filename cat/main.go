package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

// 引数でファイルパスの一覧を取得する
func main() {
	showLineNumber := flag.Bool("n", false, "行番号を表示する")
	flag.Parse()

	filePaths := flag.Args()

	lineNum := 1

	for _, filePath := range filePaths {
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ファイルをオープンできませんでした: %v\n",filePath, err)
			continue
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if *showLineNumber {
				fmt.Printf("%6d: %s\n", lineNum, line)
				lineNum++
			} else {
				fmt.Println(line)
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "エラー: %s の読み込み中にエラーが発生しました: %v\n", filePath, err)
		}
	}
}