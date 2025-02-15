<template>
      <nav class="navbar">
        <button class="back-btn" @click="goBack">← Back</button>
        <h1>Student Management System</h1>
      </nav>
    <div class="container">
      <!-- Navbar -->
  
      <!-- User List Section -->
      <div class="content">
        <h2 style="margin-left: 90px">All Users</h2>
        <table>
          <thead>
            <tr>
              <th>Username</th>
              <th>First Name</th>
              <th>Last Name</th>
              <th>Email</th>
              <th>Created On</th>
              <th>Updated On</th>
              <th>Active</th>
              <th>Role</th>
              <th>Created By</th>
              <th>Actions</th>
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
                <button class="update-btn" @click="updateUser(user)">Update</button>
                <button class="delete-btn" @click="deleteUser(user.email)">Delete</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
  
      <!-- Footer (100% width) -->
    </div>
      <footer class="footer">
        <p>&copy; 2025 Student Management System</p>
      </footer>
</template>

<script setup>
import { ref, onMounted } from "vue";
import axios from "axios";
import { useRouter } from "vue-router";

const users = ref([]);
const router = useRouter();

const formatDate = (dateString) => {
  return new Date(dateString).toLocaleDateString("en-US", {
    year: "numeric",
    month: "short",
    day: "numeric",
  });
};

const goBack = () => {
  router.push({name : "ManageUsers"}); // Go to the previous page
};

const updateUser = (user) => {
 
  router.push({
    name : "UpdateUser",
    query : {email : user.email}
  })
};

const deleteUser = async (email) => {
  if (!confirm(`Are you sure you want to delete ${email}?`)) return;

  const token = sessionStorage.getItem("jwt_token");

  try {
    await axios.delete(`http://localhost:3000/deleteUser/${email}`, {
      headers: { Authorization: `Bearer ${token}` },
    });
    users.value = users.value.filter((user) => user.email !== email);
    alert("User deleted successfully!");
  } catch (error) {
    console.error("Error deleting user:", error);
    alert("Failed to delete user.");
  }
};

onMounted(async () => {
  const token = sessionStorage.getItem("jwt_token");

  if (!token) {
    alert("Unauthorized! Please log in.");
    router.push("/");
    return;
  }

  try {
    const response = await axios.get("http://localhost:3000/viewAllUsers", {
      headers: { Authorization: `Bearer ${token}` },
    });
    users.value = response.data;
  } catch (error) {
    console.error("Error fetching users:", error);
  }
});
</script>

<style scoped>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

/* General Styles */
html, body {
  width: 100%;
  overflow-x: hidden; /* Prevent horizontal scrolling */
}

.container {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  font-family: Arial, sans-serif;
  background-color: #f8f9fa;
}

/* Navbar */
.navbar {
  width: 100%; /* No more 100vw to prevent overflow */
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
  width: 90%; /* Keeps table centered and responsive */
  max-width: 1200px;
  border-collapse: collapse;
  background: white;
  
  border-radius: 10px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  margin: 20px auto auto 110px;
}

th, td {
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
.update-btn, .delete-btn {
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
  width: 100%; /* Fix overflow issue */
  background-color: #343a40;
  color: white;
  text-align: center;
  padding: 15px 0;
  font-size: 14px;
  margin-top: auto; /* Push footer to the bottom */
}


</style>
