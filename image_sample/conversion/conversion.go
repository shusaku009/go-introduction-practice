package conversion

import (
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)


const (
	JPEG = "jpeg"
	JPG = "jpg"
	GIF = "gif"
	PNG = "png"
)

func ExtensionCheck(ext string) error {
	switch ext {
	case JPEG, JPG, GIF, PNG:
		return nil
	default:
		return fmt.Errorf("不正な拡張子です: %s", ext)
	}
}

func FilepathCheck(filepath string) error {
	switch filepath {
	case "":
		return errors.New("ファイルパスが指定されていません")
	default:
		if f, err := os.Stat(imagepath); os.IsNotExist(err) || f.IsDir() {
			return errors.New("ファイルが存在しません" + imagepath)
		} else {
			return nil
		}
	}
}

func DirpathCheck(dirpath string) error {
	switch dirpath {
	case "":
		return errors.New("変換後のファイル名が指定されていません")
	default:
		return nil
	}
}

func FileExtCheck(imagepath string) error {
	switch imagepath {
		case "." + JPEG, "." + JPG, "." + GIF, "." + PNG:
			return nil
		default:
			return errors.New("不正な拡張子です: " + imagepath)
	}
}

func FileExtension(extension string, imagepath string, dirpath string) error {
	exFile, err := os.Open(imagepath)
	defer exFile.Close()
	if err != nil {
		return errors.New("os.Create失敗")
	}

	output, err := os.Create(dirpath)
	defer output.Close()
	if err != nil {
		return errors.New("os.Create失敗")
	}

	img, _, Err := image.Decode(exFile)
	if Err != nil {
		return errors.New("画像のデコードに失敗しました")
	}

	switch extension {
	case JPEG, JPG:
		err = jpeg.Encode(output, img, nil)
		if err != nil {
			return errors.New("jpeg.Encode失敗")
		}
		fmt.Println("jpeg変換成功")
		return nil
	case GIF:
		err = gif.Encode(output, img, nil)
		if err != nil {
			return errors.New("gif.Encode失敗")
		}
		fmt.Println("gif変換成功")
		return nil
	case PNG:
		err = png.Encode(output, img)
		if err != nil {
			return errors.New("png.Encode失敗")
		}
		fmt.Println("png変換成功")
		return nil
	}

	return nil
}