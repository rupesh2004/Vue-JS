<script>
import LoginForm from './components/loginForm.vue';
import SignupForm from './components/signupForm.vue';

    export default {
        name : "App",
        components:{
            SignupForm,
            LoginForm
        },
        

    }
</script>

<template>
    
    <LoginForm/>
    <SignupForm/>
</template>