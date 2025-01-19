<script>
import { ref } from 'vue';

    export default {
        name : "Methods",
        setup(){
            let count = ref(0)
            function increment(){
                count.value+=1
            }
            function decrement(){
                if(count.value>0){
                    count.value-=1
                }else {
                    alert("Cannot decrement")
                }
            }
            return {
                count,
                increment,
                decrement
            }
        }
    }
</script>

<template>
    <p>{{ count }}</p>
    <button @click="increment">Increment</button>
    <button @click="decrement">Decrement</button>
</template>