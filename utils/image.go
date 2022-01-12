package utils

import (
	"encoding/base64"
	"io/ioutil"
	"strings"
)

//Base64ToImg base64 code to img
func Base64ToImg(base64Str, path string) error {
	ImageSource, _ := base64.StdEncoding.DecodeString(strings.Replace(base64Str, "data:image/png;base64,", "", 1)) //成图片文件并把文件写入到buffer
	err := ioutil.WriteFile(path, ImageSource, 0666)                                                               //buffer输出到jpg文件中（不做处理，直接写到文件）
	return err
}

//ImgToBase64 img to base64 code
func ImgToBase64() {

}
