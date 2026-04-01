// инкремент словаря первым словом, декремент вторым
func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    // создаем словарь
    var dict = make(map[rune]int, len(s))
    // проходим по первому слову, инкрементируем
    for _, ri := range s {
        if _, ok := dict[ri]; ok {
            dict[ri] += 1
        } else {
            dict[ri] = 1
        }
    }

    // проходя по второму слову декрементируем
    for _, rd := range t {
        if count, ok := dict[rd]; ok {
            if count == 1 {
                // удаляем из мапы, если декрементировали в 0
                delete(dict, rd)
            } else {
                dict[rd] -= 1
            }
        }
    }

    // благодаря удалению, осталось проверить наличие остатка
    if len(dict) != 0 {
        return false
    }



    return true
}
