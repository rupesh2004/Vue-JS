<script>
import { ref, onMounted } from "vue";
import { useRouter } from "vue-router";
import axios from "axios";

export default {
    name: "UserDashboard",
    setup() {
        const user = ref({});
        const router = useRouter();

        const fetchUserProfile = async () => {
            const token = sessionStorage.getItem("jwt_token");
            if (!token) {
                alert("Unauthorized! Please log in.");
                router.push({ name: "Login" });
                return;
            }

            try {
                const response = await axios.get("http://localhost:3000/userProfile", {
                    headers: { Authorization: `Bearer ${token}` }
                });

                user.value = response.data;
                console.log(user.value);
            } catch (error) {
                console.error("Error fetching user profile:", error);
                alert("Failed to fetch profile.");
            }
        };

        const goToProfile = () => {
            router.push({
                name: "UserDetails",
                query: { userData: JSON.stringify(user.value) } // Convert object to string
            });
        };

        const registerForExam = () => {
            router.push({ name: "ExamRegistration" });
        };

        const logout = () => {
            sessionStorage.removeItem("jwt_token");
            sessionStorage.removeItem("user_role");
            alert("Logged out successfully!");
            router.push({ name: "Login" });
        };

        onMounted(fetchUserProfile);

        return { user, goToProfile, registerForExam, logout };
    }
};
</script>

<template>
    <div class="container">
        <!-- Navbar -->
        <nav class="navbar">
            <div class="nav-content">
                <h2>Dashboard</h2>
                <button class="btn-nav" @click="logout">🚪 Logout</button>
            </div>
        </nav>

        <!-- Main Dashboard -->
        <div class="dashboard-container">
            <h1>Welcome, {{ user.name }}</h1>

            <div class="cards">
                <!-- View Profile Card -->
                <div class="card">
                    <h2>View Profile</h2>
                    <p>Check and update your personal details.</p>
                    <button class="btn" @click="goToProfile">View Profile</button>
                </div>

                <!-- Register for Exam Card -->
                <div class="card">
                    <h2>Register for Exam</h2>
                    <p>Apply for upcoming exams easily.</p>
                    <button class="btn" @click="registerForExam">Register for Exam</button>
                </div>
            </div>
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
    justify-content: space-between;
    align-items: center;
}

.navbar h2 {
    margin: 0;
    font-size: 24px;
}

.btn-nav {
    background-color: #ffffff;
    color: #007bff;
    border: none;
    padding: 10px 15px;
    font-size: 16px;
    border-radius: 5px;
    cursor: pointer;
    transition: 0.3s;
}

.btn-nav:hover {
    background-color: #0056b3;
    color: white;
}

/* Dashboard Content */
.dashboard-container {
    text-align: center;
    margin-top: 40px;
    padding: 80px 20px 20px; /* Adjusted padding for navbar */
    background-color: #f4f4f4;
    flex-grow: 1;
}

/* Cards */
.cards {
    display: flex;
    justify-content: center;
    gap: 20px;
    margin-bottom: 20px;
}

.card {
    width: 300px;
    background: white;
    padding: 20px;
    margin-top: 30px;
    border-radius: 10px;
    box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.6);
    text-align: center;
    transition: 0.3s;
}

.card h2 {
    font-size: 20px;
    margin-bottom: 10px;
}

.card p {
    font-size: 14px;
    color: #666;
    margin-bottom: 15px;
}

/* Buttons */
.btn {
    background-color: #007bff;
    color: white;
    border: none;
    padding: 10px 20px;
    font-size: 16px;
    border-radius: 5px;
    cursor: pointer;
    transition: 0.3s;
}

.btn:hover {
    background-color: #0056b3;
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
    .cards {
        flex-direction: column;
        align-items: center;
    }

    .card {
        width: 90%;
    }

    .nav-content {
        flex-direction: column;
        text-align: center;
    }

    .btn-nav {
        margin-top: 10px;
    }
}
</style>
