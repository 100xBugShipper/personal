/**
    @param {string} word1
    @param {string} word2
    @return {boolean}
*/

function isAnagram(word1, word2) {
    let seen = {}
    if (word1.length !== word2.length) {
        return false
    }
    for (let i = 0; i <= word1.length - 1; i++) {
        if (!seen[word1[i]]) {
            seen[word1[i]] = 1
        }else {
            seen[word1[i]]++
        }
    }

    for (let i = 0; i <= word2.length - 1; i++) {
        if (!seen[word2[i]]) {
            return false
        }else {
            seen[word2[i]]--
        }
    }

    for (const item of Object.values(seen)) {
        if (item > 0) {
            return false
        }
    }
    return true
}

console.log(isAnagram("listen", "silenty"))
