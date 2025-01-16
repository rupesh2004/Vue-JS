<script>
    export default {
        name: "BeforeUnmountAndUnMounted",
        beforeUnmount(){
            console.log("before unmount : ",this.$refs['tag'])
        },
        mounted(){
            console.log("mounted : ",this.$refs['tag'])
        }
    }

</script>

<template>
    <h1 ref="tag">BeforeUnmountAndMounted component</h1> 
</template>