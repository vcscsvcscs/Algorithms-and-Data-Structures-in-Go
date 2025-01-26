package trees

type AVLNode struct {
	Key         int
	Value       any
	Balance     int
	Left, Right *AVLNode
}

func NewAVLNode(key int, value any) *AVLNode {
	return &AVLNode{Key: key, Value: value, Balance: 0}
}

func AVLInsert(t **AVLNode, key int, value any, d *bool) {
	if *t == nil {
		*t = NewAVLNode(key, value)
	} else {
		switch {
		case key < (*t).Key:
			AVLInsert(&(*t).Left, key, value, d)
			if *d {
				leftSubTreeGrown(t, d)
			}
		case key > (*t).Key:
			AVLInsert(&(*t).Right, key, value, d)
			if *d {
				rightSubTreeGrown(t, d)
			}
		default:
			*d = false
		}
	}
}

func leftSubTreeGrown(t **AVLNode, d *bool) {
	if (*t).Balance == -1 {
		l := (*t).Left
		if l.Balance == -1 {
			balanceMMm(t, l)
		} else {
			balanceMMp(t, l)
		}
		*d = false
	} else {
		(*t).Balance--
		*d = (*t).Balance < 0
	}
}

func rightSubTreeGrown(t **AVLNode, d *bool) {
	if (*t).Balance == 1 {
		r := (*t).Right
		if r.Balance == 1 {
			balancePPp(t, r)
		} else {
			balancePPm(t, r)
		}
		*d = false
	} else {
		(*t).Balance++
		*d = (*t).Balance > 0
	}
}

func balanceMMm(t **AVLNode, l *AVLNode) {}

func balanceMMp(t **AVLNode, l *AVLNode) {}

func balancePPm(t **AVLNode, r *AVLNode) {}

func balancePPp(t **AVLNode, r *AVLNode) {
	(*t).Right = r.Left
	r.Left = *t
	(*t).Balance, r.Balance = 0, 0
	*t = r
}
