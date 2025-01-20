<script>
    export default {
        name : 'EmitsComponent',
        emits : ['customEvent'],
        setup(props,{emit}){
            const handleClick = () => {
                emit('customEvent', 'Hello, World!');
            }
            return {handleClick}
        }

    }
</script>

<template>
    <h1>Emits component</h1>
    <button @click="handleClick">Emit Event</button>
</template>