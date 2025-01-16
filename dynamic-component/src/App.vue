<script>
  import Java from "./components/java.vue"
  import Python from "./components/python.vue"
  import Flutter from "./components/flutter.vue";

  export default {
    name:"App",
    components: {
      Java,Python,Flutter
    },
    data(){
      return{
        tab: null
      }
    }
  }
</script>

<template>
  <button @click="tab='Java'">SHow Java</button>
  <button @click="tab='Python'">show python</button>
  <button @click="tab='Flutter'">show flutter</button>
  <component :is="tab"></component>
</template>
