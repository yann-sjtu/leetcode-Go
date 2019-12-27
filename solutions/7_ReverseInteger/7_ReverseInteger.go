package __ReverseInteger

func reverse(x int) int {
	if x < 10 && x > -10 {
		return x
	}
	var flag bool
	var int32X int32
	if x < 0 {
		flag = true
		int32X = -int32(x)
	} else {
		int32X = int32(x)
	}

	var reverseX, remainder, temp int32
	for ;int32X!=0;int32X/=10{
		remainder = int32X%10
		temp = 10*reverseX + remainder
		if reverseX != (temp - remainder)/10 {
			//overflow
			return 0
		}
		reverseX = temp
	}

	if flag {
		return -int(reverseX)
	}
	return int(reverseX)
}
