// 题目：最大间距
// 给定一个未经排序的数组，请找出其排序表中连续两个要素的最大间距。
// 如果数组中的要素少于 2 个，请返回 0.
// 注意事项：
// 可以假定数组中的所有要素都是非负整数，且最大不超过 32 位整数。
// 样例：
// 给定数组 [1, 9, 2, 5]，其排序表为 [1, 2, 5, 9]，其最大的
// 间距是在 5 和 9 之间，= 4.
package p_00176_maximum_gap

// 解法1
//
// 排序再遍历
//
// 复杂度：排序nlogn + 遍历n = nlogn



// 解法2
//
// 对这些数平均分桶，如果是闭区间，桶数=数组个数
// 桶内左闭右开，最后一个数会单独一桶，桶数=数组个数+1
// 数组个数比桶数大会有空桶，保证maxgap在桶间
//
// 复杂度：遍历拆桶n + 遍历桶距n = n
