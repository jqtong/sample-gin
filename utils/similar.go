package utils

import "strings"

// SimilarStr return the len of longest string both in str1 and str2 and the positions in str1 and str2
func SimilarStr(str1 []rune, str2 []rune) (int, int, int) {
	var maxLen, tmp, pos1, pos2 = 0, 0, 0, 0
	len1, len2 := len(str1), len(str2)

	for p := 0; p < len1; p++ {
		for q := 0; q < len2; q++ {
			tmp = 0
			for p+tmp < len1 && q+tmp < len2 && str1[p+tmp] == str2[q+tmp] {
				tmp++
			}
			if tmp > maxLen {
				maxLen, pos1, pos2 = tmp, p, q
			}
		}

	}

	return maxLen, pos1, pos2
}

// SimilarChar return the total length of longest string both in str1 and str2
func SimilarChar(str1 []rune, str2 []rune) int {
	maxLen, pos1, pos2 := SimilarStr(str1, str2)
	total := maxLen

	if maxLen != 0 {
		if pos1 > 0 && pos2 > 0 {
			total += SimilarChar(str1[:pos1], str2[:pos2])
		}
		if pos1+maxLen < len(str1) && pos2+maxLen < len(str2) {
			total += SimilarChar(str1[pos1+maxLen:], str2[pos2+maxLen:])
		}
	}

	return total
}

// SimilarText return a int value in [0, 100], which stands for match level
func SimilarText(str1 string, str2 string) int {
	txt1, txt2 := []rune(str1), []rune(str2)
	if len(txt1) == 0 || len(txt2) == 0 {
		return 0
	}
	return SimilarChar(txt1, txt2) * 200 / (len(txt1) + len(txt2))
}

// LevenshteinDistance text distance with levenshitein algo
func LevenshteinDistance(s1, s2 string) int {
	if s1 == s2 {
		return 0
	}
	s1Array := strings.Split(s1, "")
	s2Array := strings.Split(s2, "")
	lenS1Array := len(s1Array)
	lenS2Array := len(s2Array)
	if lenS1Array == 0 {
		return lenS2Array
	}
	if lenS2Array == 0 {
		return lenS1Array
	}
	m := make([][]int, lenS1Array+1)
	for i := range m {
		m[i] = make([]int, lenS2Array+1)
	}
	for i := 0; i < lenS1Array+1; i++ {
		for j := 0; j < lenS2Array+1; j++ {
			if i == 0 {
				m[i][j] = j
			} else if j == 0 {
				m[i][j] = i
			} else {
				if s1Array[i-1] == s2Array[j-1] {
					m[i][j] = m[i-1][j-1]
				} else {
					m[i][j] = Min(m[i-1][j]+1, m[i][j-1]+1, m[i-1][j-1]+1)
				}
			}
		}
	}
	return m[lenS1Array][lenS2Array]
}

// Min returns the minimum number of passed int slices.
func Min(is ...int) int {
	var min int
	for i, v := range is {
		if i == 0 || v < min {
			min = v
		}
	}
	return min
}

// Max returns the maximum number of passed int slices.
func Max(is ...int) int {
	var max int
	for _, v := range is {
		if max < v {
			max = v
		}
	}
	return max
}
