<template>
    <div>
      <h1>Received Data</h1>
      <p>Username: {{ username }}</p>
      <p>Password: {{ password }}</p>
    </div>
  </template>
  
  <script>
  export default {
    name : "Receiver",
    computed: {
      username() {
        return this.$route.query.username;
      },
      password() {
        return this.$route.query.password;
      }
    }
  };
  </script>
  