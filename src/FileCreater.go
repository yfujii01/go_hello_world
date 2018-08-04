package src

import (
	"os"
		"fmt"
	"errors"
)

func FileWrite(filename string, mes string) (error) {

	if filename == ""{
		//log.Fatal("ファイル名が空白")
		return errors.New("ファイル名が空白です。")
	}

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		//log.Fatal(err)
		return errors.New("ファイル書き込みエラー")
	}
	defer file.Close()
	fmt.Fprintln(file, mes)

	return nil
}
