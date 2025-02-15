<script>
import { ref, onMounted } from "vue";
import axios from "axios";
import { useRouter } from "vue-router";

export default {
    name: "ViewExam",
    setup() {
        const router = useRouter();
        const exams = ref([]);
        const applications = ref([]); // Declare applications
        const selectedExamName = ref(""); // Store selected exam name
        const showApplications = ref(false); // Track if modal is open
        const loading = ref(true);
        const error = ref("");
        const editingExam = ref(null);

        const fetchExams = async () => {
            try {
                const token = sessionStorage.getItem("jwt_token");
                const response = await axios.get("http://localhost:3000/getExams", {
                    headers: { Authorization: `Bearer ${token}` },
                });
                exams.value = response.data;
            } catch (err) {
                error.value = "Failed to load exams.";
            } finally {
                loading.value = false;
            }
        };

        const deleteExam = async (id) => {
            if (!confirm("Are you sure you want to delete this exam?")) return;
            try {
                const token = sessionStorage.getItem("jwt_token");
                await axios.delete(`http://localhost:3000/deleteExam/${id}`, {
                    headers: { Authorization: `Bearer ${token}` },
                });
                exams.value = exams.value.filter(exam => exam._id !== id);
                alert("Exam deleted successfully!");
            } catch (err) {
                console.error("Delete error:", err.response?.data || err.message);
                alert("Failed to delete exam.");
            }
        };

        const editExam = (id) => {
            editingExam.value = id;
        };

        const updateExam = async (exam) => {
            try {
                const token = sessionStorage.getItem("jwt_token");
                await axios.put(`http://localhost:3000/updateExam/${exam.examId}`, {
                    examId: exam.examId,
                    examName: exam.examName,
                    examDate: new Date(exam.examDate).toISOString(),
                    examDuration: exam.examDuration,
                    examType: exam.examType,
                    applicationCloseDate: new Date(exam.applicationCloseDate).toISOString(),
                }, {
                    headers: { Authorization: `Bearer ${token}` },
                });
                alert("Exam updated successfully!");
                editingExam.value = null;
            } catch (err) {
                console.log("Failed to update exam.", err.response?.data || err.message);
            }
        };

        const viewApplications = async (examId, examName) => {
            try {
                const token = sessionStorage.getItem("jwt_token");
                const response = await axios.get(`http://localhost:3000/getApplications/${examId}`, {
                    headers: { Authorization: `Bearer ${token}` },
                });
                applications.value = response.data;
                selectedExamName.value = examName;
                showApplications.value = true;
                console.log(response.data);
            } catch (err) {
                console.error("Error fetching applications:", err.response?.data || err.message);
                alert("Failed to fetch applications.");
            }
        };

        const goBack = () => {
            router.push({ name: "AdminDashboard" });
        };

        const formatDate = (date) => new Date(date).toLocaleDateString();

        onMounted(fetchExams);

        return {
            exams,
            applications,
            selectedExamName,
            showApplications,
            loading,
            error,
            editingExam,
            fetchExams,
            deleteExam,
            editExam,
            updateExam,
            viewApplications,
            goBack,
            formatDate,
        };
    },
};
</script>

<template>
    <nav class="navbar">
        <h1>Exam Portal</h1>
        <button class="back-button" @click="goBack">Back</button>
    </nav>

    <div class="container">
        <main class="main-content">
            <h2>All Exams</h2>

            <div v-if="loading" class="loading">Loading exams...</div>
            <div v-if="error" class="error">{{ error }}</div>

            <table v-if="exams.length" class="exam-table">
                <thead>
                    <tr>
                        <th>Exam Name</th>
                        <th>Exam Date</th>
                        <th>Close Date</th>
                        <th>Duration</th>
                        <th>Type</th>
                        <th>Actions</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="exam in exams" :key="exam._id">
                        <td>
                            <input v-if="editingExam === exam._id" v-model="exam.examName" />
                            <span v-else>{{ exam.examName }}</span>
                        </td>
                        <td>
                            <input v-if="editingExam === exam._id" v-model="exam.examDate" type="date" />
                            <span v-else>{{ formatDate(exam.examDate) }}</span>
                        </td>
                        <td>
                            <input v-if="editingExam === exam._id" v-model="exam.applicationCloseDate" type="date" />
                            <span v-else>{{ formatDate(exam.applicationCloseDate) }}</span>
                        </td>
                        <td>
                            <select v-if="editingExam === exam._id" v-model="exam.examDuration">
                                <option value="1hr">1 Hour</option>
                                <option value="2hr">2 Hours</option>
                                <option value="3hr">3 Hours</option>
                            </select>
                            <span v-else>{{ exam.examDuration }}</span>
                        </td>
                        <td>
                            <select v-if="editingExam === exam._id" v-model="exam.examType">
                                <option value="MCQ">MCQ</option>
                                <option value="Subjective">Subjective</option>
                                <option value="Practical">Practical</option>
                            </select>
                            <span v-else>{{ exam.examType }}</span>
                        </td>
                        <td>
                            <button v-if="editingExam === exam._id" @click="updateExam(exam)">Save</button>
                            <button v-else @click="editExam(exam._id)">Edit</button>
                            <button @click="deleteExam(exam._id)">Delete</button>
                            <button @click="viewApplications(exam.examId, exam.examName)">View Applications</button>
                        </td>
                    </tr>
                </tbody>
            </table>

            <div v-if="!loading && exams.length === 0" class="no-exams">No exams found.</div>
        </main>
    </div>

    <!-- Applications Modal -->
    <div v-if="showApplications" class="modal">
        <div class="modal-content">
            <span class="close" @click="showApplications = false">&times;</span>
            <h2>Applications for {{ selectedExamName }}</h2>
            <table v-if="applications.length">
                <thead>
                    <tr>
                        <th>Exam Name</th>
                        <th>Email</th>
                        <th>Applied At</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="app in applications" :key="app.email">
                        <td>{{ app.examName }}</td>
                        <td>{{ app.email }}</td>
                        <td>{{ formatDate(app.appliedAt) }}</td>
                    </tr>
                </tbody>
            </table>
            <div v-if="applications.length === 0">No applications found.</div>
        </div>
    </div>
    <footer class="footer">
    <p>© 2025 Exam Portal. All Rights Reserved.</p>
