package imageconversion

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// 画像ファイルのディレクトリをコマンドラインから取得する
	dir := flag.String("dir", "", "画像ファイルのディレクトリ")
	flag.Parse()
	// ディレクトリが指定されていない場合はエラーを返す
	if *dir == "" {
		fmt.Println("ディレクトリが指定されていません")
		return
	}

	// ディレクトリ内のファイルを取得
	f, _ := os.Open(*dir)
	fmt.Println(f)
	// ファイルの拡張子がJPGか確認

	// JPGファイルをPNGファイルに変換

	// 変換後のファイルを保存
	// 上記の処理をfor文で繰り返す
}
