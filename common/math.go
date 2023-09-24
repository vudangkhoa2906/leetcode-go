package common

func Abs[V ~int | ~int64](i V) V {
	if i > 0 {
		return i
	}
	return -i
}

func Max[V ~int | ~int64](i1, i2 V) V {
	if i1 > i2 {
		return i1
	}
	return i2
}

func Min[V ~int | ~int64](i1, i2 V) V {
	if i1 < i2 {
		return i1
	}
	return i2
}

func Square[V ~int | ~int64](i V) V {
	return i * i
}

func Compute[V ~int | ~int64](i1, i2 V, operator byte) V {
	switch operator {
	case '+':
		return i1 + i2
	case '-':
		return i1 - i2
	case '*':
		return i1 * i2
	default:
		return 0
	}
}

func Pow[V ~int | ~int64](i1, i2 V) V {
	res := V(1)
	for idx := V(1); idx <= i2; idx++ {
		res *= i1
	}
	return res
}
