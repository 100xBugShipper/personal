/**
    @param {number []}: void
*/
function concatArray(arr) {
    const n = arr.length // imp as arr will dynamically grow so better keep the length const
    for (let i = 0; i <= n - 1; i++) {
        arr.push(arr[i])
    }
    return arr
}

console.log(concatArray([1,2,3,4,5]))
