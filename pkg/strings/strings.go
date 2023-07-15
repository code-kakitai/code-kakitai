package strings

import "strings"

// ハイフンなしの文字列を返す
func RemoveHyphen(s string) string {
	return strings.ReplaceAll(s, "-", "")
}
