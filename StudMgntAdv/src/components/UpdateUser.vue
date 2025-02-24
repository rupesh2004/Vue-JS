<script>
import { useRoute, useRouter } from "vue-router";
import { ref, onMounted } from "vue";
import axios from "axios";

export default {
    name: "UpdateUser",
    setup() {
        const route = useRoute();
        const router = useRouter();
        const email = ref(route.query.email);
        const userData = ref({});
        const newPassword = ref("");

        const handleUnauthorizedError = () => {
            alert("Unauthorized User. Please Login");
            sessionStorage.removeItem("jwt_token"); 
            router.push({ name: "Login" }); 
        };

        // Fetch user data
        const fetchEmailProfile = async () => {
            if (!email.value) return;
            try {
                const token = sessionStorage.getItem("jwt_token");
                const response = await axios.get(`http://localhost:3000/getUserForUpdate/${email.value}`, {
                    headers: { Authorization: `Bearer ${token}` },
                });
                userData.value = response.data;
            } catch (error) {
                console.error("Error fetching user data:", error);
                if (error.response && error.response.status === 401) {
                    handleUnauthorizedError();
                }
            }
        };

        // Update user data
        const updateUserProfile = async () => {
            try {
                const token = sessionStorage.getItem("jwt_token");
                const updatedData = { ...userData.value };
                if (newPassword.value) {
                    updatedData.password = newPassword.value;
                }

                await axios.put(`http://localhost:3000/updateUser/${email.value}`, updatedData, {
                    headers: { Authorization: `Bearer ${token}` },
                });

                alert("User updated successfully!");
                router.push({ name: "ManageUsers" });
            } catch (error) {
                console.error("Error updating user:", error);
                if (error.response && error.response.status === 401) {
                    handleUnauthorizedError();
                } else {
                    alert("Failed to update user.");
                }
            }
        };

        onMounted(() => {
            fetchEmailProfile();
        });

        return {
            userData,
            newPassword,
            updateUserProfile,
        };
    },
};
</script>

<template>
    <div class="container">
        <!-- Navbar -->
        <nav class="navbar">
            <div class="nav-content">
                <h2>Update User</h2>
                <button class="btn-nav" @click="$router.push({ name: 'ManageUsers' })">
                    <i class="fas fa-arrow-left"></i> Back
                </button>
            </div>
        </nav>

        <!-- Form -->
        <div class="form-container">
            <form @submit.prevent="updateUserProfile" class="form">
                <div class="form-group">
                    <label><i class="fas fa-user"></i> Username</label>
                    <input v-model="userData.username" type="text" required />
                </div>

                <div class="form-group-row">
                    <div class="form-group">
                        <label><i class="fas fa-id-badge"></i> First Name</label>
                        <input v-model="userData.first_name" type="text" required />
                    </div>

                    <div class="form-group">
                        <label><i class="fas fa-id-badge"></i> Last Name</label>
                        <input v-model="userData.last_name" type="text" required />
                    </div>
                </div>

                <div class="form-group">
                    <label><i class="fas fa-envelope"></i> Email (read-only)</label>
                    <input v-model="userData.email" type="email" readonly />
                </div>

                <div class="form-group-row">
                    <div class="form-group">
                        <label><i class="fas fa-user-tag"></i> Role</label>
                        <select v-model="userData.roles">
                            <option value="admin">Admin</option>
                            <option value="user">User</option>
                        </select>
                    </div>

                    <div class="form-group">
                        <label><i class="fas fa-check-circle"></i> Active</label>
                        <select v-model="userData.is_active">
                            <option :value="true">Active</option>
                            <option :value="false">Inactive</option>
                        </select>
                    </div>
                </div>

                <div class="form-group">
                    <label><i class="fas fa-lock"></i> New Password (leave blank to keep current)</label>
                    <input v-model="newPassword" type="password" />
                </div>

                <button type="submit">
                    <i class="fas fa-save"></i> Update
                </button>
            </form>
        </div>

        <!-- Footer -->
    </div>
        <footer class="footer">
            <div class="footer-content">
                <p>&copy; 2025 Student Management System. All rights reserved.</p>
                <div class="footer-links">
                    <a href="#"><i class="fas fa-user-secret"></i> Privacy Policy</a> |
                    <a href="#"><i class="fas fa-file-contract"></i> Terms of Service</a> |
                    <a href="#"><i class="fas fa-headset"></i> Support</a>
                </div>
            </div>
        </footer>
</template>

<style scoped>
/* Main Container */
.container {
    display: flex;
    flex-direction: column;
    min-height: 100vh;
    min-width: 100%;
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
}

.btn-nav i {
    margin-right: 5px;
}

.btn-nav:hover {
    background-color: #0056b3;
    color: white;
}

/* Form Container */
.form-container {
    max-width: 500px;
    margin: 100px auto 50px;
    padding: 20px;
    background: white;
    border-radius: 10px;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

/* Form Styling */
.form {
    display: flex;
    flex-direction: column;
}

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

/* Two Inputs in One Row */
.form-group-row {
    display: flex;
    gap: 15px;
}

.form-group-row .form-group {
    flex: 1;
}

/* Button */
.form button {
    width: 100%;
    margin-top: 20px;
    padding: 10px;
    background: #007bff;
    color: white;
    border: none;
    cursor: pointer;
    border-radius: 5px;
    transition: 0.3s;
    font-size: 16px;
}

.form button i {
    margin-right: 5px;
}

.form button:hover {
    background: #0056b3;
}

/* Footer */
.footer {
    width: 100%;
    background-color: #343a40;
    color: white;
    text-align: center;
    padding: 15px 0;
    position: relative;
    bottom: 0;
    left: 0;
    box-sizing: border-box;
}

.footer-content {
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 40px;
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

.footer-links a i {
    margin-right: 5px;
}

.footer-links a:hover {
    text-decoration: underline;
}

/* Responsive: Stack Inputs on Small Screens */
@media (max-width: 600px) {
    .form-group-row {
        flex-direction: column;
    }
}
</style>
