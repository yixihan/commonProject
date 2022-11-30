package main
//设计一个类似堆栈的数据结构，将元素推入堆栈，并从堆栈中弹出出现频率最高的元素。 
//
// 实现 FreqStack 类: 
//
// 
// 
// FreqStack() 构造一个空的堆栈。 
// 
// void push(int val) 将一个整数 val 压入栈顶。 
// 
// int pop() 删除并返回堆栈中出现频率最高的元素。 
// 
// 如果出现频率最高的元素不只一个，则移除并返回最接近栈顶的元素。 
// 
// 
//
// 
//
// 示例 1： 
//
// 
//输入：
//["FreqStack","push","push","push","push","push","push","pop","pop","pop",
//"pop"],
//[[],[5],[7],[5],[7],[4],[5],[],[],[],[]]
//输出：[null,null,null,null,null,null,null,5,7,5,4]
//解释：
//FreqStack = new FreqStack();
//freqStack.push (5);//堆栈为 [5]
//freqStack.push (7);//堆栈是 [5,7]
//freqStack.push (5);//堆栈是 [5,7,5]
//freqStack.push (7);//堆栈是 [5,7,5,7]
//freqStack.push (4);//堆栈是 [5,7,5,7,4]
//freqStack.push (5);//堆栈是 [5,7,5,7,4,5]
//freqStack.pop ();//返回 5 ，因为 5 出现频率最高。堆栈变成 [5,7,5,7,4]。
//freqStack.pop ();//返回 7 ，因为 5 和 7 出现频率最高，但7最接近顶部。堆栈变成 [5,7,5,4]。
//freqStack.pop ();//返回 5 ，因为 5 出现频率最高。堆栈变成 [5,7,4]。
//freqStack.pop ();//返回 4 ，因为 4, 5 和 7 出现频率最高，但 4 是最接近顶部的。堆栈变成 [5,7]。 
//
// 
//
// 提示： 
//
// 
// 0 <= val <= 10⁹ 
// push 和 pop 的操作数不大于 2 * 10⁴。 
// 输入保证在调用 pop 之前堆栈中至少有一个元素。 
// 
//
// Related Topics 栈 设计 哈希表 有序集合 👍 271 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
type FreqStack struct {
	cnt    map[int]int
	stacks [][]int
}

func Constructor() FreqStack {
	return FreqStack{cnt: map[int]int{}}
}

func (f *FreqStack) Push(val int) {
	c := f.cnt[val]
	if c == len(f.stacks) { // 这个元素的频率已经是目前最多的，现在又出现了一次
		f.stacks = append(f.stacks, []int{val}) // 那么必须创建一个新栈
	} else {
		f.stacks[c] = append(f.stacks[c], val) // 否则就压入对应的栈
	}
	f.cnt[val]++ // 更新频率
}

func (f *FreqStack) Pop() int {
	back := len(f.stacks) - 1
	st := f.stacks[back]
	bk := len(st) - 1
	val := st[bk] // 弹出最右侧栈的栈顶
	if bk == 0 {  // 栈为空
		f.stacks = f.stacks[:back] // 删除
	} else {
		f.stacks[back] = st[:bk]
	}
	f.cnt[val]-- // 更新频率
	return val
}


/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */
//leetcode submit region end(Prohibit modification and deletion)
