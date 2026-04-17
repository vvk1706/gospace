// Contact form-specific JavaScript

document.addEventListener('DOMContentLoaded', function() {
    const form = document.getElementById('contactForm');
    const nameInput = document.getElementById('name');
    const surnameInput = document.getElementById('surname');
    const emailInput = document.getElementById('email');

    // Real-time email validation
    emailInput.addEventListener('blur', function() {
        const email = this.value.trim();
        if (email && !isValidEmail(email)) {
            this.style.borderColor = '#c33';
            showValidationMessage(this, 'Please enter a valid email address');
        } else {
            this.style.borderColor = '#e0e0e0';
            removeValidationMessage(this);
        }
    });

    // Capitalize first letter of name and surname
    [nameInput, surnameInput].forEach(input => {
        input.addEventListener('blur', function() {
            if (this.value) {
                this.value = this.value.charAt(0).toUpperCase() + this.value.slice(1).toLowerCase();
            }
        });
    });

    // Form submission handler
    form.addEventListener('submit', function(e) {
        // Client-side validation
        const name = nameInput.value.trim();
        const surname = surnameInput.value.trim();
        const email = emailInput.value.trim();

        if (!name || !surname || !email) {
            e.preventDefault();
            alert('Please fill in all fields');
            return;
        }

        if (!isValidEmail(email)) {
            e.preventDefault();
            alert('Please enter a valid email address');
            emailInput.focus();
            return;
        }

        // Add loading state
        const submitBtn = form.querySelector('button[type="submit"]');
        submitBtn.textContent = 'Submitting...';
        submitBtn.disabled = true;
    });

    // Reset form after successful submission
    const successAlert = document.querySelector('.alert-success');
    if (successAlert) {
        setTimeout(() => {
            form.reset();
            nameInput.focus();
        }, 2000);
    }
});

// Helper function to show validation message
function showValidationMessage(input, message) {
    removeValidationMessage(input);
    const errorDiv = document.createElement('div');
    errorDiv.className = 'validation-error';
    errorDiv.style.cssText = 'color: #c33; font-size: 0.875rem; margin-top: 0.25rem;';
    errorDiv.textContent = message;
    input.parentElement.appendChild(errorDiv);
}

// Helper function to remove validation message
function removeValidationMessage(input) {
    const existingError = input.parentElement.querySelector('.validation-error');
    if (existingError) {
        existingError.remove();
    }
}

// Email validation function (imported from main.js but defined here for standalone use)
function isValidEmail(email) {
    const re = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return re.test(email);
}

// Add character counter for inputs
document.addEventListener('DOMContentLoaded', function() {
    const inputs = document.querySelectorAll('#contactForm input[type="text"], #contactForm input[type="email"]');
    inputs.forEach(input => {
        input.addEventListener('input', function() {
            const maxLength = 100;
            if (this.value.length > maxLength) {
                this.value = this.value.substring(0, maxLength);
            }
        });
    });
});

// Made with Bob
