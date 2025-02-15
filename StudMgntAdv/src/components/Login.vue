<script>
import axios from "axios";
import { ref } from "vue";
import { useRouter } from "vue-router";

export default {
    name: "Login",
    setup() {
        const email = ref('');
        const password = ref('');
        const router = useRouter();

        const login = async () => {
            if (!email.value || !password.value) {
                alert("Please enter both email and password");
                return;
            }

            try {
                const response = await axios.post("http://localhost:3000/login", {
                    email: email.value,
                    password: password.value
                });

                if (response.status === 200 && response.data.token) {
                    const { token, roles } = response.data;

                    if (!roles) {
                        alert("Role not found in response. Please try again.");
                        return;
                    }
                    sessionStorage.setItem("jwt_token", token);
                    sessionStorage.setItem("user_role", roles); 

                    alert("Login successful!");
                    if (roles === "user") {
                        router.push({ name: "UserDashboard" });
                    } else if (roles === "admin") {
                        router.push({ name: "AdminDashboard" });
                    } else {
                        alert("Invalid role");
                    }
                }
            } catch (error) {
                console.error("Login Error:", error);
                alert("Login failed. Check your credentials and try again.");
            }
        };

        return { email, password, login };
    }
};
</script>

<template>
    <div class="container">
        <!-- Navbar -->
        <nav class="navbar">
            <div class="nav-content">
                <h2>Student Portal</h2>
            </div>
        </nav>

        <!-- Login Form -->
        <div class="login-wrapper">
            <form @submit.prevent="login" class="login-form">
                <h2 class="form-title">Login to Your Account</h2>

                <div class="form-group">
                    <label>Email</label>
                    <input type="email" v-model="email" required />
                </div>

                <div class="form-group">
                    <label>Password</label>
                    <input type="password" v-model="password" required />
                </div>

                <button type="submit" class="submit-btn">Login</button>
            </form>
        </div>

        <!-- Footer -->
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
/* General Layout */
.container {
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
    justify-content: center;
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
    justify-content: center;
    align-items: center;
}

.navbar h2 {
    margin: 0;
    font-size: 24px;
}

/* Center Login Form */
.login-wrapper {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-grow: 1;
    padding: 80px 20px 20px; /* Adjusted padding for navbar */
}

/* Form styling */
.login-form {
    background: white;
    padding: 20px;
    border-radius: 10px;
    box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.6);
    width: 350px;
    margin-top: -40px;
    margin-bottom: 40px;
    text-align: center;
}

/* Form Title */
.form-title {
    font-size: 22px;
    margin-bottom: 20px;
    color: #333;
}

/* Input fields */
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

.form-group input {
    width: 100%;
    padding: 10px;
    border: 1px solid #ccc;
    border-radius: 5px;
    font-size: 14px;
}

/* Button styling */
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

/* Register link */
.register-link {
    margin-top: 15px;
    font-size: 14px;
    color: #555;
}

.register-link a {
    color: #007bff;
    text-decoration: none;
    font-weight: bold;
    transition: 0.3s;
}

.register-link a:hover {
    text-decoration: underline;
}

/* Footer */
.footer {
    width: 100%;
    background-color: #343a40;
    color: white;
    text-align: center;
    padding: 15px 0;
    position: fixed;
    bottom: 0;
    left: 0;
    z-index: 1000;
}

.footer-content {
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 20px;
}

.footer-links {
    margin-top: 5px;
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

/* Responsive */
@media (max-width: 768px) {
    .nav-content {
        flex-direction: column;
        text-align: center;
    }

    .login-form {
        width: 90%;
    }
}
</style>
