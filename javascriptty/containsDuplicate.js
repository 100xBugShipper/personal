/**
    @param {number []}
    @return {boolean}
*/

function hasDuplicate(arr) {
    let frequencyCounter = {}
    for (let i = 0; i <= arr.length - 1; i++) {
        if (frequencyCounter[arr[i]]) {
            frequencyCounter[arr[i]] += 1
        }else {
            frequencyCounter[arr[i]] = 1
        }
    }
    console.log(frequencyCounter)
    for (const val of Object.values(frequencyCounter)) {
        if (val > 1) {
            return true
        }
    }
    return false
}

console.log(hasDuplicate([1,2,3,3]))
