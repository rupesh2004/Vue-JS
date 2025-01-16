<template>
    <div>
      <form @submit.prevent="submitForm">
        <input type="text" v-model="username" placeholder="Username" />
        <input type="password" v-model="password" placeholder="Password" />
        <button type="submit">Submit</button>
      </form>
    </div>
  </template>
  
  <script>
  export default {
    name:"Sender",
    data() {
      return {
        username: '',
        password: ''
      };
    },
    methods: {
      submitForm() {
        this.$router.push({ 
          name: 'secondPage',
          query: {
            username: this.username,
            password: this.password
          }
        });
      }
    }
  };
  </script>
  