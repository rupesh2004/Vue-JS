<script>
import { onMounted, ref } from 'vue'
import axios from 'axios';

export default {
  name: "App",
  setup() {
    const data = ref([]);
    const dataToSend = ref({
      id: "1234",
      name: "Rupesh Bhosale"
    });

    const fetchData = async () => {
      try {
        const res = await axios.get("http://localhost:3000");
        data.value = res.data.data;
      } catch (err) {
        console.log(err);
      }
    };

    const sendToServer = async () => {
      try {
        // Sending dataToSend as JSON to the server
        const res = await axios.post("http://localhost:3000/postdata", dataToSend.value);
        console.log(res.data);  // Logs response from server
      } catch (error) {
        console.log(error);
      }
    };

    onMounted(() => {
      fetchData();
    });

    return {
      data,
      sendToServer
    };
  }
};
</script>

<template>
  <h1>connecting to go server</h1>
  <p>{{ data }}</p>

  <!-- Button to send data to the Go server -->
  <button @click="sendToServer">Send To Server</button>
</template>
