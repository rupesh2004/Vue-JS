<script>
import axios from "axios";
import { ref } from "vue";
import { useRouter } from "vue-router";

export default {
    name: "Register",
    setup() {
        const username = ref('');
        const firstName = ref('');
        const lastName = ref('');
        const email = ref('');
        const password = ref('');
        const role = ref('user'); // Default role
        const router = useRouter();

        const register = async () => {
            console.log("Before Sending:", { username: username.value, firstName: firstName.value, lastName: lastName.value, email: email.value, password: password.value, role: role.value });

            const user = {
                username: username.value,
                first_name: firstName.value,
                last_name: lastName.value,
                email: email.value,
                password: password.value,
                roles: role.value || "user"  // Ensure role is not empty
            };

            try {
                const response = await axios.post("http://localhost:3000/register", user);
                console.log("Response Data:", response.data);

                if (response.status === 200 || response.status === 201) {
                    username.value = "";
                    firstName.value = "";
                    lastName.value = "";
                    email.value = "";
                    password.value = "";
                    role.value = "user";  // Reset after success
                    alert("Registration Successful");
                    router.push({ name: "ManageUsers" });
                } else {
                    alert("Registration Failed");
                }
            } catch (error) {
                console.error("Error:", error);
                alert("Registration Failed - Check Console for Errors");
            }
        };


        return { username, firstName, lastName, email, password, role, register };
    }
};
</script>

<template>
    <div class="page-container">
        <!-- Navbar -->
        <nav class="navbar">
            <div class="nav-content">
                <h2>User Management System</h2>
                <button class="nav-btn" @click="$router.push({ name: 'ManageUsers' })">Back</button>
            </div>
        </nav>

        <!-- Main Content -->
        <div class="content">
            <form @submit.prevent="register" class="register-form">
                <h2 class="form-title">Create an Account</h2>

                <div class="form-group">
                    <label>Username</label>
                    <input type="text" v-model="username" required />
                </div>

                <div class="form-group">
                    <label>First Name</label>
                    <input type="text" v-model="firstName" required />
                </div>

                <div class="form-group">
                    <label>Last Name</label>
                    <input type="text" v-model="lastName" required />
                </div>

                <div class="form-group">
                    <label>Email</label>
                    <input type="email" v-model="email" required />
                </div>

                <div class="form-group">
                    <label>Password</label>
                    <input type="password" v-model="password" required />
                </div>

                <div class="form-group">
                    <label>Role</label>
                    <select v-model="role">
                        <option value="user" selected>User</option>
                        <option value="admin">Admin</option>
                    </select>
                </div>

                <button type="submit" class="submit-btn">Register</button>
            </form>
        </div>

        <!-- Footer (Now below the form) -->
        <footer class="footer">
            <div class="footer-content">
                <p>&copy; 2025 Student Management System. All rights reserved.</p>
                <div class="footer-links">
                    <a href="#">Privacy Policy</a> | 
                    <a href="#">Terms of Service</a> | 
                    <a href="#">Support</a>
                </div>
            </div>
        </footer>
    </div>
</template>

<style scoped>
/* Page Layout */
.page-container {
    display: flex;
    flex-direction: column;
    min-height: 100vh;
    background-color: #f8f9fa;
    font-family: 'Arial', sans-serif;
}

/* Navbar */
.navbar {
    width: 100%;
    background-color: #007bff;
    color: white;
    padding: 15px 20px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    position: fixed;
    top: 0;
    left: 0;
    height: 80px;
    z-index: 1000;
    box-shadow: 0px 4px 8px rgba(0, 0, 0, 0.1);
}

.nav-content {
    width: 90%;
    max-width: 1200px;
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.navbar h2 {
    margin: 0;
    font-size: 22px;
}

.nav-btn {
    background-color: #ff9800;
    color: white;
    border: none;
    padding: 10px 15px;
    font-size: 14px;
    border-radius: 5px;
    cursor: pointer;
    transition: 0.3s;
}

.nav-btn:hover {
    background-color: #e68900;
}

/* Center the form */
.content {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-grow: 1;
    padding-top: 120px;
    text-align: center;
}

/* Form Styling */
.register-form {
    background: white;
    padding: 20px;
    border-radius: 10px;
    box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.1);
    width: 350px;
    text-align: center;
}

/* Form Title */
.form-title {
    font-size: 22px;
    margin-bottom: 20px;
    color: #333;
}

/* Input Fields */
.form-group {
    margin-bottom: 15px;
    text-align: left;
}

.form-group label {
    font-size: 14px;
    font-weight: bold;
    color: #555;
    display: block;
    margin-bottom: 5px;
}

.form-group input,
.form-group select {
    width: 100%;
    padding: 10px;
    border: 1px solid #ccc;
    border-radius: 5px;
    font-size: 14px;
}

/* Button */
.submit-btn {
    width: 100%;
    background-color: #007bff;
    color: white;
    padding: 10px;
    font-size: 16px;
    border: none;
    border-radius: 5px;
    cursor: pointer;
    transition: 0.3s;
}

.submit-btn:hover {
    background-color: #0056b3;
}

/* Login Link */
.login-link {
    margin-top: 15px;
    font-size: 14px;
    color: #555;
}

.login-link a {
    color: #007bff;
    text-decoration: none;
    font-weight: bold;
    transition: 0.3s;
}

.login-link a:hover {
    text-decoration: underline;
}

/* Footer (Now after the form) */
.footer {
    width: 100%;
    background-color: #343a40;
    color: white;
    text-align: center;
    padding: 15px 0;
    margin-top: 20px;
}

.footer-content {
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 20px;
}

.footer-links a {
    color: white;
    text-decoration: none;
    margin: 0 8px;
    font-size: 14px;
}

.footer-links a:hover {
    text-decoration: underline;
}
</style>
