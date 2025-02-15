<script>
import {
    ref,
    onMounted
} from "vue";
import axios from "axios";
import {
    useRouter
} from "vue-router";

export default {
    name: "ExamRegistration",
    setup() {
        const router = useRouter();
        const exams = ref([]);
        const loading = ref(true);
        const error = ref("");

        const fetchExams = async () => {
            try {
                const token = sessionStorage.getItem("jwt_token");
                const response = await axios.get("http://localhost:3000/getExams", {
                    headers: {
                        Authorization: `Bearer ${token}`
                    },
                });
                exams.value = response.data;
            } catch (err) {
                error.value = "Failed to load exams.";
            } finally {
                loading.value = false;
            }
        };

        const applyForExam = async (examId) => {
            try {
                const token = sessionStorage.getItem("jwt_token");
                await axios.post(`http://localhost:3000/applyExams/${examId}`, {}, {
                    headers: {
                        Authorization: `Bearer ${token}`
                    },
                });
                alert("Successfully applied for the exam!");
            } catch (err) {
                alert("Failed to apply for the exam.");
                console.error(err);
            }
        };

        const goBack = () => {
            router.push({
                name: "UserDashboard"
            })
        };

        const formatDate = (date) => {
            return new Date(date).toLocaleDateString();
        };

        onMounted(fetchExams);

        return {
            exams,
            loading,
            error,
            applyForExam,
            formatDate,
            goBack
        };
    },
};
</script>

<template>
<div>
    <nav class="navbar">
        <h1>Exam Registration</h1>
        <button class="back-button" @click="goBack">Back</button>
    </nav>

    <div class="container">
        <h2>Available Exams</h2>
        <div v-if="loading" class="loading">Loading exams...</div>
        <div v-if="error" class="error">{{ error }}</div>

        <table v-if="exams.length" class="exam-table">
            <thead>
                <tr>
                    <th>Exam Name</th>
                    <th>Exam Date</th>
                    <th>Duration</th>
                    <th>Type</th>
                    <th>Last Date</th>
                    <th>Action</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="exam in exams" :key="exam.examId">
                    <td>{{ exam.examName }}</td>
                    <td>{{ formatDate(exam.examDate) }}</td>
                    <td>{{ exam.examDuration }}</td>
                    <td>{{ exam.examType }}</td>
                    <td>{{ formatDate(exam.applicationCloseDate) }}</td>
                    <td>
                        <button @click="applyForExam(exam.examId)" class="apply-btn" :disabled="new Date(exam.applicationCloseDate) < new Date()" :class="{'disabled-btn': new Date(exam.applicationCloseDate) < new Date()}">
                            Apply
                        </button>
                      </td>

                </tr>
            </tbody>
        </table>

        <div v-if="!loading && exams.length === 0" class="no-exams">No exams available.</div>
    </div>

    <footer class="footer">
        <p>© 2025 Exam Portal. All Rights Reserved.</p>
    </footer>
</div>
</template>

<style>
body {
    font-family: Arial, sans-serif;
    background-color: #f4f4f4;
    margin: 0;
    padding: 0;
}
.disabled-btn {
    background-color: #ccc !important;
    cursor: not-allowed !important;
}


.navbar {
    background: #007bff;
    color: white;
    padding: 15px 20px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    text-align: center;
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

.container {
    max-width: 800px;
    margin: 20px auto;
    padding: 20px;
    background: white;
    box-shadow: 0px 4px 8px rgba(0, 0, 0, 0.1);
    border-radius: 8px;
    text-align: center;
}

.exam-table {
    width: 100%;
    border-collapse: collapse;
    margin-top: 20px;
}

.exam-table th,
.exam-table td {
    border: 1px solid #ddd;
    padding: 12px;
    text-align: center;
}

.exam-table th {
    background: #007bff;
    color: white;
}

.apply-btn {
    background-color: #28a745 !important;
    color: white !important;
    padding: 8px 15px;
    border: none;
    border-radius: 5px;
    cursor: pointer;
    font-size: 14px;
    transition: 0.3s;
}

.apply-btn:hover {
    background-color: #218838;
}

.footer {
    background: #333;
    color: white;
    text-align: center;
    padding: 10px 0;
    width: 100%;
    margin-top: 20px;
}
</style>
