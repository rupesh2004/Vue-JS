<script>
import axios from "axios";
import { onMounted, ref } from "vue";
import {useRouter} from "vue-router"
export default {
    name: "Login",
    setup() {
        const email = ref("");
        const password = ref("");
        const router = useRouter();

        const login = async() => {
            if (email.value === "admin@gmail.com" || password.value === "admin") {
                router.push({ name: "Admin" });
                return;
            }else{
                try {
                const user = {
                    email: email.value,
                    password: password.value
                }
                const response = await axios.post("http://localhost:3000/login",user)
                console.log(response.data)
                if(response.status === 200){
                    sessionStorage.setItem("token",response.data.token)
                    email.value=null
                    password.value=null
                    router.push({ name: "Dashboard" });
                    alert("Login Successful")
                }
                } catch (error) {
                    if (error.response){
                        if (error.response.status === 401){
                            alert("Invalid Email or Password")
                        }else if (error.response.status === 404){
                            alert("User Not Found")
                        }else{
                            alert("Login Failed")
                        }
                    }else{
                        alert("Network error")
                    } 
                }
            }
            
        };
        onMounted(()=>{
            const token = sessionStorage.getItem("token")
            if (token){
                router.push({ name: "Dashboard" });
            }
        })
        const gotoSignUp = () => {
            router.push({ name: "Register" });
        };

        return { email, password, login, gotoSignUp };
    }
};
</script>

<template>
    <div class="login-wrapper">
        <div v-if="!isSignUp" class="login-container">
            <h1>Login</h1>
            <form @submit.prevent="login">
                <div class="input-group">
                    <label for="email">Email:</label>
                    <input type="email" id="email" v-model="email" required>
                </div>
                <div class="input-group">
                    <label for="password">Password:</label>
                    <input type="password" id="password" v-model="password" required>
                </div>
                <button type="submit">Login</button>
            </form>
            <p class="signup-option">Don't have an account? <a href="#" @click.prevent="gotoSignUp">Sign up</a></p>
        </div>
    </div>
</template>

<style scoped>
.login-wrapper {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    background-color: #f4f4f4;
}
.login-container {
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
.signup-option {
    margin-top: 10px;
}
.signup-option a {
    color: #007bff;
    text-decoration: none;
    cursor: pointer;
}
.signup-option a:hover {
    text-decoration: underline;
}
</style>