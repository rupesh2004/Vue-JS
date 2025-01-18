<template>
    <div>
      <input v-model="data" placeholder="Enter some data" />
      <button @click="saveData">Save</button>
      <button @click="loadData">Load</button>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        data: ''
      };
    },
    methods: {
      saveData() {
        // Store data in localStorage
        localStorage.setItem('myData', this.data);
        alert('Data saved to localStorage!');
      },
      loadData() {
        // Retrieve data from localStorage
        const storedData = localStorage.getItem('myData');
        if (storedData) {
          this.data = storedData;
          alert('Data loaded from localStorage!');
        } else {
          alert('No data found in localStorage.');
        }
      }
    }
  };
  </script>
  