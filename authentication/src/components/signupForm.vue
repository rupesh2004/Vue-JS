<script>
import { useRouter } from 'vue-router';
import {useUserAuth} from '../stores/auth.js'
import { ref } from 'vue';
export default {
  name: "SignUpForm",
  setup(){
    const router = useRouter()
    const userAuth = useUserAuth();

    const username  = ref('')
    const email = ref('')
    const password = ref('')
    const hobbies = ref([])
    const gender = ref(null)

    const registerUser = ()=>{
      userAuth.userRegistration(
        username.value,
        password.value,
        email.value,
        hobbies.value,
        gender.value,

      )
      alert("User Registered")
      router.push({name : "LoginForm"})
    }
    function gotoLogin(){
      router.push({name : "LoginForm"})
    }
    return {
      userAuth,username,email,password,hobbies,gender,registerUser,gotoLogin
    }

  }
};
</script>

<template>
  <div class="signup-container">
    <div class="signup-form">
      <h2>Sign Up</h2>
      <form @submit.prevent="registerUser">
        <div class="form-group">
          <b>Username</b>
          <input
            type="text"
            id="username"
            v-model="username"
            placeholder="Enter your username"
            required
          />
        </div>
        <div class="form-group">
          <b>Email</b>
          <input
            type="email"
            id="email"
            v-model="email"
            placeholder="Enter your email"
            required
          />
        </div>
        <div class="form-group">
          <b>Password</b>
          <input
            type="password"
            id="password"
            v-model="password"
            placeholder="Enter your password"
            required
          />
        </div>
        <div class="form-group">
          <b>Hobbies</b>
          <div class="checkbox-group">
            <div>
              <input type="checkbox" id="gaming" value="Gaming" v-model="hobbies" />
              <label for="gaming">Gaming</label>
            </div>
            <div>
              <input type="checkbox" id="reading" value="Reading" v-model="hobbies" />
              <label for="reading">Reading</label>
            </div>
            <div>
              <input type="checkbox" id="traveling" value="Traveling" v-model="hobbies" />
              <label for="traveling">Traveling</label>
            </div>
          </div>
        </div>
        <div class="form-group">
          <b>Gender</b>
          <div class="radio-group">
            <label>
              <input
                type="radio"
                name="gender"
                value="Male"
                v-model="gender"
              />
              Male
            </label>
            <label>
              <input
                type="radio"
                name="gender"
                value="Female"
                v-model="gender"
              />
              Female
            </label>
            <label>
              <input
                type="radio"
                name="gender"
                value="Other"
                v-model="gender"
              />
              Other
            </label>
          </div>
        </div>
        <button type="submit" class="btn-submit">Sign Up</button>
        <p>Already have an account? <Span style="color: blue;cursor:pointer" @click="gotoLogin">Login</Span></p>
      </form>

    </div>
    <!-- display all data -->
    <div class="userData" v-if="visible">
    <h3>Username : {{ username }}</h3>
    <h3>Email : {{ email }}</h3>
    <h3>Password : {{ password }}</h3>
    <h3>Hobbies</h3>
    <ul>
      <li v-for="(hobby,index) in hobbies" :key="index">
      {{ hobby }}
      </li>
    </ul>
   
    <h3>Gender : {{ gender }}</h3>
    </div>
    <div v-else>
    </div>
  </div>
</template>

<style scoped>
.signup-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background-color: #f4f4f4;
}

.signup-form {
  background-color: #fff;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  width: 300px;
  text-align: center;
}

h2 {
  margin-bottom: 20px;
  font-size: 24px;
  color: #333;
}

.form-group {
  margin-bottom: 15px;
  text-align: left;
}

label {
  font-size: 14px;
  margin-left: 5px;
  color: #555;
  cursor: pointer;
}

.checkbox-group,
.radio-group {
  display: flex;
  gap: 15px; /* Add space between options */
  align-items: center;
}

input[type="text"],
input[type="email"],
input[type="password"] {
  width: 100%;
  padding: 8px;
  font-size: 14px;
  border: 1px solid #ddd;
  border-radius: 4px;
}

input:focus {
  outline: none;
  border-color: #3498db;
}

input[type="checkbox"],
input[type="radio"] {
  margin: 0;
  width: auto;
  height: auto;
}

.btn-submit {
  width: 100%;
  padding: 10px;
  background-color: #3498db;
  color: #fff;
  border: none;
  border-radius: 4px;
  font-size: 16px;
  cursor: pointer;
}

.btn-submit:hover {
  background-color: #2980b9;
}
</style>
