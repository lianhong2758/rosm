package rosm

import ()

type suits struct {
	S map[string]int //字符串key计数
	H map[int]int    //int key计数
}

// 一次性统计套装
func RecordSuit(s ...int) map[int]int {
	maps := make(map[int]int)
	for _, v := range s {
		maps[v]++
	}
	maps[0] = 0 //0不能作为套装id
	return maps
}

// 分次数统计套装
func MakeRecordSuit() (su *suits) {
	su = new(suits)
	return
}

// S计数
func (su *suits) AddS(str string) {
	su.S[str]++
}

// H计数
func (su *suits) AddH(i int) {
	su.H[i]++
}

func (su *suits) GetS() map[string]int {
	su.S[""] = 0
	return su.S
}

func (su *suits) GetH() map[int]int {
	su.H[0] = 0
	return su.H
}
