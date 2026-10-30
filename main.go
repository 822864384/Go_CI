package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	amount := flag.Int("amount", 386, "待拆分的金额（元，非负整数）")
	flag.Parse()

	result, err := Breakdown(*amount)
	if err != nil {
		log.Fatalf("拆分失败: %v", err)
	}

	fmt.Printf("%d 元可拆分为：\n", *amount)
	for _, note := range notes {
		if count, ok := result[note]; ok {
			fmt.Printf("  %3d 元 × %2d 张\n", note, count)
		}
	}
	fmt.Printf("合计校验: %d 元\n", Total(result))
}
