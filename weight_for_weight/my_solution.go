package weight_for_weight

import (
	"sort"
	"strings"
)

type elem []string

func (e elem) Len() int      { return len(e) }
func (e elem) Swap(i, j int) { e[i], e[j] = e[j], e[i] }
func (e elem) Less(i, j int) bool {
	a, b := e.digitSum(i), e.digitSum(j)
	if a == b {
		return e[i] < e[j]
	} else {
		return a < b
	}
}

func (e elem) digitSum(ind int) int {
	var temp int
	for _, val := range e[ind] {
		temp += int(val - '0')
	}
	return temp
}

func OrderWeight(str string) string {
	elems := strings.Fields(str)
	sort.Sort(elem(elems))
	return strings.Join(elems, " ")
}
