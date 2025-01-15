<script>
export default {
    name : "TwoWayBinding",
    data(){
        return {
            name : "John",
        }
    }

}
</script>

<template>
    <div>
        <input type="text" name="" id="" v-model="name">
        <h3>Name : {{ name }}</h3>
    </div>
</template>