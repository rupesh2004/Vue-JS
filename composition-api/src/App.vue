<script>
  import { computed } from 'vue';
import IsRefIsReactiveToRefs from './components/isRef-isReactive-toRefs.vue';
import Methods from './components/methods.vue';
import RefReactive from './components/ref-reactive.vue'
import ComputedDemo from  './components/Computed.Vue'
import Watcher from './components/watcher.vue';

  export default {
    name : "App",
    components : {
      RefReactive,
      IsRefIsReactiveToRefs,
      Methods,
      ComputedDemo,
      Watcher
    }
  }
</script>

<template>
  <RefReactive/>
  <IsRefIsReactiveToRefs/>
  <Methods/>
  <ComputedDemo/>
  <Watcher/>
</template>