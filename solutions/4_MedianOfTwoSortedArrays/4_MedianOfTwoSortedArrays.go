package __MedianOfTwoSortedArrays

/*
解题思路：
求两个有序数组的中位数，是求两个有序数组第k小的数的一个特例。

一般思路是从数组A和数组B中各取一个最小的数，把更小的删除，然后再从该数组中取一个数，删除k-1个数之后剩下的最小的数就是第k小。
改进后的思路是批量取、删。每次从两个数组中取多个数，不断剔除最小的一批
不妨假设m中的最大值小于n中的最大值，那么说明数组A中应该删除数字的个数大于m而数组B中应该删除的数字的个数小于n，因此先把数组A中m个数字全部删除，对数组B的n个数进行二分。

求中位数：
分奇偶，奇数则取中间一个，偶数则取中间两个求均值
*/

//function [findMedianSortedArrays] is a special case of function [findKthElem]
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1)+len(nums2)
	if totalLength&1 == 1 {
		return float64(findKthElem(nums1, nums2, totalLength/2+1))
	}
	return 0.5 * float64(findKthElem(nums1, nums2, totalLength/2) + findKthElem(nums1, nums2, totalLength/2+1))

}

//1<=k<=len(nums1)+len(nums2)
func findKthElem(nums1, nums2 []int, k int) int {
	var i, j, l1 int
	for {
		if len(nums1) > len(nums2) {
			nums1, nums2 = nums2, nums1
		}
		l1 = len(nums1)
		if l1 == 0 {
			return nums2[k-1]
		}
		if k == 1 {
			return min(nums1[0], nums2[0])
		}
		if k/2 > l1 {
			i = l1
		} else {
			i = k / 2
		}
		j = k - i
		// i,j <= k/2 <= l1 <= l2
		// k至少为2，且nums1和nums2都不为空，因此i，j都大于0
		if nums1[i-1] > nums2[j-1] {
			nums2 = nums2[j:]
			k -= j
		} else if nums1[i-1] < nums2[j-1] {
			nums1 = nums1[i:]
			k -= i
		} else {
			return nums1[i-1]
		}
	}
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}