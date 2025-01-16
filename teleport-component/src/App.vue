<script>
import TeleportComponent from "./components/teleportComponent.vue";
export default {
  name: 'App',
  components: {
    TeleportComponent
  }
}
</script>
<template>
  <teleport to='.footer'><TeleportComponent/></teleport>

</template> 