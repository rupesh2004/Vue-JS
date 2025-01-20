<script>
  import { computed } from 'vue';
import IsRefIsReactiveToRefs from './components/isRef-isReactive-toRefs.vue';
import Methods from './components/methods.vue';
import RefReactive from './components/ref-reactive.vue'
import ComputedDemo from  './components/Computed.Vue'
import Watcher from './components/watcher.vue';
import ProvideInject from './components/provide-inject.vue';
import {provide} from 'vue'
import EmitsComponent from './components/emitsComponent.vue';
  export default {
    
    name : "App",
    components : {
      RefReactive,
      IsRefIsReactiveToRefs,
      Methods,
      ComputedDemo,
      Watcher,
      ProvideInject,
      EmitsComponent
    },
    setup(){
      provide('name','rupesh')
      provide('age','21')

      const handleCustomEvent=(payload)=>{
        console.log(payload)
      }
      return {
        handleCustomEvent
      }
    }
  }
</script>

<template>
  <RefReactive/>
  <IsRefIsReactiveToRefs/>
  <Methods/>
  <ComputedDemo/>
  <Watcher/>
  <ProvideInject/>
  <EmitsComponent  @customEvent = "handleCustomEvent"/>
</template>