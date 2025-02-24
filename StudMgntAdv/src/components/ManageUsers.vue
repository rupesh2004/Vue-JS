<script setup>
import {
    ref,
    onMounted
} from "vue";
import axios from "axios";
import {
    useRouter
} from "vue-router";

const users = ref([]);
const router = useRouter();

const username = ref('');
const firstName = ref('');
const lastName = ref('');
const email = ref('');
const password = ref('');
const role = ref('user');

const formatDate = (dateString) => {
    return new Date(dateString).toLocaleDateString("en-US", {
        year: "numeric",
        month: "short",
        day: "numeric",
    });
};

const goBack = () => {
    router.push({
        name: "AdminDashboard"
    }); // Go to the previous page
};

const updateUser = (user) => {

    router.push({
        name: "UpdateUser",
        query: {
            email: user.email
        }
    })
};

const deleteUser = async (email) => {
    if (!confirm(`Are you sure you want to delete ${email}?`)) return;

    const token = sessionStorage.getItem("jwt_token");

    try {
        await axios.delete(`http://localhost:3000/deleteUser/${email}`, {
            headers: {
                Authorization: `Bearer ${token}`
            },
        });
        users.value = users.value.filter((user) => user.email !== email);
        alert("User deleted successfully!");
    } catch (error) {
        console.error("Error deleting user:", error);
        alert("Failed to delete user.");
    }
};

const fetchAllProfile = async () => {
    const token = sessionStorage.getItem("jwt_token");

    if (!token) {
        alert("Unauthorized! Please log in.");
        router.push("/");
        return;
    }

    try {
        const response = await axios.get("http://localhost:3000/viewAllUsers", {
            headers: {
                Authorization: `Bearer ${token}`
            },
        });
        users.value = response.data;
    } catch (error) {
        console.error("Error fetching users:", error);
    }
}
const register = async () => {
    console.log("Before Sending:", {
        username: username.value,
        firstName: firstName.value,
        lastName: lastName.value,
        email: email.value,
        password: password.value,
        role: role.value
    });

    const user = {
        username: username.value,
        first_name: firstName.value,
        last_name: lastName.value,
        email: email.value,
        password: password.value,
        roles: role.value || "user" // Ensure role is not empty
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
            role.value = "user"; // Reset after success
            alert("Registration Successful");
            fetchAllProfile()
        } else {
            alert("Registration Failed");
        }
    } catch (error) {
        console.error("Error:", error);
        alert("Registration Failed - Check Console for Errors");
    }
};
onMounted(() => {
    fetchAllProfile();
})
</script>

<template>
<nav class="navbar">
    <button class="back-btn" @click="goBack">← Back</button>
    <h3>Student Management System</h3>
</nav>
<div class="container">
    <div class="reg-content">
        <form @submit.prevent="register" class="    ">
            <div class="register-form">
            <h2 class="form-title">Create an Account</h2>

    <div class="form-group">
        <label><i class="fas fa-user"></i> Username</label>
        <input type="text" v-model="username" required placeholder="rupesh2004" />
    </div>

    <div class="form-group-row">
        <div class="form-group">
            <label><i class="fas fa-id-badge"></i> First Name</label>
            <input type="text" v-model="firstName" required  placeholder="Rupesh"/>
        </div>

        <div class="form-group">
            <label><i class="fas fa-id-badge"></i> Last Name</label>
            <input type="text" v-model="lastName" required placeholder="Bhosale"/>
        </div>
    </div>

    <div class="form-group-row">
        <div class="form-group">
            <label><i class="fas fa-envelope"></i> Email</label>
            <input type="email" v-model="email" required placeholder="bhosaerupesh@gmail.com" />
        </div>

        <div class="form-group">
            <label><i class="fas fa-lock"></i> Password</label>
            <input type="password" v-model="password" required placeholder = "★★★★★★★★" />
        </div>
    </div>

    <div class="form-group">
        <label><i class="fas fa-user-tag"></i> Role</label>
        <select v-model="role">
            <option value="user" selected>User</option>
            <option value="admin">Admin</option>
        </select>
    </div>

    <button type="submit" class="submit-btn">Register</button>
