<script>
    export default {
        name:"BeforeMountAndMounted",
        data(){
            return {
                name: 'Rupesh',
            }
        },
        beforeMount(){
            console.log('Before Mount',this.$el);
        },
        mounted(){
            console.log('Mounted',this.$el);
        }
    }
</script>

<template>
    <h1>Hello from {{ name }}</h1>
</template>