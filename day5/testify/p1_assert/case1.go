package p1_assert

func getValue(t int) any {
	name := "Bob"
	age := 10
	if t == 1 {
		return name
	} else {
		return age
	}
}
