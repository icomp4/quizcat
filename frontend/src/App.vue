<template>
  <AppNavbar />
  <div class="main-container">
    <FilterBox />
    <QuizGrid />
  </div>
</template>

<script>
import AppNavbar from './components/AppNavbar.vue'
import FilterBox from './components/FilterBox.vue'
import QuizGrid from './components/QuizGrid.vue'

export default {
  name: 'App',
  components: {
    AppNavbar,
    FilterBox,
    QuizGrid
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
