<script>
    export default {
        name:"Child",
        props:{
            name: String,
            users : Object,
            getMobile:Function
        },
        data(){
            return {
                h1Tag: "<h1>Hello from rupesh</h1>",
                h2Tag : "<h2>Hello from rupesh</h2>",
               
            }
        },
        
    }
</script>

<template>
    <h1>child component</h1>
    <div v-html="h1Tag">
    </div>
    <div v-html="h2Tag"></div>
    <h3>Name : {{ name }}</h3>
    <h3>Users Data</h3>
    <ul>
        <li v-for="user in users" :key="user.user">
            {{ user.user }} {{ user.age }} {{ user.address }}
        </li>
    </ul>

    <button @click="getMobile()">Click to get number</button>

    
</template>