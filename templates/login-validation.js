function loginValidateForm() {
    var username = document.getElementById("username").value;
    var password = document.getElementById("password").value;
    console.log("Username value:", document.getElementById("username").value.trim());
    console.log("Password value:", document.getElementById("password").value.trim());

    if (username.trim() === "" || password.trim() === "") {
        alert("Please enter both username and password.");
        return false;
    }
}