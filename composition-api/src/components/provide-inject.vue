<script>
import { inject } from 'vue';

export default {
  name: "ProvideAndInject",
  setup() {
   const name =inject('name')
   const age = inject('age') 
   return {
    name,age
   }
  },
};
</script>
<template>
  <h1>Provide and Inject</h1>
  <p>{{name}} {{age}}</p>
</template>
