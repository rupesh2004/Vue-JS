<script>
    export default {
        name:"RefDemo",
        methods:{
            submitForm(){
                const userName = this.$refs.userName.value;
                const email  = this.$refs.email.value
                alert(userName+email)
            }
        }
    }
</script>

<template>
    <div>
        <input type="text" ref="userName">
        <input type="email" ref="email" id="">
        <button @click="submitForm">Submit</button>
    </div>
</template>