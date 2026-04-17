// Calculator-specific JavaScript

document.addEventListener('DOMContentLoaded', function() {
    const form = document.querySelector('.calculator-form');
    const num1Input = document.getElementById('num1');
    const num2Input = document.getElementById('num2');
    const operationSelect = document.getElementById('operation');

    // Add keyboard shortcuts
    document.addEventListener('keydown', function(e) {
        // Ctrl/Cmd + Enter to submit form
        if ((e.ctrlKey || e.metaKey) && e.key === 'Enter') {
            e.preventDefault();
            form.submit();
        }
    });

    // Add input validation
    [num1Input, num2Input].forEach(input => {
        input.addEventListener('input', function() {
            // Remove any non-numeric characters except decimal point and minus
            this.value = this.value.replace(/[^0-9.-]/g, '');
        });
    });

    // Add visual feedback for operation selection
    operationSelect.addEventListener('change', function() {
        const selectedOption = this.options[this.selectedIndex];
        this.style.color = '#667eea';
        setTimeout(() => {
            this.style.color = '';
        }, 300);
    });

    // Form submission handler
    form.addEventListener('submit', function(e) {
        const num1 = parseFloat(num1Input.value);
        const num2 = parseFloat(num2Input.value);
        const operation = operationSelect.value;

        // Client-side validation
        if (isNaN(num1) || isNaN(num2)) {
            e.preventDefault();
            alert('Please enter valid numbers');
            return;
        }

        if (operation === 'divide' && num2 === 0) {
            e.preventDefault();
            alert('Cannot divide by zero');
            return;
        }

        // Add loading state
        const submitBtn = form.querySelector('button[type="submit"]');
        submitBtn.textContent = 'Calculating...';
        submitBtn.disabled = true;
    });
});

// Clear form function
function clearForm() {
    document.getElementById('num1').value = '';
    document.getElementById('num2').value = '';
    document.getElementById('operation').selectedIndex = 0;
    
    // Remove any result or error messages
    const alerts = document.querySelectorAll('.alert');
    alerts.forEach(alert => alert.remove());
    
    // Focus on first input
    document.getElementById('num1').focus();
}

// Add keyboard shortcut hint
window.addEventListener('load', function() {
    const form = document.querySelector('.calculator-form');
    if (form) {
        const hint = document.createElement('p');
        hint.style.cssText = 'text-align: center; color: #666; font-size: 0.9rem; margin-top: 1rem;';
        hint.textContent = 'Tip: Press Ctrl+Enter (Cmd+Enter on Mac) to calculate';
        form.appendChild(hint);
    }
});

// Made with Bob
