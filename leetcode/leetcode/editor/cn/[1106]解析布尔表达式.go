package main
//给你一个以字符串形式表述的 布尔表达式（boolean） expression，返回该式的运算结果。 
//
// 有效的表达式需遵循以下约定： 
//
// 
// "t"，运算结果为 True 
// "f"，运算结果为 False 
// "!(expr)"，运算过程为对内部表达式 expr 进行逻辑 非的运算（NOT） 
// "&(expr1,expr2,...)"，运算过程为对 2 个或以上内部表达式 expr1, expr2, ... 进行逻辑 与的运算（AND） 
// "|(expr1,expr2,...)"，运算过程为对 2 个或以上内部表达式 expr1, expr2, ... 进行逻辑 或的运算（OR） 
// 
//
// 
//
// 示例 1： 
//
// 输入：expression = "!(f)"
//输出：true
// 
//
// 示例 2： 
//
// 输入：expression = "|(f,t)"
//输出：true
// 
//
// 示例 3： 
//
// 输入：expression = "&(t,f)"
//输出：false
// 
//
// 示例 4： 
//
// 输入：expression = "|(&(t,f,t),!(t))"
//输出：false
// 
//
// 
//
// 提示： 
//
// 
// 1 <= expression.length <= 20000 
// expression[i] 由 {'(', ')', '&', '|', '!', 't', 'f', ','} 中的字符组成。 
// expression 是以上述形式给出的有效表达式，表示一个布尔值。 
// 
//
// Related Topics 栈 递归 字符串 👍 164 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func parseBoolExpr(expression string) bool {
	stk := []rune{}
	for _, c := range expression {
		if c != '(' && c != ')' && c != ',' {
			stk = append(stk, c)
		} else if c == ')' {
			var t, f int
			for stk[len(stk)-1] == 't' || stk[len(stk)-1] == 'f' {
				if stk[len(stk)-1] == 't' {
					t++
				} else {
					f++
				}
				stk = stk[:len(stk)-1]
			}
			op := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			c = 'f'
			if (op == '!' && f > 0) || (op == '&' && f == 0) || (op == '|' && t > 0) {
				c = 't'
			}
			stk = append(stk, c)
		}
	}
	return stk[0] == 't'
}
//leetcode submit region end(Prohibit modification and deletion)
