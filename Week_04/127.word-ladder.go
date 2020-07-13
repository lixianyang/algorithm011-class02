func ladderLength(beginWord string, endWord string, wordList []string) int {
    words := make(map[string]int, len(wordList))
    for i, word := range wordList { words[word] = i }
    visited := make([]bool, len(wordList))  // 下标
    if _, ok := words[endWord]; !ok {
        return 0
    }

    step, length := 1, len(beginWord)
    queue := []string{ beginWord }
    for len(queue) > 0 {
        size := len(queue)
        for i := 0; i < size; i++ {
            beginWord = queue[i]
            if beginWord == endWord { return step }
            chars := []byte(beginWord)
            for j := 0; j < length; j++ {
                o := chars[j]
                for c := 'a'; c <= 'z'; c++ {
                    chars[j] = byte(c)
                    tmp := string(chars)
                    if index, ok := words[tmp]; ok {
                        if visited[index] {
                            continue
                        }
                        visited[index] = true
                        queue = append(queue, tmp)
                    } else {
                        continue
                    }
                }
                chars[j] = o
            }
        }
        queue = queue[size:]
        step++
    }

    return 0
}
