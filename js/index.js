const clickableItems = document.querySelectorAll('.citem')

// Loop through each item and add a click event listener
clickableItems.forEach((item, index) => {
    item.addEventListener('click', () => {
        // Redirect the user to a new path
        window.location.href = `/group/${index + 1}`
    })
})