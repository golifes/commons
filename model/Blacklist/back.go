package Blacklist

/**
白名单
*/

type Blacklist struct {
	Id   int
	Ip   string
	Name string //谁的名称
	Type int    //是谁的标记id
}
