<script>
import { ref, onMounted } from "vue";
import axios from "axios";
import { useRouter } from "vue-router";

export default {
  name: "ExamRegistration",
  setup() {
    const router = useRouter();
    const exams = ref([]);
    const appliedExams = ref(new Set()); // Stores applied exam IDs
    const loading = ref(true);
    const error = ref("");

    const fetchExams = async () => {
      try {
        const token = sessionStorage.getItem("jwt_token");
        const response = await axios.get("http://localhost:3000/getExams", {
          headers: { Authorization: `Bearer ${token}` },
        });
        exams.value = response.data;

        // Fetch applied exams after fetching the exam list
        await fetchAppliedExams();
      } catch (err) {
        error.value = "Failed to load exams.";
      } finally {
        loading.value = false;
      }
    };

    const fetchAppliedExams = async () => {
      try {
        const token = sessionStorage.getItem("jwt_token");
        const userEmail = JSON.parse(atob(token.split(".")[1])).email; // Decode JWT to get email
        const response = await axios.get(`http://localhost:3000/getUserApplications/${userEmail}`, {
          headers: { Authorization: `Bearer ${token}` },
        });

        appliedExams.value = new Set(response.data.map(app => app.examId));
      } catch (err) {
        console.error("Error fetching applied exams:", err);
      }
    };

    const applyForExam = async (examId) => {
      if (!confirm(`Are you sure you want to apply for this exam`)) return;
      try {
        const token = sessionStorage.getItem("jwt_token");
        await axios.post(`http://localhost:3000/applyExams/${examId}`, {}, {
          headers: { Authorization: `Bearer ${token}` },
        });
        alert("Successfully applied for the exam!");
        appliedExams.value.add(examId);
      } catch (err) {
        alert("Failed to apply for the exam.");
        console.error(err);
      }
    };

    const isApplicationClosed = (closeDate) => {
      return new Date(closeDate) < new Date();
    };

    const hasApplied = (examId) => {
      return appliedExams.value.has(examId);
    };

    const goBack = () => {
      router.push({ name: "UserDashboard" });
    };

    const formatDate = (date) => {
      return new Date(date).toLocaleDateString();
    };

    onMounted(fetchExams);

    return {
      exams,
      appliedExams,
      loading,
      error,
      applyForExam,
      formatDate,
      isApplicationClosed,
      hasApplied,
      goBack
    };
  },
};
</script>
<template>
  <div>
      <!-- Navbar -->
      <nav class="navbar">
          <h3><i class="fas fa-file-alt"></i> Exam Registration</h3>
          <button class="back-button" @click="goBack">
              <i class="fas fa-arrow-left"></i> Back
          </button>
      </nav>

      <!-- Content -->
      <div class="container">
          <h2><i class="fas fa-list-alt"></i> Available Exams</h2>
          <div v-if="loading" class="loading"><i class="fas fa-spinner fa-spin"></i> Loading exams...</div>
          <div v-if="error" class="error"><i class="fas fa-exclamation-triangle"></i> {{ error }}</div>

          <table v-if="exams.length" class="exam-table">
              <thead>
                  <tr>
                      <th><i class="fas fa-book"></i> Exam Name</th>
                      <th><i class="fas fa-calendar-alt"></i> Exam Date</th>
                      <th><i class="fas fa-clock"></i> Duration</th>
                      <th><i class="fas fa-tags"></i> Type</th>
                      <th><i class="fas fa-hourglass-end"></i> Last Date</th>
                      <th><i class="fas fa-tasks"></i> Action</th>
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
                          <button 
                              v-if="!isApplicationClosed(exam.applicationCloseDate)"
                              :disabled="hasApplied(exam.examId)"
                              :class="{ 'disabled-btn': hasApplied(exam.examId) }"
                              @click="applyForExam(exam.examId)"
                              class="apply-btn"
                          >
                              <i :class="hasApplied(exam.examId) ? 'fas fa-check-circle' : 'fas fa-edit'"></i>
                              {{ hasApplied(exam.examId) ? "Applied" : "Apply" }}
                          </button>
                      </td>
                  </tr>
              </tbody>
          </table>

          <div v-if="!loading && exams.length === 0" class="no-exams">
              <i class="fas fa-info-circle"></i> No exams available.
          </div>
      </div>

      <!-- Footer -->
      <footer class="footer">
          <p><i class="fas fa-copyright"></i> 2025 Exam Portal. All Rights Reserved.</p>
      </footer>
  </div>
</template>

<style scoped>
/* General Styles */
body {
  font-family: Arial, sans-serif;
  margin: 0;
  padding: 0;
}
.disabled-btn {
  background-color: #ccc !important;
  cursor: not-allowed !important;
}

/* Navbar */
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
  display: flex;
  align-items: center;
  gap: 8px;
}

.back-button:hover {
  background: #e0e0e0;
}

/* Content Container */
.container {
  max-width: 90%;
  margin: 20px auto;
  padding: 20px;
  background: #f5f5f5;
  text-align: center;
}

/* Table Styling */
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

/* Apply Button */
.apply-btn {
  background-color: #28a745 !important;
  color: white !important;
  padding: 8px 15px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  font-size: 14px;
  transition: 0.3s;
  display: flex;
  align-items: center;
  gap: 5px;
}

.apply-btn:hover {
  background-color: #218838;
}

/* Footer */
.footer {
  background: #333;
  color: white;
  text-align: center;
  padding: 10px 0;
  width: 100%;
  position: fixed;
  bottom: 0;
  left: 0;
}

/* Responsive */
@media (max-width: 768px) {
  .exam-table {
      font-size: 14px;
  }
}
</style>
