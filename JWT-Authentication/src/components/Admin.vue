<script>
import axios from "axios";
import { ref, onMounted } from "vue";
import { useRouter } from "vue-router";

export default {
  name: "Admin",
  setup() {
    const users = ref([]);
    const router = useRouter();

    // Fetch all users
    const fetchAllUsers = async () => {
      try {
        const response = await axios.get("http://localhost:3000/users");
        users.value = response.data.users;
        console.log(response.data);
      } catch (error) {
        console.error("Error fetching users:", error);
      }
    };

    // Call the fetch function when the component is mounted
    onMounted(() => {
      fetchAllUsers();
    });

    // Logout function to redirect to login page
    const logout = () => {
      router.push({ name: "Login" });
    };

    // Delete user function
    const deleteUser = async (email) => {
      try {
        const response = await axios.delete(`http://localhost:3000/deleteUser/${email}`);
        console.log("User deleted:", response.data);
        fetchAllUsers();  
      } catch (error) {
        console.error("Error deleting user:", error);
      }
    };

    return { users, logout, deleteUser };
  },
};
</script>

<template>
  <div class="container">
    <div class="header">
      <h1 class="title">Admin Panel</h1>
      <button class="logout-button" @click="logout">Logout</button>
    </div>
    <table class="user-table">
      <thead>
        <tr>
          <th>Name</th>
          <th>Email</th>
          <th>Password</th>
          <th>Actions</th> <!-- New column for delete action -->
        </tr>
      </thead>
      <tbody>
        <tr v-for="(user, index) in users" :key="index">
          <td>{{ user.name }}</td>
          <td>{{ user.email }}</td>
          <td>{{ user.password }}</td>
          <td>
            <button class="delete-button" @click="deleteUser(user.email)">Delete</button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<style scoped>
.container {
  width: 80%;
  margin: 20px auto;
  text-align: center;
}

.header {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-bottom: 20px;
}

.title {
  font-size: 24px;
  color: #333;
  flex-grow: 1;
  text-align: center;
}

.logout-button {
  padding: 10px 20px;
  background-color: #ff4757;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

.logout-button:hover {
  background-color: #ff6b81;
}

.user-table {
  width: 100%;
  border-collapse: collapse;
  background: white;
  border-radius: 10px;
  overflow: hidden;
  box-shadow: 0px 4px 6px rgba(0, 0, 0, 0.2);
}

.user-table th,
.user-table td {
  padding: 12px;
  border-bottom: 1px solid #ddd;
}

.user-table th {
  background: #007bff;
  color: white;
  text-transform: uppercase;
}

.user-table tr:hover {
  background: #f1f1f1;
  transition: 0.3s ease-in-out;
}

.delete-button {
  padding: 5px 10px;
  background-color: #ff4757;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

.delete-button:hover {
  background-color: #ff4757;
}
</style>
