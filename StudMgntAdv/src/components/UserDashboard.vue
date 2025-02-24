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
                <button class="btn-nav" @click="logout">
                    <i class="fas fa-sign-out-alt"></i> Logout
                </button>
            </div>
        </nav>

        <!-- Main Dashboard -->
        <div class="dashboard-container">
            <div class="cards">
                <!-- View Profile Card -->
                <div class="card">
                    <i class="fas fa-user-circle card-icon"></i>
                    <h2>View Profile</h2>
                    <p>Check and update your personal details.</p>
                    <button class="btn" @click="goToProfile">
                        <center><i class="fas fa-id-badge"></i> View Profile</center>
                    </button>
                </div>

                <!-- Register for Exam Card -->
                <div class="card">
                    <i class="fas fa-edit card-icon"></i>
                    <h2>Register for Exam</h2>
                    <p>Apply for upcoming exams easily.</p>
                    <button class="btn" @click="registerForExam">
                        <i class="fas fa-pen"></i> Register for Exam
                    </button>
                </div>
            </div>
        </div>

        <!-- Footer -->
    </div>
        <footer class="footer">
    <div class="footer-content">
        <p>&copy; 2025 Student Management System. All rights reserved.</p>
        <div class="footer-links">
            <a href="#"><i class="fas fa-lock"></i> Privacy Policy</a> 
            <span class="divider">|</span>
            <a href="#"><i class="fas fa-file-contract"></i> Terms of Service</a> 
            <span class="divider">|</span>
            <a href="#"><i class="fas fa-life-ring"></i> Support</a>
        </div>
        <div class="social-icons">
            <a href="#" class="social-link"><i class="fab fa-facebook-f"></i></a>
            <a href="#" class="social-link"><i class="fab fa-twitter"></i></a>
            <a href="#" class="social-link"><i class="fab fa-instagram"></i></a>
            <a href="#" class="social-link"><i class="fab fa-linkedin-in"></i></a>
        </div>
    </div>
</footer>
</template>

<style scoped>
/* General Layout */
.container {
    display: flex;
    min-height: 100vh;
    width: 100%;
    background-color: #f5f5f5;
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
    display: flex;
    align-items: center;
    gap: 10px;
}

/* Logout Button */
.btn-nav {
    background-color: #ffffff;
    color: #007bff;
    border: none;
    padding: 10px 15px;
    font-size: 16px;
    border-radius: 5px;
    cursor: pointer;
    transition: 0.3s;
    display: flex;
    align-items: center;
    gap: 8px;
}

.btn-nav:hover {
    background-color: #0056b3;
    color: white;
}

/* Dashboard Content */
.dashboard-container {
    margin-top: 20px;
    margin-left: -50px;
    text-align: center;
    padding: 100px 20px 20px;
    background-color: #f4f4f4;
}

/* Cards */
.cards {
    display: flex;
    justify-content: center;
    gap: 20px;
}

/* Card Styling */
.card {
    width: 300px;
    background: white;
    padding: 20px;
    border-radius: 10px;
    box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.6);
    text-align: center;
    transition: 0.3s ease-in-out;
    opacity: 0;
    transform: translateY(20px);
    animation: fadeInUp 0.8s forwards;
}

/* Fade-in animation */
@keyframes fadeInUp {
    to {
        opacity: 1;
        transform: translateY(0);
    }
}

/* Card Icons */
.card-icon {
    font-size: 50px;
    color: #007bff;
    margin-bottom: 10px;
    opacity: 0;
    animation: fadeIn 1s forwards 0.3s;
}

@keyframes fadeIn {
    to {
        opacity: 1;
    }
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
    padding: 12px 20px;
    font-size: 16px;
    border-radius: 5px;
    cursor: pointer;
    transition: 0.3s ease-in-out;
    
    /* Flexbox for proper centering */
    display: flex;
    align-items: center;
    justify-content: center; /* Center text and icon */
    gap: 8px;

    /* Fade-in and scale animation */
    opacity: 0;
    transform: scale(0.9);
    animation: fadeInScale 0.6s forwards 0.5s;
}

@keyframes fadeInScale {
    from {
        opacity: 0;
        transform: scale(0.9);
    }
    to {
        opacity: 1;
        transform: scale(1);
    }
}

@keyframes fadeInScale {
    to {
        opacity: 1;
        transform: scale(1);
    }
}

.btn:hover {
    background-color: #0056b3;
}

.footer {
    width: 100%;
    background-color: #343a40;
    color: white;
    text-align: center;
    position: relative;
    bottom: 0;
    left: 0;
    margin-top: -100px;
    box-shadow: 0px -4px 8px rgba(0, 0, 0, 0.2);
}

.footer-content {
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 20px;
}

/* Footer Links */
.footer-links {
    margin-top: 10px;
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 12px;
    flex-wrap: wrap;
}

.footer-links a {
    color: white;
    text-decoration: none;
    font-size: 14px;
    transition: 0.3s;
    display: flex;
    align-items: center;
    gap: 5px;
}

.footer-links a:hover {
    text-decoration: underline;
    color: #007bff;
}

/* Divider */
.divider {
    color: rgba(255, 255, 255, 0.5);
    font-size: 14px;
}

/* Social Icons */
.social-icons {
    margin-top: 15px;
    display: flex;
    justify-content: center;
    gap: 12px;
}

.social-link {
    color: white;
    font-size: 18px;
    width: 35px;
    height: 35px;
    display: flex;
    align-items: center;
    justify-content: center;
    background-color: rgba(255, 255, 255, 0.1);
    border-radius: 50%;
    transition: 0.3s;
}

.social-link:hover {
    background-color: #007bff;
    transform: scale(1.1);
}

/* Responsive */
@media (max-width: 768px) {
    .footer-links {
        flex-direction: column;
        gap: 8px;
    }
}
</style>
