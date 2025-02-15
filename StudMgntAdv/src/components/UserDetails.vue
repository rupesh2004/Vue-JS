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
                <h2>User Profile</h2>
                <button class="btn-nav" @click="goToDashboard">🏠 Dashboard</button>
            </div>
        </nav>

        <!-- Profile Card -->
        <div class="profile-container">
            <div class="profile-card">
                <h3>{{ userData.name }}</h3>
                <div class="details">
                    <p><strong>Username:</strong> {{ userData.username }}</p>
                    <p><strong>Name :</strong> {{ userData.first_name }} {{ userData.last_name }}</p>
                    <p><strong>Email:</strong> {{ userData.email }}</p>
                    <p><strong>Role:</strong> {{ userData.roles }}</p>
                    <p><strong>Account Created:</strong> {{ new Date(userData.created_on).toLocaleString() }}</p>
                    <p><strong>Last Updated:</strong> {{ new Date(userData.updated_on).toLocaleString() }}</p>
                    <p><strong>Created By:</strong> {{ userData.created_by }}</p>
                    
                    <p>
                        <strong>Status:</strong> 
                        <span :class="{'active': userData.is_active, 'inactive': !userData.is_active}">
                            {{ userData.is_active ? 'Active' : 'Inactive' }}
                        </span>
                    </p>
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
/* General Styles */
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

/* Dashboard Button */
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

/* Profile Container */
.profile-container {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-grow: 1;
    padding: 20px;
}

/* Profile Card */
.profile-card {
    background: white;
    padding: 25px;
    border-radius: 10px;
    box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.6);
    text-align: center;
    width: 400px;
}

.profile-card h3 {
    font-size: 26px;
    margin-bottom: 15px;
    color: #333;
}

.details p {
    font-size: 16px;
    color: #555;
    margin-bottom: 10px;
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
