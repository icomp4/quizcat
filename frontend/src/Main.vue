<template>
  <AppNavbar />
  <HeroSection />
  <FaqSection />
  <MainFooter />
</template>
<script>
import AppNavbar from './components/AppNavbar.vue'
import HeroSection from './components/HeroSection.vue'
import FaqSection from './components/FaqSection.vue'
import MainFooter from './components/MainFooter.vue'

export default {
  name: 'App',
  components: {
    AppNavbar,
    HeroSection,
    FaqSection,
    MainFooter
  }
}
</script>

<style>
.main-container {
  display: flex;
  justify-content: center;
  padding: 0 10%;
}
</style>
