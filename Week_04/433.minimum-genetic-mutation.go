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
    visited := make(map[string]struct{})
    step, length := 0, len(start)
    for len(queue) > 0 {
        n := len(queue)
        for i := 0; i < n; i++ {
            start = queue[i]
            if start == end { return step }
            for j := 0; j < length; j++ {
                for _, c := range changeMap[start[j]] {
                    arr := []byte(start)
                    arr[j] = c
                    tmp := string(arr)
                    if _, ok := bankMap[tmp]; ok {
                        if _, ok = visited[tmp]; ok {
                            continue
                        }
                        visited[tmp] = struct{}{}
                        queue = append(queue, tmp)
                    }
                }
            }
        }
        queue = queue[n:]
        step++
    }
    return -1
}
