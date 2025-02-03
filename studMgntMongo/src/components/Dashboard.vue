
<script>
import {useRouter} from 'vue-router';

    export default {
        name : "Dashboard",
        setup(){
            const router = useRouter();
            const addStudent = () => {
                router.push({name : 'AddStudent'});
            }
            const fetchAllStudent = () => {
                router.push({name : 'FetchAllStudent'});
            }
            
            return {addStudent,fetchAllStudent};
        }
    }
</script>
<template>
    <div class="container">
      <h1>Student Management System</h1>
      <div class="sub-container">
        <div class="card text-center">
          <img src="https://cdn4.iconfinder.com/data/icons/social-messaging-ui-color-and-shapes-7/177800/310-512.png" class="card-img-top student-logo" alt="Student Logo">
          <div class="card-body">
            <h5 class="card-title">Add Students</h5>
            <button class="btn btn-primary add-btn" @click="addStudent"  >
              <i class="fas fa-user-plus"></i> Add Student
            </button>
          </div>
        </div>
        <div class="card text-center">
          <img src="https://cdn-icons-png.flaticon.com/512/3135/3135755.png" class="card-img-top student-logo" alt="Student Logo">
          <div class="card-body">
            <h5 class="card-title">View Students</h5>
            <button class="btn btn-primary add-btn" @click="fetchAllStudent" >
              <i class="fas fa-user-plus"></i> View  Student
            </button>
          </div>
        </div>
      </div>
    </div>
    
  </template>
  
  <style scoped>
  .container {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    gap: 20px;
    flex-direction: column;
  }
  .sub-container {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 40px;
  }
  .card {
    width: 10rem;
    border-radius: 12px;
    box-shadow: 0px 6px 15px rgba(0, 0, 0, 0.15);
    transition: transform 0.3s ease-in-out, box-shadow 0.3s ease-in-out;
    background: linear-gradient(to bottom, #ffffff, #f8f9fa);
    padding: 20px;
    
    border: 1px solid blue;
  }
  
  .card:hover {
    transform: scale(1.05);
    box-shadow: 0px 8px 20px rgba(0, 0, 0, 0.25);
  }
  
  .student-logo {
    width: 80px;
    margin-left: 30px;
    display: block;
    border-radius: 50%;
    background-color: #eef2ff;
    padding: 10px;
  }
  
  .card-title {
    font-size: 1.2rem;
    font-weight: 600;
    margin-left : 30px;
    color: #333;
    margin-bottom: 15px;
  }
  
  /* Add Button */
  .add-btn {
    background: #007bff;
    border: none;
    color:white;
    font-size: 0.9rem;
    font-weight: 500;
    padding: 10px 20px;
    border-radius: 8px;
    margin-left:20px;
    cursor:pointer;
    transition: background 0.3s ease-in-out;
  }
  
  .add-btn i {
    margin-right: 8px;
  }
  
  .add-btn:hover {
    background: #0056b3;

  }
  </style>
  
  