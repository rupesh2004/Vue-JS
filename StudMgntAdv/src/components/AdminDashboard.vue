<script>
import { ref, onMounted } from "vue";
import { useRouter } from "vue-router";
import axios from "axios"; // Ensure axios is installed

export default {
    name: "AdminDashboard",
    setup() {
        const router = useRouter();
        const adminProfile = ref(null); // Store admin details
        const errorMessage = ref("");

        // Function to fetch admin profile
        const fetchAdminProfile = async () => {
            const token = sessionStorage.getItem("jwt_token");
            if (!token) {
                alert("Unauthorized! Please log in.");
                router.push({ name: "Login" });
                return;
            }

            try {
                const response = await axios.get("http://localhost:3000/adminProfile", {
                    headers: { Authorization: `Bearer ${token}` },
                });
                adminProfile.value = response.data;
            } catch (error) {
                console.error("Error fetching admin profile:", error);
                errorMessage.value = "Failed to load admin profile.";
            }
        };
        const manageUsers=()=>{
            router.push({ name: "ManageUsers" });
        }

        // Logout function
        const logout = () => {
            sessionStorage.removeItem("jwt_token");
            sessionStorage.removeItem("user_role");
            alert("Logged out successfully!");
            router.push({ name: "Login" });
        };

        const createExam  = ()=>{
            router.push({ name: "CreateExam" });
        }
        

        // Fetch admin profile on component mount
        onMounted(fetchAdminProfile);

        return { adminProfile, errorMessage, logout, manageUsers, createExam };
    },
};
</script>


<template>
    <div class="dashboard-container">
        <!-- Navbar -->
        <nav class="navbar">
            <div class="nav-content">
                <h2>Admin Dashboard</h2>
                <button class="logout-btn" @click="logout">
                    <i class="fas fa-sign-out-alt"></i> Logout
                </button>
            </div>
        </nav>

        <!-- Main Content -->
        <div class="content">
            <div class="cards">
                <!-- User Management Card -->
                <div class="card">
                    <i class="fas fa-users-cog card-icon"></i>
                    <h3>User Management</h3>
                    <p>Manage registered users.</p>
                    <button class="btn" @click="manageUsers">
                        <i class="fas fa-users"></i> Manage Users
                    </button>
                </div>

                <!-- Exam Management Card -->
                <div class="card">
                    <i class="fas fa-file-alt card-icon"></i>
                    <h3>Exam Management</h3>
                    <p>Manage new exams for students.</p>
                    <button class="btn" @click="createExam">
                        <i class="fas fa-pencil-alt"></i> Manage Exam
                    </button>
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
.dashboard-container {
    display: flex;
    flex-direction: column;
    min-height: 100vh;
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
    font-size: 24px;
}

/* Logout Button */
.logout-btn {
    background-color: #dc3545;
    color: white;
    border: none;
    padding: 10px 15px;
    font-size: 14px;
    border-radius: 5px;
    cursor: pointer;
    transition: 0.3s;
}

.logout-btn:hover {
    background-color: #c82333;
}

/* Content Section */
.content {
    align-items: center;
    flex-grow: 1;
    height: 100vh;
    text-align: center;
    padding: 80px 20px;
}

/* Card container */
.cards {
    display: flex;
    margin-top: 30px;
    gap: 30px;
    animation: fadeIn 1s ease-in-out;
}

/* Individual card styling */
.card {
    width: 300px;
    background: white;
    padding: 20px;
    border-radius: 10px;
    box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.1);
    transition: 0.3s;
    transform: translateY(0);
    text-align: center;
}

.card:hover {
    transform: translateY(-3px);
}

/* Card Icons */
.card-icon {
    font-size: 40px;
    color: #007bff;
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
    .content {
        padding-top: 100px;
    }

    .cards {
        flex-direction: column;
        align-items: center;
    }

    .card {
        width: 90%;
    }
}

/* Animations */
@keyframes fadeIn {
    from {
        opacity: 0;
    }
    to {
        opacity: 1;
    }
}
</style>
