<script>
import axios from 'axios';
import { onMounted, ref, nextTick } from 'vue';

export default {
    name: "Dashboard",
    setup() {
        const username = ref("");
        const email = ref("");
        const oldPassword = ref("")
        const newPassword = ref("")
        const confirmPassword = ref("")
        const changePasswordModel = ref(false);

        // Fetch user profile
        const fetchProfileInfo = async () => {
            try {
                const token = sessionStorage.getItem("token");
                const response = await axios.get("http://localhost:3000/r/profile", {
                    headers: {
                        "Authorization": `Bearer ${token}`
                    }
                });

                if (response.status === 200) {
                    username.value = response.data.name;
                    email.value = response.data.email;
                } else {
                    alert("Failed to fetch profile info");
                }
            } catch (error) {
                console.error("Error fetching profile:", error);
            }
        };

        // Open Change Password Modal
        const openChangePasswordModal = async () => {
            console.log("Opening modal...");
            changePasswordModel.value = true;
            await nextTick(); // Ensures Vue updates DOM
            console.log("Modal State:", changePasswordModel.value);
        };

        // Close Change Password Modal
        const closeChangePasswordModal = () => {
            console.log("Closing modal...");
            oldPassword.value = "";
            newPassword.value = "";
            confirmPassword.value = "";
            changePasswordModel.value = false;
        };

        // Logout
        const logout = () => {
            sessionStorage.removeItem("token");
            location.reload();
        };
        const submitUpdatedPassword = async () => {
    if (!oldPassword.value || !newPassword.value || !confirmPassword.value) {
        alert("Please fill in all fields");
        return;
    }

    if (newPassword.value !== confirmPassword.value) {
        alert("New passwords do not match");
        return;
    }

    try {
        const token = sessionStorage.getItem("token");


        const response = await axios.post("http://localhost:3000/r/changepassword", {
                    oldPassword: oldPassword.value,
                    newPassword: newPassword.value,
                    confirmPassword: confirmPassword.value
                }, { headers: { "Authorization": `Bearer ${token}` } });

        if (response.status === 200) {
            alert("Password updated successfully");
            closeChangePasswordModal();
        } else {
            alert(response.data.error || "Failed to update password");
        }

    } catch (error) {
        console.error("Error:", error);
        alert(error.response?.data?.error || "Something went wrong");
    }
};


        onMounted(fetchProfileInfo);

        return {
            username,
            email,
            changePasswordModel,
            openChangePasswordModal,
            closeChangePasswordModal,
            logout,
            oldPassword,
            newPassword,
            confirmPassword,
            submitUpdatedPassword

        };
    },
};
</script>

<template>
  <div class="dashboard-container">
    <div class="card">
      <h2 class="welcome-title">Welcome, {{ username }}!</h2>
      <div class="user-info">
        <p class="user-detail"><strong>Name:</strong> {{ username }}</p>
        <p class="user-detail"><strong>Email:</strong> {{ email }}</p>
      </div>
      <div class="button-group">
        <button @click="openChangePasswordModal" class="btn btn-change-password">
          Change Password
        </button>
        <button @click="logout" class="btn btn-logout">
          Logout
        </button>
      </div>
    </div>
  </div>

  <!-- Change Password Modal -->
  <div v-if="changePasswordModel" class="modal-overlay">
    <div class="modal">
      <h3>Change Password</h3>
      <input type="password" placeholder="Enter Old Password" class="input-field" v-model="oldPassword">
      <input type="password" placeholder="Enter New Password" class="input-field" v-model="newPassword">
      <input type="password" placeholder="Confirm New Password" class="input-field" v-model="confirmPassword">
      <div class="modal-buttons">
        <button class="btn btn-submit" @click="submitUpdatedPassword">Submit</button>
        <button @click="closeChangePasswordModal" class="btn btn-cancel">Cancel</button>
      </div>
    </div>
  </div>
</template>

<style scoped>
/* Dashboard Layout */
.dashboard-container {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    background-color: #f8f9fa;
}

/* Card Styling */
.card {
    background-color: white;
    border-radius: 12px;
    padding: 20px;
    box-shadow: 0 4px 20px rgba(0, 0, 0, 0.2);
    width: 24rem;
    text-align: center;
}

.welcome-title {
    font-size: 2rem;
    font-weight: 700;
    color: #343a40;
    margin-bottom: 1.5rem;
}

.user-info {
    color: #495057;
    margin-bottom: 2rem;
}

.user-detail {
    font-size: 1.125rem;
    font-weight: 500;
    margin: 0.5rem 0;
}

/* Buttons */
.button-group {
    display: flex;
    flex-direction: column;
    gap: 1rem;
}

.btn {
    display: inline-block;
    padding: 1rem 1.5rem;
    font-size: 1.125rem;
    font-weight: 600;
    border-radius: 1.5rem;
    width: 100%;
    text-align: center;
    cursor: pointer;
    transition: all 0.3s ease;
    box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
    border: none;
}

.btn-change-password {
    background-color: #007bff;
    color: white;
}

.btn-change-password:hover {
    background-color: #0056b3;
    color: white;

}

.btn-logout {
    background-color: #dc3545;
    color: white;
}

.btn-logout:hover {
    background-color: #c82333;
    color: white;

}

/* Modal Overlay (FIXED) */
.modal-overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 20%;
    height: 70%;
  
    display: flex !important;  /* FORCES DISPLAY */
    justify-content: center;
    align-items: center;
    z-index: 9999; /* Ensures modal appears above everything */
}

/* Modal Box (FIXED) */
.modal {
    position: absolute;  /* FIXES POSITION */
    background: white;
    padding: 20px;
    margin-left: 400px;
    margin-top: 100px;
    border-radius: 8px;
    width: 30rem;
    box-shadow: 0 4px 10px rgba(0, 0, 0, 0.3);
    text-align: center;
    display: flex; /* Ensures modal is shown */
    flex-direction: column;
    align-items: center;
    justify-content: center;
}

.modal h3 {
    margin-bottom: 1rem;
}

/* Input Fields */
.input-field {
    width: 100%;
    padding: 10px;
    margin: 8px 0;
    border: 1px solid #ccc;
    border-radius: 5px;
}

/* Modal Buttons */
.modal-buttons {
    display: flex;
    justify-content: space-between;
    margin-top: 1rem;
    gap: 30px;
}

.btn-submit {
    background-color: #007bff;
    color: white;
    padding: 10px 20px;
    border: none;
    cursor: pointer;
    border-radius: 5px;
}

.btn-cancel {
    background-color: #dc3545;
    color: white;
    padding: 10px 20px;
    border: none;
    cursor: pointer;
    border-radius: 5px;
}

.btn-submit:hover {
    background-color: #0056b3;
    color: white;
}

.btn-cancel:hover {
    background-color: #c82333;
    color: white;

}
</style>
