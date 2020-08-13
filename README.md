# find-numbers-with-even-number-of-digits

## 題目解讀：

### 題目來源:
[find-numbers-with-even-number-of-digits](https://leetcode.com/problems/find-numbers-with-even-number-of-digits/)

### 原文:

Given an array nums of integers, return how many of them contain an even number of digits.

### 解讀：

給定一個正整數陣列nums

計算裡面有多陣列元素nums[i]的digits數目是偶數

## 初步解法:
### 初步觀察:

對於這題 有想到兩種作法

作法1: log 

對於每個正整數

可以藉由log10這個 數值的 floor +1 得到該integer的digits數值

舉例來說： log10(10) = 1 => 其 digits = 1 + 1 = 2
         log10(100) = 2 => 其 digits = 2 + 1 = 3

作法2: 

對於每個正整數

可以藉由把數字轉換成字串 然後藉由 計算字串長度得到該integer的digits數值

### 初步設計:
這邊由於for loop部份比較簡單 所以這邊就只敘述怎麼去推算正整數的digit
Given an integer num
**作法1**:

step1: let integer log10Val = Math.Log10(num)

step2: let integer digits = Math.floor(log10Val) + 1

step3: return digits 

**作法2**：
step1: let string strVal = strconv.Itoa(num)

step2: let integer digits = len(strVal)

step3: return digits

## 遇到的困難
### 題目上理解的問題
因為英文不是筆者母語

所以在題意解讀上 容易被英文用詞解讀給搞模糊

### pseudo code撰寫

一開始不習慣把pseudo code寫下來

因此 不太容易把自己的code做解析

### golang table driven test不熟
對於table driven test還不太熟析

所以對於寫test還是耗費不少時間
## 我的github source code
```golang
package find_numbers

import (
	"math"
	"strconv"
)

func findNumbers(nums []int) int {
	result := 0
	for _, val := range nums {
		if calculateDigits(val)%2 == 0 {
			result += 1
		}
	}
	return result
}

func calculateDigits(num int) int {
	return int(math.Floor(math.Log10(float64(num)))) + 1
}

func calculateDigitsV1(num int) int {
	result := strconv.Itoa(num)
	return len(result)
}

```
## 測資的撰寫

```golang
package find_numbers

import "testing"

func Test_findNumbers(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				nums: []int{12, 345, 2, 6, 7896},
			},
			want: 2,
		},
		{
			name: "Example2",
			args: args{
				nums: []int{555, 901, 482, 1771},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumbers(tt.args.nums); got != tt.want {
				t.Errorf("findNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

```
## my self record
[golang ironman30 day 15th day](https://hackmd.io/1HTH-P3uQoqwROS5ESoLMw?view)

## 參考文章

[golang test](https://ithelp.ithome.com.tw/articles/10204692)