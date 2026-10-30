func groupAnagrams(strs []string) [][]string {
    if len(strs) == 0 {
        return nil
    }

    
    groups := make(map[[26]int][]string)

    for _, str := range strs {
        var count [26]int
        for _, char := range str {
            count[char-'a']++
        }

        groups[count] = append(groups[count], str)
    }

    var sublists [][]string
    for _, group := range groups {
        sublists = append(sublists, group)
    }

    return sublists
}
