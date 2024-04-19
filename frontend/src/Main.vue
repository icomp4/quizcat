<template>
  <div class="main-container">
    <AppNavbar />
    <div class="content">
      <HeroSection />
      <FaqSection />
    </div>
    <MainFooter />
  </div>
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
    flex-direction: column;
    min-height: 100vh;
  }
  .content {
    flex: 1;
  }
</style>
