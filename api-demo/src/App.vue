<script>
  import GetAPI from "./components/getData.vue"
  import SetData from "./components/setData.vue";
  export default {
    name : "App",
    components :{
      GetAPI,
      SetData
    },
     
  }
</script>

<template>
  <GetAPI/>
  <SetData/>
  </template>