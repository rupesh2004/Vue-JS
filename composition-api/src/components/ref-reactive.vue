<script>
    import {ref,reactive} from 'vue'
    export default {
        name : "RefReactive",
        setup(){
            const name = ref("Rupesh")
            setTimeout(()=>name.value = "abc",3000)

            const userData = reactive({
                firstName : "Rupesh",
                lastName : "Bhosale",
                age : 21
            })
            return {
                name,
                userData
            }
        }
    }
</script>

<template>
    <div>
        <h1>Hello from {{ name }}</h1>
        <p>User Data : {{ userData }}</p>
    </div>
</template>