<script>
export default {
    name : "SetData",
    data(){
        return {
            email : null,
            work : null
        }
    },
    methods :{
        addWork(){
            alert("hello")
        }
    }
}
</script>

<template>
    <input type="email" name="" id="" placeholder="Email" v-model="email"><br /><br />
    <input type="text" name="" id="" placeholder="work" v-model="work"><br /><br />
    <button @click="addWork">Submit</button>
    
</template>