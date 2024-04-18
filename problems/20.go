package problems

func isValid(s string) bool {
	var stack []byte
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			var p = stack[len(stack)-1]
			if (s[i] == ')' && p != '(') ||
				(s[i] == '}' && p != '{') ||
				(s[i] == ']' && p != '[') {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
