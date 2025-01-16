<template>
    <div class="not-found">
      <h1>404 - Page Not Found</h1>
      <p>The page you are looking for does not exist.</p>
      <router-link to="/">Go back to Home</router-link>
    </div>
  </template>
  
  <script>
  export default {
    name: "NotFound"
  };
  </script>
  
  <style scoped>
  .not-found {
    text-align: center;
    margin-top: 50px;
  }
  
  .not-found h1 {
    font-size: 48px;
    color: #e74c3c;
  }
  
  .not-found p {
    font-size: 18px;
    color: #555;
  }
  
  router-link {
    display: inline-block;
    margin-top: 20px;
    padding: 10px 20px;
    background-color: #3498db;
    color: white;
    text-decoration: none;
    border-radius: 4px;
  }
  
  router-link:hover {
    background-color: #2980b9;
  }
  </style>
  