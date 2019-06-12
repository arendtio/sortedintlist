package sortedIntList

type Interface interface {
	Insert(int)
	Remove(int)
	Length() int
	At(int) int
}

// hiding implementation details
type List struct {
	list []int
}

func NewSortedIntList() *List {
	l := &List{make([]int, 0, 128)}
	return l
}

// TODO: build a linked list version and benchmark it against this one

func (l *List) Insert(x int) {
	// find position
	to := len(l.list) - 1
	pos := find(x, l.list, 0, to) // find can handle a negative 'to'

	l.list = append(l.list, 0)         // increase size
	copy(l.list[pos+1:], l.list[pos:]) // shift
	l.list[pos] = x
}

func (l *List) Remove(x int) {
	pos := find(x, l.list, 0, len(l.list)-1)
	pos--
	if pos < 0 {
		pos = 0
	}
	l.list = append(l.list[:pos], l.list[pos+1:]...)
}

func (l List) Length() int {
	return len(l.list)
}

func (l List) At(pos int) int {
	return l.list[pos]
}

// recursive function to find the correct spot to insert the new number
// invariable: low <= x < high
func find(x int, list []int, low int, high int) int {
	if len(list) == 0 { // not much to do if the list is empty
		return 0
	} else if low == high {
		if list[low] < x {
			return low + 1
		}
		return low
	} else if list[high] <= x {
		return high + 1
	} else if list[low] > x {
		return low
	} else if low+1 == high {
		if list[low] >= x {
			return high
		} else if list[high] >= x {
			return high
		} else {
			return high + 1
		}
	} else {
		middle := low + (high-low)/2
		newLow := low
		newHigh := high
		if list[middle] < x {
			newLow = middle
		} else {
			newHigh = middle
		}
		return find(x, list, newLow, newHigh)
	}
}
