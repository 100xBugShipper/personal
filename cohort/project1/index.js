let val1 = document.getElementById("textArea1")
let val2 = document.getElementById("textArea2")


const buttonElement = document.getElementById("sumButton")
buttonElement.addEventListener("click", function() {
    let sum = Number(val1.value) + Number(val2.value)
    alert(`Sum is ${sum}`)
})

