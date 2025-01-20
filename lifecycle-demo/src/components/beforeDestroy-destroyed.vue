<script>
    export default {
        name : "BeforeDestroyAndDestroyed",
        // beforeDestroyed and destroyed is replaced with beforeUnmount and unmounted
        beforeUnmount() {

            console.log("Before Destroy");
        },
        unmounted() {
            setTimeout(()=>{
                console.log("Destroyed");
            },3000)
        }
    }
</script>

<template>
    <h1>BeforeDestroy and Destroyed Component</h1>
</template>