<script>
import { computed, ref } from 'vue';

    export default{
        name:"ComputedDemo",
        setup(){
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
    }
</script>

<template>
    <br />
    <br />
    <input type="number" name="" id="" v-model.lazy="num1"><br /><br />
    <input type="number" name="" id="" v-model.lazy="num2">
    <p>{{ result }}</p>
</template>