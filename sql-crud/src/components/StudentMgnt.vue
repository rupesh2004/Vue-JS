<script>
import {
    ref,
    reactive,
    onMounted
} from "vue";
import axios from 'axios';

export default {
    setup() {
        const user = reactive({
            id: null, // Track ID for update
            name: "",
            mobile: "",
            prn: "",
            dob: "",
        });

        const users = ref([]);
        const isEditing = ref(false); // Track if updating

        const formatDate = (dateString) => {
            const parts = dateString.split("-");
            return `${parts[2]}-${parts[1]}-${parts[0]}`;
        };

        const addUser = async () => {
            try {
                const payload = {
                    name: user.name,
                    mobileNo: user.mobile,
                    prn: parseInt(user.prn, 10),
                    birthDate: formatDate(user.dob),
                };

                if (isEditing.value) {
                    // Update student if editing
                    await axios.put(`http://localhost:3000/updateStudent/${user.id}`, payload);
                    alert("Student updated successfully");
                } else {
                    // Add new student
                    await axios.post("http://localhost:3000/addStudent", payload);
                    alert("Student created successfully");
                }

                // Reset form
                clearForm();
                FetchAllStudents();
            } catch (err) {
                console.log(err);
            }
        };

        const FetchAllStudents = async () => {
            try {
                const response = await axios.get("http://localhost:3000/getStudents");
                users.value = response.data; // Store fetched students
            } catch (err) {
                console.log(err);
            }
        };

        const deleteStudent = async (id) => {
            if (confirm("Are you sure you want to delete this student?")) {
                try {
                    await axios.delete(`http://localhost:3000/deleteStudent/${id}`);
                    alert("Student deleted successfully");
                    FetchAllStudents();
                } catch (err) {
                    console.log(err);
                }
            }
        };

        const editStudent = (student) => {
            user.id = student._id; // Store MongoDB ID for updating
            user.name = student.name;
            user.mobile = student.mobileNo;
            user.prn = student.prn;
            user.dob = student.birthDate.split("-").reverse().join("-"); // Convert back to input format

            isEditing.value = true;
        };

        const clearForm = () => {
            Object.assign(user, {
                id: null,
                name: "",
                mobile: "",
                prn: "",
                dob: ""
            });
            isEditing.value = false;
        };

        onMounted(FetchAllStudents);

        return {
            user,
            users,
            addUser,
            deleteStudent,
            editStudent,
            isEditing
        };
    },
};
</script>

<template>
<div class="container">
    <h2 class="title">{{ isEditing ? "Update Student" : "Add Student" }}</h2>

    <!-- Form Card -->
    <div class="form-card">
        <form @submit.prevent="addUser">
            <div class="form-group">
                <label>Name:</label>
                <input type="text" v-model="user.name" required />
            </div>

            <div class="form-group">
                <label>Mobile No:</label>
                <input type="tel" v-model="user.mobile" required />
            </div>

            <div class="form-group">
                <label>PRN:</label>
                <input type="text" v-model="user.prn" required />
            </div>

            <div class="form-group">
                <label>Date of Birth:</label>
                <input type="date" v-model="user.dob" required />
            </div>

            <button type="submit">{{ isEditing ? "Update" : "Submit" }}</button>
        </form>
    </div>

    <!-- Table -->
    <h3 class="title">User Details</h3>
    <div class="table-container" v-if="users.length">
        <table>
            <thead>
                <tr>
                    <th>Name</th>
                    <th>Mobile No</th>
                    <th>PRN</th>
                    <th>DOB</th>
                    <th>Actions</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="(item, index) in users" :key="index">
                    <td>{{ item.name }}</td>
                    <td>{{ item.mobileNo }}</td>
                    <td>{{ item.prn }}</td>
                    <td>{{ item.birthDate }}</td>
                    <td>
                        <button class="edit-btn" @click="editStudent(item)">Edit</button>
                        <button class="delete-btn" @click="deleteStudent(item.prn)">Delete</button>
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</div>
</template>

<style scoped>
/* General Styles */
.container {
    max-width: 600px;
    margin: auto;
    text-align: center;
    font-family: 'Arial', sans-serif;
}

/* Title */
.title {
    margin-bottom: 15px;
    font-size: 24px;
    color: #333;
}

/* Form Card */
.form-card {
    background: #fff;
    padding: 20px;
    border-radius: 10px;
    box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.5);
    margin-bottom: 20px;
}

/* Form Group */
.form-group {
    display: flex;
    flex-direction: column;
    text-align: left;
    margin-bottom: 10px;
}

label {
    font-weight: bold;
    margin-bottom: 5px;
    color: #555;
}

input {
    padding: 10px;
    border: 1px solid #ccc;
    border-radius: 5px;
    transition: 0.3s;
}

input:focus {
    border-color: #007bff;
    outline: none;
    box-shadow: 0px 0px 5px rgba(0, 123, 255, 0.5);
}

/* Submit Button */
button {
    padding: 10px;
    border: none;
    color: white;
    cursor: pointer;
    border-radius: 5px;
    font-size: 16px;
    transition: 0.3s;
    background-color:blue;
}

button:hover {
    opacity: 0.8;
}

/* Edit Button */
.edit-btn {
    background: #ffc107;
    margin-right: 5px;
}

/* Delete Button */
.delete-btn {
    background: #dc3545;
}

/* Table */
.table-container {
    overflow-x: auto;
}

table {
    width: 100%;
    border-collapse: collapse;
    margin-top: 20px;
}

th,
td {
    border: 1px solid #ccc;
    padding: 10px;
    text-align: left;
}

th {
    background: #007bff;
    color: white;
}

tbody tr:nth-child(odd) {
    background: #f9f9f9;
}

tbody tr:hover {
    background: #e0f0ff;
    transition: 0.3s;
}
</style>
