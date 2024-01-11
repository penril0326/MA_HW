package parenthesisremoval

import (
	"strings"
)

func ParenthesisRemoval(expr string) string {
	last, next := make([]int, len(expr)), make([]int, len(expr))

	keep := make([]bool, len(expr))
	for i := 0; i < len(expr); i++ {
		keep[i] = true
	}

	// iterate from head, store last operator
	op := -1
	for i := 0; i < len(expr); i++ {
		last[i] = op
		if expr[i] == '+' || expr[i] == '-' || expr[i] == '*' || expr[i] == '/' {
			op = int(expr[i])
		}
	}

	// iterate from tail, store next operator
	op = -1
	for i := len(expr) - 1; i >= 0; i-- {
		next[i] = op
		if expr[i] == '+' || expr[i] == '-' || expr[i] == '*' || expr[i] == '/' {
			op = int(expr[i])
		}
	}

	st := make([]int, 0)
	opIdx := make(map[byte]int)
	ops := []byte{'+', '-', '*', '/'}

	for i := 0; i < len(expr); i++ {
		for _, op := range ops {
			if op == expr[i] {
				opIdx[op] = i
				break
			}
		}

		if expr[i] == '(' {
			st = append(st, i)
		} else if expr[i] == ')' {
			// pop
			leftBracketIdx := st[len(st)-1]
			st = st[:len(st)-1]

			rightBracketIdx := i

			nxt := next[rightBracketIdx]
			lst := last[leftBracketIdx]

			// check which operator is in current brackets
			isInBrackets := make(map[byte]bool)
			for _, op := range ops {
				if idx, isExist := opIdx[op]; isExist && (idx >= leftBracketIdx) {
					isInBrackets[op] = true
				}
			}

			redundant := false

			// check if there is other duplicated parenthesis outside: ((a+b))
			if (leftBracketIdx > 0) && (rightBracketIdx+1 < len(expr)) &&
				((expr[leftBracketIdx-1] == '(') && (expr[rightBracketIdx+1] == ')') ||
					(expr[leftBracketIdx-1] == '(' && ((nxt == '+') || (nxt == '-')))) {
				redundant = true
			} else {
				// remove the op if it was be checked and there is no other outside duplicated parenthesis.
				for op, _ := range isInBrackets {
					if _, isExist := opIdx[op]; isExist {
						delete(opIdx, op)
					}
				}
			}

			// check the case that no any ops in the current brackets: (a)
			if !isInBrackets['+'] && !isInBrackets['*'] && !isInBrackets['-'] && !isInBrackets['/'] {
				redundant = true
			}

			if lst == '/' {
				// if last operator is /, need to keep current brackets
			} else if lst == '-' && (isInBrackets['+'] || isInBrackets['-']) {
				// if last operator is -, and current brackets contain - or +, then we need to keep brackets
			} else if !isInBrackets['-'] && !isInBrackets['+'] {
				redundant = true
			} else {
				if (lst == -1 || lst == '+' || lst == '-') && (nxt == -1 || nxt == '+' || nxt == '-') {
					redundant = true
				}
			}

			// update keep so that we know which parentheis can be removed
			if redundant {
				keep[leftBracketIdx] = false
				keep[rightBracketIdx] = false
			}
		}
	}

	// Final string
	var sb strings.Builder
	for i := 0; i < len(expr); i++ {
		if keep[i] {
			sb.WriteByte(expr[i])
		}
	}

	return sb.String()
}
