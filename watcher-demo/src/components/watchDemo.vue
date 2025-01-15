<script>
    export default {
        name : "WatchDemo",
        data(){
            return {
                count : 0
            }
        },
        methods : {
            increment(){
                this.count += 1
            },
            decrement(){
                this.count-=1
            }
        },
        watch:{
            count(val,prev){
                alert(this.count)
                alert(prev)
            }
        }
    }
</script>

<template>
    <div>
        <h1>{{ count }}</h1>
        <button @click="increment">+</button>
        <button @click="decrement">-</button>
    </div>
</template>