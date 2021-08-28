func maxArea(height []int) int {
    i := 0
    j := len(height) - 1
    
    ret := 0
    for i < j {
        ret = max(ret, (j - i) * min(height[i], height[j]))
        if height[i] < height[j] {
            i++
        } else {
            j--
        }
    }
    
    return ret
}

func min(a int, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a int, b int) int {
    if a < b {
        return b
    }
    return a
}