<script>
import { computed, ref } from 'vue';

const setup = ()=>{
        const num1 = ref(0)
        const num2  = ref(0)
    
        const result = computed(()=>{
            return num1.value+num2.value
        })
    
        return {
            num1,num2,
            result
        }
    }
    export default{
        name:"ComputedDemo",
        setup
    }
</script>

<template>
    <br />
    <br />
    <input type="number" name="" id="" v-model.lazy="num1"><br /><br />
    <input type="number" name="" id="" v-model.lazy="num2">
    <p>{{ result }}</p>
</template>