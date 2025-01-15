<script>
  import SlotsDemo from './component/SlotsDemo.vue';

  export default {
    name: "App",
    components: {
      SlotsDemo
    }
  }
</script>

<template>
  <SlotsDemo>
    <template v-slot:h1Tag>
      <h1>Hello From Rupesh</h1>
    </template>
    <template v-slot:pTag>
      <p>Hello from MR</p>
    </template>
  </SlotsDemo>
</template>
