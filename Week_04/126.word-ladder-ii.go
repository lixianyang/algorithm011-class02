func findLadders(beginWord string, endWord string, wordList []string) [][]string {
    words := make(map[string]struct{}, len(wordList))
    for _, word := range wordList {
        words[word] = struct{}{}
    }
    if _, ok := words[endWord]; !ok {
        return nil
    }

    result := make([][]string, 0)
    queue := [][]string{ []string{beginWord} }  // 直接将路径存储在队列里
    visited := make(map[string]bool)

    length := len(beginWord)
    for len(queue) > 0 {
        levelVisited := make(map[string]bool)
        size := len(queue)
        for i := 0; i < size; i++ {
            path := queue[i]
            beginWord = path[len(path)-1]
            if beginWord == endWord {
                // 有结果，由于所有单词长度都相等，所以能找到的都是最短路径
                result = append(result, path)
                continue
            }
            chars := []byte(beginWord)
            for j := 0; j < length; j++ {
                o := beginWord[j]
                for c := 'a'; c <= 'z'; c++ {
                    chars[j] = byte(c)
                    tmp := string(chars)                        
                    if _, ok := words[tmp]; ok {
                        if visited[tmp] {
                            continue
                        }
                        levelVisited[tmp] = true
                        newPath := make([]string, len(path)+1)
                        copy(newPath, path)
                        newPath[len(newPath)-1] = tmp
                        queue = append(queue, newPath)
                    } else {
                        continue
                    }
                }
                chars[j] = o
            }
        }
        for k, v := range levelVisited {
            visited[k] = v
        }
        queue = queue[size:]
    }
    return result
}
