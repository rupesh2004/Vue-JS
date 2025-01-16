<script>
    import axios from 'axios'
    export default {
        name : "GetAPI",

        data(){
            return {
                data :[]
            };
        },
        async mounted(){
            let result = await axios.get("https://reqres.in/api/users?page=1")
            this.data = result.data.data
        } 
    }
</script>

<template>
    <h1>API Data</h1>
    <p>{{ data }}</p>
    <ul>
        <li v-for="item in data" :key="item.id">
            {{ item.id }} {{ item.email }} {{ item.first_name }} {{ item.last_name }} {{ item.avatar }}
        </li>
    </ul>
</template>