</div>
        </form>
    </div>

    <div class="content">
        <h2 style="margin-left: 40px"><i class="fas fa-users"></i> All Users</h2>
        <table>
            <thead>
                <tr>
                    <th><i class="fas fa-user"></i> Username</th>
                    <th><i class="fas fa-id-badge"></i> <br>First Name</th>
                    <th><i class="fas fa-id-badge"></i><br> Last Name</th>
                    <th><i class="fas fa-envelope"></i> <br>Email</th>
                    <th><i class="fas fa-calendar-plus"></i><br> Created On</th>
                    <th><i class="fas fa-calendar-check"></i> <br>Updated On</th>
                    <th><i class="fas fa-toggle-on"></i> <br>Active</th>
                    <th><i class="fas fa-user-tag"></i><br>Role</th>
                    <th><i class="fas fa-user-cog"></i><br> Created By</th>
                    <th><i class="fas fa-tasks"></i><br>Actions</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="user in users" :key="user.email">
                    <td>{{ user.username }}</td>
                    <td>{{ user.first_name }}</td>
                    <td>{{ user.last_name }}</td>
                    <td>{{ user.email }}</td>
                    <td>{{ formatDate(user.created_on) }}</td>
                    <td>{{ formatDate(user.updated_on) }}</td>
                    <td>
                        <span :class="user.is_active ? 'active' : 'inactive'">
                            {{ user.is_active ? "Active" : "Inactive" }}
                        </span>
                    </td>
                    <td>{{ user.roles }}</td>
                    <td>{{ user.created_by }}</td>
                    <td>
                        <button class="update-btn" @click="updateUser(user)">
                            <i class="fas fa-edit"></i> Update
                        </button>
                        <button class="delete-btn" @click="deleteUser(user.email)">
                            <i class="fas fa-trash"></i> Delete
                        </button>
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</div>
<footer class="footer">
    <p>&copy; 2025 Student Management System</p>
</footer>
</template>

<style scoped>
* {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
}

.reg-content {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-grow: 1;
    padding-top: 20px;
    text-align: center;
}

/* Form Styling */
.register-form {
    background: white;
    padding: 20px;
    border-radius: 10px;
    box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.6);
    width: 550px;
    text-align: center;
}

/* Form Title */
.form-title {
    font-size: 22px;
    margin-bottom: 20px;
    color: #333;
}

/* Group Two Inputs in One Row */
.form-group-row {
    display: flex;
    gap: 15px; /* Space between inputs */
}

/* Ensure Inputs Take Equal Space */
.form-group-row .form-group {
    flex: 1;
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

/* Responsive: Stack Inputs on Small Screens */
@media (max-width: 600px) {
    .form-group-row {
        flex-direction: column;
    }
}

.container {
    display: flex;
    flex-direction: column;
    min-height: 100vh;
    min-width: 100%;
    font-family: Arial, sans-serif;
    background-color: #f5f5f5;
}

/* Navbar */
.navbar {
    width: 100%;
    /* No more 100vw to prevent overflow */
    background-color: #007bff;
    color: white;
    padding: 15px 0;
    text-align: center;
    font-size: 24px;
    font-weight: bold;
    display: flex;
    justify-content: center;
    align-items: center;
}

.back-btn {
    position: absolute;
    left: 20px;
    background: white;
    color: #007bff;
    border: none;
    padding: 8px 15px;
    font-size: 16px;
    cursor: pointer;
    border-radius: 5px;
    font-weight: bold;
}

.back-btn:hover {
    background: #e9ecef;
}

/* Content */
.content {
    flex: 1;
    padding: 20px;
    text-align: center;
}

/* Table Styles */
table {
    width: 100%;
    /* Keeps table centered and responsive */
    max-width: 1200px;
    border-collapse: collapse;
    background: white;
    border-radius: 10px;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
    margin: 20px auto auto 10px;
}

th,
td {
    padding: 12px;
    border: 1px solid #ddd;
    text-align: center;
}

th {
    background-color: #007bff;
    color: white;
}

tbody tr:nth-child(even) {
    background-color: #f2f2f2;
}

/* Active & Inactive Status */
.active {
    color: green;
    font-weight: bold;
}

.inactive {
    color: red;
    font-weight: bold;
}

/* Action Buttons */
.update-btn,
.delete-btn {
    border: none;
    padding: 8px 12px;
    margin: 2px;
    font-size: 14px;
    cursor: pointer;
    border-radius: 5px;
    font-weight: bold;
}

.update-btn {
    background-color: #28a745;
    color: white;
}

.update-btn:hover {
    background-color: #218838;
}

.delete-btn {
    background-color: #dc3545;
    color: white;
}

.delete-btn:hover {
    background-color: #c82333;
}

/* Footer */
.footer {
    width: 100%;
    /* Fix overflow issue */
    background-color: #343a40;
    color: white;
    text-align: center;
    padding: 15px 0;
    font-size: 14px;
    /* margin-top: auto; Push footer to the bottom */
}
</style>