</footer>
</template>

<style>
body {
    font-family: Arial, sans-serif;
    background-color: #f4f4f4;
    margin: 0;
    padding: 0;
}

.container {
    display: flex;
    flex-direction: column;
    min-height: 100vh;
    align-items: center;
}

/* Navbar */
.navbar {
    background: #007bff;
    color: white;
    padding: 15px 20px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    width: 100%;
    box-shadow: 0px 4px 8px rgba(0, 0, 0, 0.1);
}

.navbar h1 {
    margin: 0;
    font-size: 1.5rem;
}

.back-button {
    background-color: white !important;
    color: black !important;
    border: none;
    padding: 8px 15px;
    font-size: 1rem;
    border-radius: 5px;
    cursor: pointer;
    transition: 0.3s;
}

.back-button:hover {
    background: #e0e0e0;
}

/* Main Content */
.main-content {
    flex-grow: 1;
    padding: 20px;
    max-width: 90%;
    width: 800px;
    margin: 20px auto;
    background: white;
    box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.1);
    border-radius: 8px;
}

h2 {
    text-align: center;
    color: #333;
}

/* Table Styling */
.exam-table {
    width: 100%;
    border-collapse: collapse;
    margin-top: 20px;
    overflow-x: auto;
}

.exam-table th,
.exam-table td {
    border: 1px solid #ddd;
    padding: 12px;
    text-align: center;
    font-size: 14px;
}

.exam-table th {
    background: #007bff;
    color: white;
    font-weight: bold;
}

.exam-table tr:nth-child(even) {
    background-color: #f9f9f9;
}

.exam-table tr:hover {
    background-color: #f1f1f1;
}

/* Buttons */
.exam-table button {
    padding: 6px 12px;
    margin: 2px;
    border: none;
    cursor: pointer;
    border-radius: 5px;
    font-size: 14px;
    transition: 0.3s;
}

button:hover {
    opacity: 0.8;
}

button:first-child {
    background-color: #28a745;
    color: white;
}

button:nth-child(2) {
    background-color: red;
    color: white;
}

button:last-child {
    background-color: #3562DCFF;
    color: white;
}

/* Responsive Design */
@media (max-width: 768px) {
    .main-content {
        width: 95%;
        padding: 15px;
    }

    .exam-table th,
    .exam-table td {
        padding: 8px;
        font-size: 12px;
    }

    .exam-table {
        display: block;
        overflow-x: auto;
        white-space: nowrap;
    }

    .back-button {
        padding: 6px 12px;
        font-size: 14px;
    }
}

/* Status Messages */
.loading,
.error,
.no-exams {
    text-align: center;
    font-size: 1.2rem;
    margin-top: 20px;
    color: #333;
}

/* Footer */
.footer {
    background: #333;
    color: white;
    text-align: center;
    padding: 10px 0;
    width: 100%;
    margin-top: auto;
}

.modal {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: rgba(0, 0, 0, 0.5);
    display: flex;
    align-items: center;
    justify-content: center;
}

.modal-content {
    background: white;
    padding: 20px;
    border-radius: 10px;
    text-align: center;
    width: 50%;
}

.close {
    position: absolute;
    top: 10px;
    right: 20px;
    font-size: 1.5rem;
    cursor: pointer;
}

</style>
    