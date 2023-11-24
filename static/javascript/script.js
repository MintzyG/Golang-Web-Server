document.addEventListener('DOMContentLoaded', function() {
  // Add click event listener to all elements with class 'deleteButton'
  const deleteButtons = document.querySelectorAll('.deleteButton');
  deleteButtons.forEach(button => {
    button.addEventListener('click', function() {
      const userId = button.getAttribute('data-user-id');
      
      // Send DELETE request
      fetch(`/dashboard/${userId}`, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
          // Add any additional headers as needed
        },
        // You can add a body if necessary
        // body: JSON.stringify({ key: 'value' }),
      })
      .then(response => {
        if (response.ok) {
          // Handle success (e.g., remove the list item from the DOM)
          button.closest('li').remove();
        } else {
          // Handle error
          console.error('Failed to delete user');
        }
      })
      .catch(error => {
        console.error('Error:', error);
      });
    });
  });
});
