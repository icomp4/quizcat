<template>
  <div class="main-container">
      <AppNavbar />
    <div class="content">
      <FilterBox @search="handleSearch" @filter-changed="handleFilterChange" />
      <QuizGrid :filter="currentFilter" :searchQuery="currentSearchQuery" />
  </div>
  <MainFooter />
</div>
</template>

<script>
import AppNavbar from './components/AppNavbar.vue';
import FilterBox from './components/FilterBox.vue';
import QuizGrid from './components/QuizGrid.vue';
import MainFooter from './components/MainFooter.vue';

export default {
  components: {
    FilterBox,
    QuizGrid,
    AppNavbar,
    MainFooter,
  },
  data() {
    return {
      currentFilter: 'weekly',
      currentSearchQuery: '', 
    };
  },
  methods: {
    handleFilterChange(newFilter) {
      this.currentFilter = newFilter;
      this.currentSearchQuery = '';
    },
    handleSearch(query) {
      this.currentSearchQuery = query;
    },
  },
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
