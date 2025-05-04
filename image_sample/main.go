package main

import (
	"flag"
	"fmt"
	"go-introduction-practice/image_sample/conversion"
	"os"
	"path/filepath"
)

var (
	extension string
	imagepath string
	dirpath string
)

func main() {
	flag.StringVar(&extension, "e", "jpeg", "拡張子の指定")
	flag.StringVar(&imagepath, "f", "", "変換前の画像ファイルのパス")
	flag.StringVar(&dirpath, "d", "", "変換後の画像ファイルのパス")
	flag.Parse()

	err := conversion.FileExtension(extension, imagepath, dirpath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = conversion.FilepathCheck(imagepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = conversion.DirpathCheck(dirpath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	f := filepath.Ext(imagepath)
	err = conversion.FileExtCheck(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = conversion.FileExtension(extension, imagepath, dirpath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("変換が完了しました")
}