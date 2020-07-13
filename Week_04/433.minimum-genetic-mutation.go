func minMutation(start string, end string, bank []string) int {
    bankMap := make(map[string]struct{}, len(bank))
    for i := range bank { bankMap[bank[i]] = struct{}{} }

    if _, ok := bankMap[end]; !ok { return -1 }

    changeMap := map[byte][]byte{
        'A': []byte{'C', 'G', 'T'},
        'C': []byte{'A', 'G', 'T'},
        'G': []byte{'A', 'C', 'T'},
        'T': []byte{'A', 'C', 'G'},
    }

    queue := []string{ start }
    step := 0
    for len(queue) > 0 {
        n := len(queue)
        for i := 0; i < n; i++ {
            start = queue[0]
            queue = queue[1:]
            if start == end { return step }
            // 这里可以优化，只变更需要变更的字符;(最开始循环一遍字符串就知道哪些字符不一样，需要多少步)
            for j := 0; j < len(start); j++ {
                for _, c := range changeMap[start[j]] {
                    arr := []byte(start)
                    arr[j] = c
                    tmp := string(arr)
                    if _, ok := bankMap[tmp]; ok {
                        queue = append(queue, tmp)
                    }
                }
            }
        }
        step++
    }
    return -1
}
