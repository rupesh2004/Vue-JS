<script>
import axios from "axios";
import { ref } from "vue";

import {useRouter} from "vue-router"
export default {
    name: "Register",
    setup() {
        const name = ref("");
        const email = ref("");
        const password = ref("");
        const router = useRouter();

        const register = async () => {
            const user = {
                name :name.value,
                email : email.value,
                password : password.value
            }
            console.log(user)
            try {
                const response = await axios.post("http://localhost:3000/addUser", user);
                if (response.status === 200){
                    alert("User created successfully");
                    // clear felids
                    name.value=null
                    email.value=null
                    password.value=null
                    router.push({name: "Login"})
                }
            } catch (error) {
                console.log(error)
            }
        };
        const gotoLogin =() =>{
            router.push({name :"Login"})
        }

        return { name, email, password, register,gotoLogin };
    }
};
</script>

<template>
    <div class="register-wrapper">
        <div class="register-container">
            <h1>Sign Up</h1>
            <form @submit.prevent="register">
                <div class="input-group">
                    <label for="name">Name:</label>
                    <input type="text" id="name" v-model="name" required>
                </div>
                <div class="input-group">
                    <label for="email">Email:</label>
                    <input type="email" id="email" v-model="email" required>
                </div>
                <div class="input-group">
                    <label for="password">Password:</label>
                    <input type="password" id="password" v-model="password" required>
                </div>
                <button type="submit">Sign Up</button>
            </form>
            <p class="login-option">Already have an account? <a href="#" @click = "gotoLogin">Login</a></p>
        </div>
    </div>
</template>

<style scoped>
.register-wrapper {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    background-color: #f4f4f4;
}
.register-container {
    width: 350px;
    padding: 30px;
    background: white;
    box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
    border-radius: 8px;
    text-align: center;
}
.input-group {
    margin-bottom: 15px;
    text-align: left;
}
label {
    display: block;
    font-weight: bold;
}
input {
    width: 100%;
    padding: 10px;
    border: 1px solid #ccc;
    border-radius: 4px;
}
button {
    width: 100%;
    padding: 10px;
    background-color: #007bff;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
}
button:hover {
    background-color: #0056b3;
}
.login-option {
    margin-top: 10px;
}
.login-option a {
    color: #007bff;
    text-decoration: none;
    cursor: pointer;
}
.login-option a:hover {
    text-decoration: underline;
}
</style>
