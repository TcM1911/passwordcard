package utils

// Shufle is an implementation of the java Collections.shuffle() method.
func Shufle(list []byte, rnd Randomer) {
	size := len(list)
	for i := size; i > 1; i-- {
		swap(list, i-1, rnd.nextInt(i))
	}
}

func swap(list []byte, i, j int) {
	tmp := list[i]
	list[i] = list[j]
	list[j] = tmp
}
