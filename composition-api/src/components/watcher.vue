<script>
import { computed, reactive, ref, watch } from 'vue';

    export default{
        name:"Watcher",
        setup(){
            const num1 = ref()
            const num2  = ref()

            const formData = reactive({
                num1,num2
            })

            watch(()=>{
                return {...formData}
            },(newVal,oldVal)=>{
                console.log(newVal,oldVal)
            })

            return {
                num1,num2,
            }
        }
    }
</script>

<template>
    <br />
    <br />
    <input type="number" name="" id="" v-model.lazy="num1"><br /><br />
    <input type="number" name="" id="" v-model.lazy="num2">
</template>