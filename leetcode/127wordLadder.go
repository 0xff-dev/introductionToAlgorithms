package leetcode

func ladderLength(beginWord string, endWord string, wordList []string) int {
    distance := 1
    // 检查单词是否访问过了。
    visited := make(map[string]struct{})
    words := make(map[string]struct{})
    for _, w := range wordList {
        words[w] = struct{}{}
    }
    if _, ok := words[endWord]; !ok {
        return 0
    }
    queue := []string{beginWord}

    for len(queue) > 0 {
        nextQueue := make(map[string]struct{})
        for _, o := range queue {
            visited[o] = struct{}{}
        }
        for _, o := range queue {
            bs := []byte(o)
            for idx := 0; idx < len(bs); idx++ {
                for char:= 'a'; char <= 'z'; char++ {
                    if byte(char) == bs[idx] {
                        continue
                    }
                    bs[idx] = byte(char)
                    if string(bs) == endWord {
                        return distance+1
                    }
                    newWord := string(bs)
                    if _, ok := words[newWord]; ok {
                        if _, ok = visited[newWord]; !ok {
                            nextQueue[newWord] = struct{}{}
                        }
                    }
                }
                bs = []byte(o)
            }
        }
        distance++
        queue = []string{}
        for k := range nextQueue {
            queue = append(queue, k)
        }
    }
    return 0
}
