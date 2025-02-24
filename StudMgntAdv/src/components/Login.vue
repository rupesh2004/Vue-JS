<script>
import axios from "axios";
import {
    onMounted,
    ref
} from "vue";
import {
    useRouter
} from "vue-router";

export default {
    name: "Login",
    setup() {
        const email = ref('');
        const password = ref('');
        const router = useRouter()

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
                    const {
                        token,
                        roles
                    } = response.data;

                    if (!roles) {
                        alert("Role not found in response. Please try again.");
                        return;
                    }
                    sessionStorage.setItem("jwt_token", token);
                    sessionStorage.setItem("user_role", roles);
                    if (roles === "user") {
                        router.push({
                            name: "UserDashboard"
                        });
                    } else if (roles === "admin") {
                        router.push({
                            name: "AdminDashboard"
                        });
                    } else {
                        alert("Invalid role");
                    }
                }
            } catch (error) {
                console.error("Login Error:", error);
                alert("Login failed. Check your credentials and try again.");
            }
        }

        return {
            email,
            password,
            login
        };
    }
};
</script>

<template>
        <nav class="navbar">
            <div class="nav-content">
                <h2>Student Portal</h2>
            </div>
        </nav>
    <div class="container">
        <!-- Navbar -->

        <!-- Login Form -->
        <div class="login-wrapper">
            <form @submit.prevent="login" class="login-form">
                <h2 class="form-title">Login In</h2>

                <!-- Login Icon Below Title -->
                <div class="login-icon">
                    <i class="fas fa-user-circle"></i>
                </div>

                <div class="form-group">
                    <label><i class="fas fa-envelope"></i> Email</label>
                    <input type="email" v-model="email" placeholder="joe@gmail.com" required />
                </div>

                <div class="form-group">
                    <label><i class="fas fa-lock"></i> Password</label>
                    <input type="password" v-model="password" placeholder="★★★★★★★★" required />
                </div>

                <button type="submit" class="submit-btn"><i class="fa-solid fa-lock"></i> Login</button>
            </form>
        </div>

        <!-- Footer (Now below login form) -->
    </div>
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
</template>

<style scoped>
/* General Layout */
.container {
    display: flex;
    flex-direction: column;
    min-height: 100vh;
    font-family: 'Arial', sans-serif;
}

/* Navbar */
.navbar {
    width: 100%;
    background-color: #007bff;
    color: white;
    display: flex;
    justify-content: center;
    align-items: center;
    height: 80px;
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
    margin-top: -100px;
}

/* Form styling */
.login-form {
    background: white;
    padding: 25px;
    border-radius: 10px;
    box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.3);
    width: 450px;
    text-align: center;
}

/* Form Title */
.form-title {
    font-size: 24px;
    font-weight: bolder;
    margin-bottom: 0px;
    color: #060505;
}

/* Login Icon */
.login-icon {
    font-size: 80px;
    color: #007bff;
}

/* Input fields */
.form-group {
    margin-bottom: 15px;
    text-align: left;
}

.form-group label {
    font-size: 15px;
    font-weight: bold;
    color: #555;
    display: flex;
    align-items: center;
}

.form-group label i {
    margin-right: 10px;
    color: #007bff;
}

.form-group input {
    width: 100%;
    padding: 12px;
    border: 1px solid #ccc;
    border-radius: 5px;
    font-size: 14px;
}

/* Button styling */
.submit-btn {
    width: 100%;
    background-color: #007bff;
    color: white;
    padding: 12px;
    font-size: 16px;
    border: none;
    border-radius: 5px;
    cursor: pointer;
    transition: 0.3s;
}

.submit-btn:hover {
    background-color: #0056b3;
}

/* Footer - Now Dynamic */
.footer {
    width: 100%;
    background-color: #343a40;
    color: white;
    text-align: center;
    padding: 15px 0;
    margin-top: -90px; /* Pushes footer below login form */
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
