package solutions

func convert(s string, numRows int) string {
	l := len(s)
	if numRows == 1 || len(s) <= numRows {
		return s
	}

	byteS := make([]byte, 0, l)
	var left, right int
	//遍历每一行
	for n:=0;n<numRows;n++ {
		//以每第n+2*numRows-2个为中心，依次遍历。
		//行数越小的，距离中心越近，从左往右依次加入新的字符串中
		//边界处理：最后多一个虚拟中心
		for i:=0;i<l+2*numRows-2;i+=2*numRows-2 {
			if left=i-n;left>=0 &&left<l{
				byteS = append(byteS, s[left])
			}
			if right = i+n;right<l && n !=0 && n != numRows -1{
				byteS = append(byteS, s[right])
			}
		}
	}
	return string(byteS)
}