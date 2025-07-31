func ladderLength(beginWord string, endWord string, wordList []string) int {
    wordSet := make(map[string]bool)
    for _, word := range wordList {
        wordSet[word] = true
    }

    if !wordSet[endWord] {
        return 0
    }

    queue := []string{beginWord}
    visited := make(map[string]bool)
    visited[beginWord] = true
    steps := 1

    for len(queue) > 0 {
        levelSize := len(queue)
        for i := 0; i < levelSize; i++ {
            word := queue[0]
            queue = queue[1:]

            if word == endWord {
                return steps
            }

            // Try changing every letter
            for j := 0; j < len(word); j++ {
                for c := 'a'; c <= 'z'; c++ {
                    newWord := word[:j] + string(c) + word[j+1:]
                    if wordSet[newWord] && !visited[newWord] {
                        queue = append(queue, newWord)
                        visited[newWord] = true
                    }
                }
            }
        }
        steps++
    }

    return 0
}