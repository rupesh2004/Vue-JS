<template>
    <button @click="sendData">Send to Parent</button>
  </template>
  
  <script>
  export default {
    name: "Child2",
    data() {
      return {
        userName: "rupesh", // Data to send to the parent
      };
    },
    methods: {
      sendData() {
        this.$emit('sendData', this.userName); // Emit event with data
      },
    },
  };
  </script>
  