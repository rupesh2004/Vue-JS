<script>
export default {
    name: "BeforeCreateAndCreated",
    beforeCreate(){
        alert("Before create")
    },
    created(){
        alert("Created")
    }
};
</script>

<template>
    <h1>Hello</h1>
</template>

