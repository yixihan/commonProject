/**
<p>给你一个整数数组 <code>arr</code>，只有可以将其划分为三个和相等的 <strong>非空</strong> 部分时才返回&nbsp;<code>true</code>，否则返回 <code>false</code>。</p>

<p>形式上，如果可以找出索引&nbsp;<code>i + 1 &lt; j</code>&nbsp;且满足&nbsp;<code>(arr[0] + arr[1] + ... + arr[i] == arr[i + 1] + arr[i + 2] + ... + arr[j - 1] == arr[j] + arr[j + 1] + ... + arr[arr.length - 1])</code>&nbsp;就可以将数组三等分。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>arr = [0,2,1,-6,6,-7,9,1,2,0,1]
<strong>输出：</strong>true
<strong>解释：</strong>0 + 2 + 1 = -6 + 6 - 7 + 9 + 1 = 2 + 0 + 1
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>arr = [0,2,1,-6,6,7,9,-1,2,0,1]
<strong>输出：</strong>false
</pre>

<p><strong>示例 3：</strong></p>

<pre>
<strong>输入：</strong>arr = [3,3,6,5,-2,2,5,1,-9,4]
<strong>输出：</strong>true
<strong>解释：</strong>3 + 3 = 6 = 5 - 2 + 2 + 5 + 1 - 9 + 4
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
 <li><code>3 &lt;= arr.length &lt;= 5 * 10<sup>4</sup></code></li>
 <li><code>-10<sup>4</sup> &lt;= arr[i] &lt;= 10<sup>4</sup></code></li>
</ul>

<div><div>Related Topics</div><div><li>贪心</li><li>数组</li></div></div><br><div><li>👍 182</li><li>👎 0</li></div>
*/

//leetcode submit region begin(Prohibit modification and deletion)
package main

func canThreePartsEqualSum(arr []int) bool {
	sum := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	if sum%3 != 0 {
		return false
	}

	l, r := 0, len(arr)-1
	lSum, rSum := arr[l], arr[r]

	avg := sum / 3
	for l+1 < r {
		if lSum == avg && rSum == avg {
			return true
		}
		if lSum != avg {
			l++
			lSum += arr[l]
		}

		if rSum != avg {
			r--
			rSum += arr[r]
		}


	}

	return false
}

//leetcode submit region end(Prohibit modification and deletion)
