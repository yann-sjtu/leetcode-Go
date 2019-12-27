package __StringToInteger

func myAtoi(str string) int {
	byteS := trimSpace(str)
	l := len(byteS)
	if l == 0 {
		return 0
	}

	plus := byte(43)
	minus := byte(45)
	num0 := byte(48)
	num9 := byte(57)
	maxInt32 := int32(0x7fffffff)
	var symbol byte
	var myInt32, temp, num int32
	if byteS[0] == minus || byteS[0] == plus {
		symbol = byteS[0]
		byteS = byteS[1:]
	}
	for _, c := range byteS {
		if c < num0 || c > num9 {
			break
		}
		num = int32(c - num0)
		temp = myInt32*10 - num
		if myInt32 != (temp+num)/10 || temp > myInt32 {
			//overflow
			if symbol == minus {
				return int(maxInt32 + 1)
			}
			return int(maxInt32)
		}
		myInt32 = temp
	}
	if symbol != minus {
		if myInt32 == maxInt32+1 {
			return int(maxInt32)
		}
		return -int(myInt32)
	}
	return int(myInt32)
}

func trimSpace(s string) []byte {
	space := byte(32)
	byteS := []byte(s)
	for len(byteS) != 0 && byteS[0] == space {
		byteS = byteS[1:]
	}
	return byteS
}
