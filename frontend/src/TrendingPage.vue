<template>
    <AppNavbar />
    <QuizGrid />
    <FilterBox />
    <MainFooter />
  </template>
  <script>
  import AppNavbar from './components/AppNavbar.vue'
  import MainFooter from './components/MainFooter.vue'
  import QuizGrid from './components/QuizGrid.vue'
  import FilterBox from './components/FilterBox.vue'
  
  export default {
    name: 'App',
    components: {
      AppNavbar,
      MainFooter,
      QuizGrid,
      FilterBox
    }
  }
  </script>
  
  