func minSpeedOnTime(dist []int, hour float64) int {
    i := 1
    j := int(1e7) + 1
    for i < j {
        m := i + (j - i) / 2
        time := 0.0
        
        for k := 0; k < len(dist) - 1; k++ {
            time += float64(dist[k] / m) * 1.0
            if dist[k] % m != 0 {
                time += 1.0
            }
        }
        
        time += float64(dist[len(dist) - 1]) / float64(m)
        
        if time > hour {
            i = m + 1
        } else {
            j = m
        }
    }
    
    if i == int(1e7) + 1 {
        return -1
    }
    return i
}