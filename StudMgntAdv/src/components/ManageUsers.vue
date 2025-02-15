<script>
import { useRouter } from "vue-router";

export default {
    name: "ManageUsers",
    setup() {
        const router = useRouter();

        const addUser = () => {
            router.push({ name: "Register" });
        };

        const viewUsers = () => {
            router.push({ name: "ViewAllUsers" });
        };

        return { addUser, viewUsers };
    }
};
</script>

<template>
    <div class="user-management-container">
        <!-- Navbar -->
        <nav class="navbar">
            <div class="nav-content">
                <h2>User Management</h2>
                <button class="back-btn" @click="$router.push({ name: 'AdminDashboard' })">Dashboard</button>
            </div>
        </nav>

        <!-- Main Content -->
        <div class="content">
            <h1>Manage Users</h1>
            <div class="cards">
                <!-- Add User -->
                <div class="card" @click="addUser">
                    <h2>Add User</h2>
                    <p>Create a new user.</p>
                    <button class="btn">Add</button>
                </div>

                <!-- View Users -->
                <div class="card" @click="viewUsers">
                    <h2>View Users</h2>
                    <p>See all registered users.</p>
                    <button class="btn">View</button>
                </div>
            </div>
        </div>

        <!-- Footer -->
        <footer class="footer">
            <p>&copy; 2025 Student Management System. All rights reserved.</p>
        </footer>
    </div>
</template>

<style scoped>
/* General Layout */
.user-management-container {
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
    font-size: 24px;
}

/* Back Button */
.back-btn {
    background-color: rgb(237, 215, 18);

    color: black;
    border: none;
    padding: 10px 15px;
    font-size: 14px;
    border-radius: 5px;
    cursor: pointer;
    transition: 0.3s;
}

.back-btn:hover {
    background-color: #e0a800;
}

/* Content Section */
.content {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    flex-grow: 1;
    height: 100vh;
    text-align: center;
    padding: 100px 20px;
}

h1 {
    margin-bottom: 30px;
    color: #333;
    font-size: 28px;
}

/* Card Container */
.cards {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 20px;
    flex-wrap: wrap;
}

/* Card Styling */
.card {
    width: 250px;
    background: white;
    padding: 20px;
    border-radius: 10px;
    box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.1);
    text-align: center;
    transition: transform 0.3s, box-shadow 0.3s;
    cursor: pointer;
}

.card:hover {

    box-shadow: 0px 6px 15px rgba(0, 0, 0, 0.6);
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
}
</style>
