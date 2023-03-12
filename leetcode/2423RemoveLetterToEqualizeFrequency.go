package leetcode

func equalFrequency(word string) bool {
    tmp := []byte(word)
    for i := 0; i < len(tmp); i++ {
        temp, m := make([]byte, len(tmp)), make([]int, 26)
        copy(temp, tmp)
        for _, c := range append(temp[:i], temp[i + 1:]...) { m[c - 'a']++ }
        c := make(map[int]int)
        for _, v := range m { if v != 0 { c[v]++ } }
        if len(c) == 1 { return true }
    }
    return false
}
