// Package main 实现金额面额拆分命令行工具。
package main

import (
	"errors"
	"fmt"
)

// ErrNegativeAmount 表示传入了负数金额。
var ErrNegativeAmount = errors.New("denom: amount must not be negative")

// notes 是可用的纸币面额（单位：元），必须从大到小排列。
var notes = []int{100, 50, 20, 10, 5, 1}

// Breakdown 把 amount 贪心拆分为各面额张数。
// 返回的 map 只包含张数大于 0 的面额；张数为 0 的面额不写入。
// amount 为负数时返回 ErrNegativeAmount。
func Breakdown(amount int) (map[int]int, error) {
	if amount < 0 {
		return nil, fmt.Errorf("%w: got %d", ErrNegativeAmount, amount)
	}

	result := make(map[int]int, len(notes))
	rest := amount
	for _, note := range notes {
		count := rest / note
		if count > 0 {
			result[note] = count
			rest -= count * note
		}
	}

	return result, nil
}

// Total 校验拆分结果的面额合计，返回总金额。
func Total(result map[int]int) int {
	sum := 0
	for note, count := range result {
		sum += note * count
	}
	return sum
}
