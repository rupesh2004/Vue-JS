<script>
    export default {
        name:"BeforeUpdateAndUpdated",
        data(){
            return{
                count:0
            }
        },
        beforeUpdate(){
            console.log("Before update",this.$refs.count.textContent)
        },
        updated(){
            console.log("update",this.$refs.count.textContent)
        }
    }
</script>

<template>
    <h1 ref="count">{{ count }}</h1>
    <button @click="count+=1">+</button>
</template>