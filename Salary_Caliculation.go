package main

import (
	"fmt"
	"strings"
)

func main() {
	Item := []string{"個別時間数", "補講時間数", "業務時間数", "通勤手当", "手当", "割り増し手当"}

	// Salary := []float64{個別給,補講給,業務給,以下は実費換算のため"1"とする}
	Salary := []float64{1230, 1120, 1120, 1, 1, 1}

	var chief int = 3000
	total_salary := []int{chief}

	var variable_input [6]float64

	for input := 0; input < len(Item); input++ {
		fmt.Println(Item[input] + "を入力してください")
		fmt.Scan(&variable_input[input])
		result := int(variable_input[input] * Salary[input])
		total_salary = append(total_salary, result)
	}

	sum := 0

	for _, value := range total_salary {
		sum += value
	}
	fmt.Println("合計：" + comma(sum) + "円")
}

// comma関数は数値をカンマ区切りの文字列に変換します
func comma(sum int) string {
	sum_str := fmt.Sprintf("%d", sum)
	length_of_sum := len(sum_str)

	// 文字列が3桁以下の場合、そのまま返す
	if length_of_sum <= 3 {
		return sum_str
	}

	// 結果を格納するビルダー
	var result_madecomma strings.Builder

	pre := length_of_sum % 3

	// 先頭の行を追加
	if pre > 0 {
		result_madecomma.WriteString(sum_str[:pre])
		if length_of_sum > pre {
			result_madecomma.WriteString(",")
		}
	}

	// 3桁ごとにカンマを追加
	for i := pre; i < length_of_sum; i += 3 {
		result_madecomma.WriteString(sum_str[i : i+3])
		if i+3 < length_of_sum {
			result_madecomma.WriteString(",")
		}
	}
	return result_madecomma.String()
}

// ビルダーとWriteStringの機能を調べる
