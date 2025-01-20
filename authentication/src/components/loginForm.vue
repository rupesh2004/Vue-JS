<script>
  import {useUserAuth} from '@/stores/auth.js'
  import {ref} from 'vue'
  import {useRouter} from 'vue-router'
export default {
  name: "LoginForm",
  setup(){
    const router = useRouter();
    const userAuth = useUserAuth()
    const username  = ref('')
    const password = ref('')

    function handleLogin(){
      const user = userAuth.userData.find((u)=>{
        return u.username === username.value && u.password === password.value
      })  
      if(user){
        alert("Login Successful")
      }else{
        alert("Invalid credential")
      }
    }
    function gotoRegister (){
      router.push({name : "SignUpForm"})
    }

    return {
      userAuth,username,password,handleLogin,gotoRegister
    }
  }

};
</script>

<template>
  <div class="login-container">
    <div class="login-form">
      <h2>Login</h2>
      <form @submit.prevent="handleLogin">
        <div class="form-group">
          <b>Username</b>
          <input
            type="text"
            id="username"
            placeholder="Enter your username"
            v-model="username"
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
        <button type="submit" class="btn-submit">Login</button>
        <p class="noAcc">
          Don't Have an Account? 
          <span class="register-link" @click="gotoRegister">Register Here</span>
        </p>
      </form>
    </div>
  </div>
</template>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background-color: #f4f4f4;
}

.login-form {
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
  display: block;
  font-size: 14px;
  margin-bottom: 5px;
  color: #555;
}

input {
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

.noAcc {
  margin-top: 15px;
  font-size: 14px;
  color: #666;
  text-align: center;
}

.register-link {
  color: #3498db;
  font-weight: bold;
  cursor: pointer;
  text-decoration: underline;
}

.register-link:hover {
  color: #2980b9;
  text-decoration: none;
}
</style>
