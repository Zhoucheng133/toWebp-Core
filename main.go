package main

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"towebp_core/utils"

	"github.com/ncruces/zenity"
)
import "C"

//export Convert
func Convert(path *C.char, width, height *C.int, output *C.char) *C.char {
	err := utils.Convert(C.GoString(path), int(C.int(*width)), int(C.int(*height)), C.GoString(output))
	if err != nil {
		return C.CString(fmt.Sprint("ERR: ", err.Error()))
	}
	return C.CString("OK")
}

//export Scan
func Scan(path *C.char) *C.char {
	jsonData, err := json.Marshal(utils.Scan(C.GoString(path)))
	if err != nil {
		return C.CString(fmt.Sprint("ERR: ", err.Error()))
	}
	return C.CString(string(jsonData))
}

//export ConvertFromDir
func ConvertFromDir(path *C.char, width, height *C.int, output *C.char) *C.char {
	err := utils.ConvertFromDir(C.GoString(path), int(C.int(*width)), int(C.int(*height)), C.GoString(output))
	if err != nil {
		return C.CString(fmt.Sprint("ERR: ", err.Error()))
	}
	return C.CString("OK")
}

func main() {

	inputPath, err := zenity.SelectFile(
		zenity.Title("选择文件"),
		zenity.FileFilters{
			{Name: "图片文件", Patterns: []string{"*.jpeg", "*.jpg", "*.png"}},
		},
	)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	baseName := filepath.Base(inputPath)
	ext := filepath.Ext(baseName)
	outName := baseName[:len(baseName)-len(ext)] + ".webp"
	utils.Convert(inputPath, 0, 1000, outName)
}
