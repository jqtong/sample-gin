package utils

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"
	"unicode"
)

// FilterSpace 将检索词中中文部分内容的空格替换为英文逗号，将中文逗号更改为英文逗号
// 用于用户输入查询关键词的清洗
func FilterSpace(s string) string {

	keywords := strings.TrimSpace(s)
	keywords = strings.ReplaceAll(keywords, "，", ",")

	var sb strings.Builder
	tokens := strings.Split(keywords, ",")

	for _, tokenStr := range tokens {
		innerTokens := strings.Split(tokenStr, " ")
		for _, token := range innerTokens {

			token = strings.TrimSpace(token)
			if token == "" || token == " " {
				continue
			}

			isMBString := false
			for _, r := range token {
				switch {
				case unicode.Is(unicode.Han, r): // 汉字
					isMBString = true
					break
				case unicode.Is(unicode.Hiragana, r): // 平假名
					isMBString = true
					break
				case unicode.Is(unicode.Katakana, r): // 片假名
					isMBString = true
					break
				}

				if unicode.Is(unicode.Han, r) || unicode.Is(unicode.Hiragana, r) || unicode.Is(unicode.Katakana, r) {
					isMBString = true
					break
				}
			}
			if isMBString {
				sb.WriteString(", " + token + ", ")
			} else {
				sb.WriteString(token + " ")
			}
		}

		sb.WriteString(",")
	}

	handledTokens := make([]string, 0)
	for _, token := range strings.Split(sb.String(), ",") {
		token = strings.TrimSpace(token)
		token = strings.TrimRight(token, ",")

		if token == "" || token == " " {
			continue
		}

		handledTokens = append(handledTokens, token)
	}

	return strings.Join(handledTokens, ", ")
}

//TopicHash 计算主题词的哈希值
func TopicHash(topic string) string {
	t := sha1.New()
	io.WriteString(t, topic)

	return fmt.Sprintf("%x", t.Sum(nil))
}

//KeywordJSON 返回主题词的json 字符串
func KeywordJSON(topic, subTopic, thirdTopic, except string) string {
	arr := map[string]string{
		"topic":      topic,
		"subtopic":   subTopic,
		"thirdtopic": thirdTopic,
		"except":     except,
	}

	jsonStr, _ := json.Marshal(arr)

	return string(jsonStr[:])
}

// ParseKeyword 解析关键词
func ParseKeyword(keywordJSON string) (keywords map[string]string, err error) {
	keywords = make(map[string]string)
	err = json.Unmarshal([]byte(keywordJSON), &keywords)
	return
}

// StructToMap type struct to map
func StructToMap(obj interface{}) (res map[string]interface{}) {
	j, _ := json.Marshal(obj)
	_ = json.Unmarshal(j, &res)
	return
}

//CurDateTime 返回当前时间
func CurDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
