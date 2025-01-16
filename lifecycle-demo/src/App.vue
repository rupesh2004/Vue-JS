<script>
import BeforeCreateCreated from './components/beforeCreate-created.vue'
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
        BeforeUnmountAndUnMounted
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
</template>
