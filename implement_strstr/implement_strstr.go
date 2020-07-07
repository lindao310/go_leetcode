package implement_strstr

func strPos(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	for i := 0; i < len(haystack) - len(needle) + 1; i++ {
		//fmt.Printf("i=%d, substr=%s \n", i, haystack[i:i+len(needle)])
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}

	return -1
}
