package list

import "container/list"

type HighList struct {
	list.List
	priority int
}

type ShortList struct {
	list.List
	priority    int
	needRunTime int
}

type LongList struct {
	list.List
	priority    int
	needRunTime int
}
