<script>
    import  {useCounterStore} from '@/store/counter'
    export default {
        name : "IncrementComponent",
        setup(){
            const counterStore = useCounterStore();

            return {
                counterStore
            }
        }
        
    }
</script>

<template>
    <div>
        <h1>Count : {{ counterStore.count }} </h1>
        <h2>Double Count {{ counterStore.doubleCount }}</h2>
        <button @click="counterStore.increment">Increment</button>
    </div>
</template>