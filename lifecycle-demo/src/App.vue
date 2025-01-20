<script>
import BeforeCreateCreated from './components/beforeCreate-created.vue'
import BeforeDestroyDestroyed from './components/beforeDestroy-destroyed.vue';
import BeforeMountMounted from './components/beforeMount-mounted.vue';
import BeforeUpdateUpdated from './components/beforeUpdate-updated.vue';
import BeforeUnmountAndUnMounted from './components/beforUnmounte-unmounted.vue';

export default {
    name: "App",
    data(){
        return {
            show:false
        }
    },
    components: {
        BeforeCreateCreated,
        BeforeMountMounted,
        BeforeUpdateUpdated,
        BeforeUnmountAndUnMounted,
        BeforeDestroyDestroyed
    }
}
</script>
<template>
<BeforeCreateCreated />
<BeforeMountMounted />
<BeforeUpdateUpdated />
<BeforeUnmountAndUnMounted v-if="show" />
<br /><br />
<button @click="show=!show">Toggle </button>

<BeforeDestroyDestroyed />
</template>
