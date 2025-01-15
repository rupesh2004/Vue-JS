<script>
export default {
    name: "ForDemo",
    data() {
        return {
            users: [{
                    id: 1,
                    name: 'John',
                    age: 25
                },
                {
                    id: 2,
                    name: 'Jane',
                    age: 30
                },
                {
                    id: 3,
                    name: 'Bob',
                    age: 35
                }
            ]
        }
    },
}
</script>

<template>
<div>
    <h1>For Demo</h1>
    <ul>
        <li v-for="user in users" :key="user.id">
            {{ user.id }} {{ user.name }} {{ user.age }}
        </li>
    </ul>
</div>
</template>
