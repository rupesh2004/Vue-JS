<script>
import axios from "axios";
import { onMounted, ref } from "vue";
import {useRouter} from "vue-router";
export default {
  name: "FetchAllStudent",
  setup() {
    const router = useRouter();
    const students = ref([]);
    const currentStudent = ref(null);
    const isEditing = ref(false);

    const getAllStudentsData = async () => {
      try {
        const response = await axios.get("http://localhost:3000/");
        students.value = response.data.students;
      } catch (error) {
        console.error("Error fetching students:", error);
      }
    };

    const deleteStudent = async (studentID) => {
      try {
        const response = await axios.delete(`http://localhost:3000/deleteStudent/${studentID}`);
        if (response.status === 200) {
          getAllStudentsData();
          alert("Student deleted successfully");
        }
      } catch (error) {
        console.error("Error deleting student:", error);
      }
    };

    const updateStudent = (studentID) => {
      const student = students.value.find((s) => s.studentID === studentID);
      currentStudent.value = { ...student };
      isEditing.value = true;
    };

const submitUpdate = async () => {
  try {
    const updatedStudent = {
      fullName: currentStudent.value.fullName,
      email: currentStudent.value.email,
      phone: currentStudent.value.phone,
      dob: currentStudent.value.dob,
    };

    console.log("Sending update request:", JSON.stringify(updatedStudent, null, 2));
    

    const response = await axios.put(
      `http://localhost:3000/updateStudent/${currentStudent.value.studentID}`,
      updatedStudent,
      { headers: { "Content-Type": "application/json" } }
    );

    if (response.status === 200) {
      // Update the student directly in the list
      const updatedIndex = students.value.findIndex(s => s.studentID === currentStudent.value.studentID);
      students.value[updatedIndex] = {
        ...currentStudent.value,
      };

      // Reset form and stop editing
      currentStudent.value = null;
      isEditing.value = false;

      alert("Student updated successfully");
    }
  } catch (error) {
    console.error("❌ Error updating student:", error.response?.data || error);
  }
};




  const gotoDashboard = () => {
    router.push({ name: 'Dashboard' });
  }


    onMounted(() => {
      getAllStudentsData();
    });

    return {
      students,
      deleteStudent,
      updateStudent,
      currentStudent,
      isEditing,
      submitUpdate,
      gotoDashboard,

    };
  },
};
</script>

<template>
  <button class = "gotoDashboard" @click = "gotoDashboard" >Go to Dashboard</button>

  <div class="container">
    <!-- Update Form -->
    <div v-if="isEditing" class="update-form">
      <h2>Update Student</h2>
      <form @submit.prevent="submitUpdate">
        <label for="fullName">Full Name:</label>
        <input type="text" id="fullName" v-model="currentStudent.fullName" required />

        <label for="email">Email:</label>
        <input type="email" id="email" v-model="currentStudent.email" required />

        <label for="phone">Phone:</label>
        <input type="text" id="phone" v-model="currentStudent.phone" required />

        <label for="dob">Date of Birth:</label>
        <input type="date" id="dob" v-model="currentStudent.dob" required />


        <div class="btn-group">
          <button type="submit" class="update-btn">Update</button>
          <button type="button" @click="isEditing = false" class="cancel-btn">Cancel</button>
        </div>
      </form>
    </div>

    <!-- Students Table -->
    <table class="styled-table" v-else>
      <thead>
        <tr>
          <th>Student ID</th>
          <th>Full Name</th>
          <th>Email</th>
          <th>Phone</th>
          <th>Date of Birth</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="student in students" :key="student.studentID">
          <td>{{ student.studentID }}</td>
          <td>{{ student.fullName }}</td>
          <td>{{ student.email }}</td>
          <td>{{ student.phone }}</td>
          <td>{{ student.dob }}</td>
          <td class="button-group">
            <button class="update-btn" @click="updateStudent(student.studentID)">Update</button>
            <button class="delete-btn" @click="deleteStudent(student.studentID)">Delete</button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<style scoped>
.gotoDashboard {
  background-color : #007bff;
  color : white;
  border : none;
  border-radius:8px;
  padding  : 15px;
  cursor:pointer;
  font-size :bold; 
  margin : 10px;
}
.container {
  max-width: 90%;
  margin: 20px auto;
  text-align: center;
}

h1 {
  font-size: 24px;
  color: #333;
  margin-bottom: 20px;
}

.styled-table {
  width: 100%;
  border-collapse: collapse;
  background: #fff;
  box-shadow: 0px 2px 10px rgba(0, 0, 0, 0.1);
  border-radius: 8px;
  overflow: hidden;
}

.styled-table th, .styled-table td {
  padding: 12px;
  border-bottom: 1px solid #ddd;
  text-align: center;
}

.styled-table th {
  background-color: #007bff;
  color: white;
  font-weight: bold;
}

.styled-table tr:hover {
  background-color: #f1f1f1;
}

.button-group {
  display: flex;
  gap: 10px;
  justify-content: center;
}

.update-btn, .delete-btn {
  padding: 8px 12px;
  border: none;
  cursor: pointer;
  font-size: 14px;
  border-radius: 4px;
  transition: all 0.3s ease;
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

/* Update Form */
.update-form {
  width: 50%;
  margin: 30px auto;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0px 2px 10px rgba(0, 0, 0, 0.1);
  background-color: #f8f9fa;
}

.update-form label {
  display: block;
  margin-bottom: 8px;
  font-weight: bold;
}

.update-form input {
  width: 100%;
  padding: 10px;
  margin-bottom: 15px;
  border: 1px solid #ccc;
  border-radius: 4px;
}

.btn-group {
  display: flex;
  gap: 10px;
  justify-content: center;
}
</style>
