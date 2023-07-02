function validateForm() {
    var username = document.getElementById("username").value;
    var password = document.getElementById("password").value;
    
    if (username.trim() === "" || password.trim() === "") {
        alert("Please enter both username and password.");
        return false;
    }
}