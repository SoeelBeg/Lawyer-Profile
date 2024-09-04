document.addEventListener('DOMContentLoaded', function () {
    const contentDiv = document.getElementById('content');
    const getLawyerResultDiv = document.getElementById('get-lawyer-result');

    // Add Lawyer
    if (document.getElementById('create-lawyer-form')) {
        document.getElementById('create-lawyer-form').addEventListener('submit', function (e) {
            e.preventDefault();

            const lawyer = {
                name: document.getElementById('create-name').value,
                specialty: document.getElementById('create-specialty').value,
                email: document.getElementById('create-email').value,
                mobile_number: document.getElementById('create-mobile').value,
                address: document.getElementById('create-address').value
            };

            fetch('/api/lawyers', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(lawyer)
            })
            .then(response => response.json())
            .then(data => {
                alert('Lawyer added successfully!');
                window.location.href = '/index.html';  // Redirect to homepage
            })
            .catch(error => {
                console.error('Error adding lawyer:', error);
            });
        });
    }

    // Get Lawyer by ID
    if (document.getElementById('get-lawyer-id')) {
        document.getElementById('get-lawyer-btn').addEventListener('click', function () {
            const lawyerId = document.getElementById('get-lawyer-id').value;
            if (!lawyerId) {
                alert('Please enter a lawyer ID');
                return;
            }

            fetch(`/api/lawyers/${lawyerId}`)
                .then(response => response.json())
                .then(lawyer => {
                    if (lawyer) {
                        getLawyerResultDiv.innerHTML = `
                            <div>
                                <strong>${lawyer.name}</strong><br>
                                Specialty: ${lawyer.specialty}<br>
                                Email: ${lawyer.email}<br>
                                Mobile: ${lawyer.mobile_number}<br>
                                Address: ${lawyer.address}
                            </div>
                        `;
                    } else {
                        getLawyerResultDiv.innerHTML = "<p>No lawyer found with this ID.</p>";
                    }
                })
                .catch(error => {
                    console.error('Error fetching lawyer:', error);
                });
        });
    }

    // Update Lawyer
    if (document.getElementById('update-lawyer-form')) {
        document.getElementById('update-lawyer-form').addEventListener('submit', function (e) {
            e.preventDefault();

            const lawyerId = document.getElementById('update-lawyer-id').value;
            if (!lawyerId) {
                alert('Please enter a lawyer ID');
                return;
            }

            const updatedLawyer = {
                name: document.getElementById('update-name').value,
                specialty: document.getElementById('update-specialty').value,
                email: document.getElementById('update-email').value,
                mobile_number: document.getElementById('update-mobile').value,
                address: document.getElementById('update-address').value
            };

            fetch(`/api/lawyers/${lawyerId}`, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(updatedLawyer)
            })
            .then(response => response.json())
            .then(data => {
                alert('Lawyer updated successfully!');
                window.location.href = '/index.html';  // Redirect to homepage
            })
            .catch(error => {
                console.error('Error updating lawyer:', error);
            });
        });
    }

    // Delete Lawyer
    if (document.getElementById('delete-lawyer-form')) {
        document.getElementById('delete-lawyer-form').addEventListener('submit', function (e) {
            e.preventDefault();

            const lawyerId = document.getElementById('delete-lawyer-id').value;
            if (!lawyerId) {
                alert('Please enter a lawyer ID');
                return;
            }

            fetch(`/api/lawyers/${lawyerId}`, {
                method: 'DELETE'
            })
            .then(response => {
                if (response.ok) {
                    alert('Lawyer deleted successfully!');
                    window.location.href = '/index.html';  // Redirect to homepage
                } else {
                    alert('Failed to delete lawyer');
                }
            })
            .catch(error => {
                console.error('Error deleting lawyer:', error);
            });
        });
    }

    // List All Lawyers
    if (document.getElementById('lawyer-list')) {
        fetch('/api/lawyers')
            .then(response => response.json())
            .then(lawyers => {
                const lawyerListDiv = document.getElementById('lawyer-list');

                if (lawyers.length === 0) {
                    lawyerListDiv.innerHTML = "<p>No lawyers found.</p>";
                } else {
                    lawyers.forEach(lawyer => {
                        lawyerListDiv.innerHTML += `
                            <div class="lawyer-item">
                                <h2>${lawyer.name}</h2>
                                <p>ID: ${lawyer.id}</p>
                                <p>Specialty: ${lawyer.specialty}</p>
                                <p>Email: ${lawyer.email}</p>
                                <p>Mobile: ${lawyer.mobile_number}</p>
                                <p>Address: ${lawyer.address}</p>
                            </div>
                            <hr>
                        `;
                    });
                }
            })
            .catch(error => {
                console.error('Error fetching lawyers:', error);
            });
    }
});
