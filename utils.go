package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/fatih/color"
	"sort"
	"strconv"
	"errors"
)

var mahjong = [...]string{
	"1m", "2m", "3m", "4m", "5m", "6m", "7m", "8m", "9m",
	"1p", "2p", "3p", "4p", "5p", "6p", "7p", "8p", "9p",
	"1s", "2s", "3s", "4s", "5s", "6s", "7s", "8s", "9s",
	"1z", "2z", "3z", "4z", "5z", "6z", "7z",
}

var mahjongZH = [...]string{
	"1万", "2万", "3万", "4万", "5万", "6万", "7万", "8万", "9万",
	"1饼", "2饼", "3饼", "4饼", "5饼", "6饼", "7饼", "8饼", "9饼",
	"1索", "2索", "3索", "4索", "5索", "6索", "7索", "8索", "9索",
	"东", "南", "西", "北", "白", "发", "中",
}

func _errorExit(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
	os.Exit(1)
}

// e.g. "3m" => 2
func _convert(tile string) (int, error) {
	for i, m := range mahjong {
		if m == tile {
			return i, nil
		}
	}
	return -1, fmt.Errorf("参数错误: %s", tile)
}

// e.g. "13m 24p" => (4, [0, 2, 10, 12])
func convert(tiles string) (num int, cnt []int, err error) {
	tiles = strings.TrimSpace(tiles)
	if tiles == "" {
		return 0, nil, errors.New("参数错误: 处理的手牌不能为空")
	}

	var result []int
	for _, split := range strings.Split(tiles, " ") {
		split = strings.TrimSpace(split)
		if len(split) <= 1 {
			return 0, nil, errors.New("参数错误: " + split)
		}
		for i := range split[:len(split)-1] {
			single := split[i:i+1] + split[len(split)-1:]
			tile, err := _convert(single)
			if err != nil {
				return -1, nil, err
			}
			result = append(result, tile)
		}
	}

	cnt = make([]int, len(mahjong))
	for _, m := range result {
		cnt[m]++
		if cnt[m] > 4 {
			return 0, nil, errors.New("参数错误: 超过4张一样的牌！")
		}
	}

	return len(result), cnt, nil
}

func countToString(cnt []int) (string, error) {
	if len(cnt) != len(mahjong) {
		return "", fmt.Errorf("cnt 长度必须为 %d", len(mahjong))
	}

	sb := strings.Builder{}
	for i, type_ := range [...]string{"m", "p", "s", "z"} {
		wrote := false
		for j := 0; j < 9; j++ {
			idx := 9*i + j
			if idx >= len(mahjong) {
				break
			}
			for k := 0; k < cnt[idx]; k++ {
				sb.WriteString(strconv.Itoa(j + 1))
				wrote = true
			}
		}
		if wrote {
			sb.WriteString(type_ + " ")
		}
	}
	return strings.TrimSpace(sb.String()), nil
}

func in(a string, arr []string) bool {
	for _, _a := range arr {
		if _a == a {
			return true
		}
	}
	return false
}

func uniqueStrings(strings []string) []string {
	u := make([]string, 0, len(strings))
	mp := make(map[string]struct{}, len(strings))

	for _, val := range strings {
		if _, ok := mp[val]; !ok {
			mp[val] = struct{}{}
			u = append(u, val)
		}
	}

	sort.Strings(u)
	return u
}

func countDui(cnt []int) int {
	dui := 0
	for i := 0; i < len(mahjong); i++ {
		if cnt[i] >= 2 {
			dui++
		}
	}
	return dui
}

//

func _getAttr(n float64) []color.Attribute {
	var colors []color.Attribute
	switch {
	case n < 14:
		colors = append(colors, color.FgBlue)
	case n <= 18:
		colors = append(colors, color.FgYellow)
	case n < 24:
		colors = append(colors, color.FgHiRed)
	default:
		colors = append(colors, color.FgRed)
	}
	return colors
}

func colorNumber1(n int) {
	color.New(_getAttr(float64(n))...).Printf("%d", n)
}

func colorNumber2(n int) {
	color.New(_getAttr(float64(n) / 2)...).Printf("%d", n)
}

func alert(s string) {
	color.Yellow(s)
	fmt.Println()
}
