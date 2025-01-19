<script>
import { ref, reactive, isRef, isReactive, toRefs } from 'vue';

export default {
    name: "Is-RefReactiveToRefs",
    setup() {
        // Create a ref
        const name = ref("Rupesh");
        console.log('Name :', isRef(name)); // Logs true

        // Change ref value after 3 seconds
        setTimeout(() => name.value = "Shubhya", 3000);

        // Create a reactive object
        const userData = reactive({
            firstName: "Rupesh",
            lastName: "Bhosale",
            age: 21
        });
        console.log('Data :', isReactive(userData)); // Logs true

        // Return values for the template
        return {
            name,
            userData,
            ...toRefs(userData)
        };
    }
};
</script>

<template>
    <div>
        <p>Age : {{age}}</p>
    </div>
</template>
