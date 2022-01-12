package utils

import (
	"fmt"
	"math/rand"
	"regexp"
	"sort"
	"strings"
	"time"
	"unsafe"
)

var src = rand.NewSource(time.Now().UnixNano())

const keywordTag = "<span style='color: %s; font-style: normal;'>${1}</span>"
const keywordTagWithBackground = "<span style='color: %s; background-color: %s; font-style: normal;'>${1}</span>"

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// HighLightColor 高亮标签颜色定义
type HighLightColor struct {
	Color           string
	BackgroundColor string
}

// MatchField 命中词
type MatchField struct {
	Keyword string         `json:"keyword"`
	Freq    int            `json:"freq"`
	Color   HighLightColor `json:"-"`
}

// RandString generate random string
func RandString(n int) string {

	b := make([]byte, n)

	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

// MatchWords 计算词频
func MatchWords(keywords []string, content string) []MatchField {

	// keywords 清洗（截断，滤重）
	keywordMap := make(map[string]bool, 1<<5)
	for _, keyword := range keywords {
		k := strings.TrimSpace(keyword)
		if k != "" && k != " " {
			keywordMap[k] = true
		}
	}

	matchFields := make([]MatchField, 0, 1<<5)

	for keyword := range keywordMap {

		matches := findSingle(keyword, content)

		if matches > 0 {
			matchFields = append(matchFields, MatchField{
				Keyword: keyword,
				Freq:    matches,
			})
		}
	}

	return matchFields
}

// MatchColors 客官，选个颜色?
func MatchColors(keywords []MatchField, colors []HighLightColor) []MatchField {

	if len(keywords) == 0 || len(colors) == 0 {
		return keywords
	}

	index := 0
	for i := 0; i < len(keywords); i++ {
		if index == len(colors) {
			index = 0
		}
		keywords[i].Color = colors[index]
		index++
	}

	return keywords
}

func findSingle(keyword string, content string) int {

	reg, err := regexp.Compile(keyword)
	if err != nil {
		// compile error
		return 0
	}

	matches := reg.FindAllString(content, -1)

	return len(matches)
}

// AppendRainbowFlag 增加彩色高亮标签
func AppendRainbowFlag(keywords []MatchField, content string) string {

	//matchWords := make([]MatchField, 0, len(keywords))
	//copy(matchWords, keywords)

	if len(keywords) > 0 {
		// sort match words by Freq
		sort.SliceStable(keywords, func(i, j int) bool {
			return keywords[i].Freq > keywords[j].Freq
		})
	}

	replacedContent := content
	if len(keywords) > 0 {
		for _, field := range keywords {
			replacedContent = replaceSingle(field.Keyword, field.Color, replacedContent)
		}
	}

	return replacedContent
}

func replaceSingle(keyword string, color HighLightColor, content string) string {

	reg, err := regexp.Compile("(" + keyword + ")")
	if err != nil {
		return content
	}

	if color.BackgroundColor == "" {
		return reg.ReplaceAllString(content, fmt.Sprintf(keywordTag, color.Color))
	}

	return reg.ReplaceAllString(content, fmt.Sprintf(keywordTagWithBackground, color.Color, color.BackgroundColor))
}

//TrimHTML filter html tag
func TrimHTML(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	return strings.TrimSpace(src)
}
