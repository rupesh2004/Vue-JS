<script>
  import { ref, onMounted } from 'vue';  // Import onMounted from vue
  import axios from 'axios';

  export default {
    name: 'App',
    setup() {
      const message = ref('');

      const callServer = async () => {
        try {
          const response = await axios.get('http://localhost:3000');
          message.value = response.data;
        } catch (err) {
          console.error(err.message);
        }
      };
      
      onMounted(() => {
        callServer();
      });

      return {
        message,
        callServer
      };
    }
  };
</script>

<template>
  <div>
    <h1>{{ message }}</h1>
  </div>
</template>

<style>
/* You can add styling here */
</style>
