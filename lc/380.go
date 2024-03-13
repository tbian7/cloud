package lc

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	m map[int]int
	s []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		m: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	this.s = append(this.s, val)
	this.m[val] = len(this.s) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	i, ok := this.m[val]

	if ok {
		this.s[i] = this.s[len(this.s)-1]
		this.m[this.s[i]] = i
		this.s = this.s[:len(this.s)-1]
		delete(this.m, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	fmt.Println(this.s)
	return this.s[rand.Intn(len(this.s))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

// func main() {
// 	rs := Constructor()

// 	fmt.Println(rs.Insert(1))
// 	fmt.Println(rs.Insert(0))
// 	fmt.Println(rs.Remove(0))
// 	fmt.Println(rs.Insert(2))
// 	fmt.Println(rs.Remove(1))
// 	fmt.Println(rs.GetRandom())

// }
