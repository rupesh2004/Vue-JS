<script setup>
import { ref } from "vue";
import axios from "axios";
import { useRouter } from "vue-router";

const router = useRouter();

const examForm = ref({
    examName: "",
    examDate: "",
    applicationCloseDate: "",
    examDuration: "",
    examType: "",
});

const submitExam = async () => {
    const token = sessionStorage.getItem("jwt_token");
    if (!token) {
        alert("Unauthorized! Please log in.");
        router.push({ name: "Login" });
        return;
    }

    // Convert dates to ISO format
    if (examForm.value.examDate) {
        examForm.value.examDate = new Date(examForm.value.examDate).toISOString();
    }
    if (examForm.value.applicationCloseDate) {
        examForm.value.applicationCloseDate = new Date(examForm.value.applicationCloseDate).toISOString();
    }

    console.log("Sending payload:", JSON.stringify(examForm.value));

    try {
        const response = await axios.post("http://localhost:3000/createExam", examForm.value, {
            headers: { Authorization: `Bearer ${token}` },
        });
        alert("Exam created successfully!");
        router.push({ name: "AdminDashboard" });
    } catch (error) {
        console.error("Error creating exam:", error.response?.data || error.message);
        alert(error.response?.data?.error || "Failed to create exam.");
    }
};

const goBack = () => {
    router.push({ name: "AdminDashboard" });
};
</script>

<template>
    <nav class="navbar">
        <h1>Exam Portal</h1>
        <button class="back-button" style="background-color: #D4E0EDFF;color : black" @click="goBack">Back</button>
    </nav>

    <div class="container">
        <main class="main-content">
            <h2>Create New Exam</h2>
            <form @submit.prevent="submitExam" class="exam-form">
                <div class="form-group">
                    <label>Exam Name</label>
                    <input v-model="examForm.examName" type="text" required />
                </div>
                <div class="form-group">
                    <label>Exam Date</label>
                    <input v-model="examForm.examDate" type="date" :min="new Date().toISOString().split('T')[0]" required />
                </div>
                <div class="form-group">
                    <label>Application Close Date</label>
                    <input
                        v-model="examForm.applicationCloseDate"
                        type="date"
                        :min="new Date().toISOString().split('T')[0]"
                        required
                    />
                </div>
                <div class="form-group">
                    <label>Exam Duration</label>
                    <select v-model="examForm.examDuration" required>
                        <option value="">Select Duration</option>
                        <option value="1hr">1 Hour</option>
                        <option value="2hr">2 Hours</option>
                        <option value="3hr">3 Hours</option>
                    </select>
                </div>
                <div class="form-group">
                    <label>Exam Type</label>
                    <select v-model="examForm.examType" required>
                        <option value="">Select Type</option>
                        <option value="MCQ">MCQ</option>
                        <option value="Subjective">Subjective</option>
                        <option value="Practical">Practical</option>
                    </select>
                </div>
                <button type="submit" style = "background-color: #007bff;" class="submit-button">Create Exam</button>
            </form>
        </main>
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
}

.navbar {
    background: #007bff;
    color: white;
    padding: 15px 20px;
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.navbar h1 {
    margin: 0;
    font-size: 1.5rem;
}

.back-button {
    /* background-color: ; */
    background-color: #007bff;
    border: none;
    padding: 8px 15px;
    font-size: 1rem;
    border-radius: 5px;
    cursor: pointer;
}

.back-button:hover {
    background: #e0e0e0;
}

.main-content {
    flex-grow: 1;
    padding: 20px;
    max-width: 500px;
    width: 100%;
    margin: 20px auto;
    background: white;
    box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1);
    border-radius: 8px;
}

.exam-form {
    display: flex;
    flex-direction: column;
}

.form-group {
    margin-bottom: 15px;
}

.form-group label {
    font-weight: bold;
    margin-bottom: 5px;
    display: block;
}

.form-group input,
.form-group select {
    width: 100%;
    padding: 10px;
    border: 1px solid #ccc;
    border-radius: 5px;
    font-size: 1rem;
}

.form-group input:focus,
.form-group select:focus {
    outline: none;
    border-color: #007bff;
}

.submit-button {
    background-color: #007bff;
    color: white;
    padding: 10px;
    font-size: 1rem;
    border: none;
    border-radius: 5px;
    cursor: pointer;
}

.submit-button:hover {
    background-color: #7FB5EFFF;
}

.footer {
    background: #333;
    color: white;
    text-align: center;
    padding: 10px 0;
}
</style>
