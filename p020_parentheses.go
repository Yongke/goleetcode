package goleetcode

func isValid(s string) bool {
	m := map[byte]byte{
		')': '(', ']': '[', '}': '{',
	}
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		}
		if s[i] == ')' || s[i] == ']' || s[i] == '}' {
			if len(stack) == 0 || m[s[i]] != stack[len(stack)-1] {
				return false
			}
			stack = stack[0 : len(stack)-1]
		}
	}
	return len(stack) == 0
}
