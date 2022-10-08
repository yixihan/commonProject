package main

// getRow 最佳算法
func getRow(rowIndex int) []int {
	ans := make([]int, rowIndex+1)
	tmp := 1

	for i := 0; i <= rowIndex; i++ {
		ans[i] = tmp
		tmp = tmp * (rowIndex - i) / (i + 1)
	}

	return ans
}

/*
List<Integer> result = new ArrayList<>();
long temp = 1;

for (int i = 0; i <= rowIndex; i++) {
	result.add(Math.toIntExact(temp));
	temp = temp * (rowIndex - i) / (i + 1);
}

return result;
*/
