<script>
import Footer from './components/Footer.vue';
import Navbar from './components/Navbar.vue';

    export default {
        name :"App",
        components: {
          Navbar,
          Foote
        }
    }
</script>

<template>
    <Navbar/>
    <router-view></router-view>
    <Footer/>
</template>