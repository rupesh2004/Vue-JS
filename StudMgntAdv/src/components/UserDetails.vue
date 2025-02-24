<script>
import { defineComponent, ref, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';

export default defineComponent({
    name: "UserDetails",
    setup() {
        const router = useRouter();
        const route = useRoute();
        const userData = ref({});

        onMounted(() => {
            if (route.query.userData) {
                try {
                    userData.value = JSON.parse(route.query.userData); // Convert string back to object
                } catch (error) {
                    console.error("Error parsing user data:", error);
                }
            } else {
                alert("No user data found!");
                router.push({ name: "UserDashboard" });
            }
        });

        const goToDashboard = () => {
            router.push({ name: "UserDashboard" });
        };

        return { userData, goToDashboard };
    }
});
</script>

<template>
    <div class="container">
        <!-- Navbar -->
        <nav class="navbar">
            <div class="nav-content">
                <h2><i class="fas fa-user-circle"></i> User Profile</h2>
                <button class="btn-nav" @click="goToDashboard">
                    <i class="fas fa-home"></i> Dashboard
                </button>
            </div>
        </nav>

        <!-- Profile Card -->
        <div class="profile-container">
            <div class="profile-card">
                <div class="profile-header">
                    <i class="fas fa-user-alt profile-icon"></i>
                    <h3>{{ userData.name }}</h3>
                </div>
                <div class="details">
                    <div class="detail-item">
                        <i class="fas fa-user-tag"></i>
                        <span><strong>Username:</strong> {{ userData.username }}</span>
                    </div>
                    <div class="detail-item">
                        <i class="fas fa-user"></i>
                        <span><strong>Name:</strong> {{ userData.first_name }} {{ userData.last_name }}</span>
                    </div>
                    <div class="detail-item">
                        <i class="fas fa-envelope"></i>
                        <span><strong>Email:</strong> {{ userData.email }}</span>
                    </div>
                    <div class="detail-item">
                        <i class="fas fa-user-shield"></i>
                        <span><strong>Role:</strong> {{ userData.roles }}</span>
                    </div>
                    <div class="detail-item">
                        <i class="fas fa-calendar-plus"></i>
                        <span><strong>Account Created:</strong> {{ new Date(userData.created_on).toLocaleString() }}</span>
                    </div>
                    <div class="detail-item">
                        <i class="fas fa-edit"></i>
                        <span><strong>Last Updated:</strong> {{ new Date(userData.updated_on).toLocaleString() }}</span>
                    </div>
                    <div class="detail-item">
                        <i class="fas fa-user-cog"></i>
                        <span><strong>Created By:</strong> {{ userData.created_by }}</span>
                    </div>
                    <div class="detail-item">
                        <i class="fas fa-user-check"></i>
                        <span>
                            <strong>Status:</strong>
                            <span :class="{'active': userData.is_active, 'inactive': !userData.is_active}">
                                {{ userData.is_active ? 'Active' : 'Inactive' }}
                            </span>
                        </span>
                    </div>
                </div>
            </div>
        </div>

        <!-- Footer -->
    </div>
        <footer class="footer">
            <div class="footer-content">
                <p>&copy; 2025 Student Management System. All rights reserved.</p>
                <div class="footer-links">
                    <a href="#"><i class="fas fa-user-secret"></i> Privacy Policy</a> | 
                    <a href="#"><i class="fas fa-balance-scale"></i> Terms of Service</a> | 
                    <a href="#"><i class="fas fa-life-ring"></i> Support</a>
                </div>
            </div>
        </footer>
</template>

<style scoped>
/* General Styles */
.container {
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

/* Profile Container */
.profile-container {
    flex: 1;
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 100px 20px 20px;
}

/* Profile Card */
.profile-card {
    background: white;
    padding: 25px;
    border-radius: 10px;
    box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.6);
    width: 400px;
    animation: fadeIn 0.6s ease-in-out;
}

.profile-header {
    display: flex;
    flex-direction: column;
    align-items: center;
    margin-bottom: 15px;
}

.profile-icon {
    font-size: 50px;
    color: #007bff;
    margin-bottom: 10px;
}

.profile-card h3 {
    font-size: 26px;
    color: #333;
}

/* Profile Details */
.details {
    display: flex;
    flex-direction: column;
    gap: 10px;
}

.detail-item {
    display: flex;
    align-items: center;
    gap: 10px;
    font-size: 16px;
    color: #555;
    border-bottom: 1px solid #ddd;
    padding-bottom: 5px;
}

.detail-item i {
    font-size: 18px;
    color: #007bff;
    width: 25px;
    text-align: center;
}

/* Status */
.active {
    color: green;
    font-weight: bold;
}
.inactive {
    color: red;
    font-weight: bold;
}

/* Footer */
.footer {
    width: 100%;
    background-color: #343a40;
    color: white;
    text-align: center;
    padding: 15px 0;
    margin-top: auto;
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
    display: inline-flex;
    align-items: center;
    gap: 5px;
}

.footer-links a:hover {
    text-decoration: underline;
}

/* Fade-in Animation */
@keyframes fadeIn {
    from {
        opacity: 0;
        transform: translateY(-10px);
    }
    to {
        opacity: 1;
        transform: translateY(0);
    }
}

/* Responsive */
@media (max-width: 768px) {
    .profile-card {
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

