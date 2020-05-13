package valid_parentheses

//depend on stack.go
func Valid(s string) bool {
	//
	maps := map[uint8]uint8{'(':')','{':'}','[':']'}

	stack := NewStack()

	for i:=0; i < len(s); i++ {
		if stack.len() > 0 {
			val, exist := maps[stack.peak().(uint8)]

			if exist && val == s[i] { //if stack top in ['(','[','{'] && current byte equal it's map value,  pop the top
				stack.pop()
			} else {
				stack.push(s[i])
			}
		} else {
			stack.push(s[i])
		}
	}

	return stack.len() == 0
}

//view the discussions, found another solution, by slice ... so smart
func Valid2(s string) bool {
	maps := map[rune]rune{'(':')','{':'}','[':']'}

	stack := make([]rune, 0)

	for _,v := range s {
		if len(stack) == 0 {
			stack = append(stack, v)
			continue
		}

		val, exist := maps[stack[len(stack)-1]]

		if exist && val == v { //if stack top in ['(','[','{'] && current byte equal it's map value,  pop the top
			stack = stack[0: len(stack)-1]
		} else {
			stack = append(stack,v)
		}
	}

	return len(stack) == 0
}
