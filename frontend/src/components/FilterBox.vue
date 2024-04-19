<template>
    <div class="filter-box">
        <div class="search-area">
            <input type="text"
                placeholder="Search quizzes..."
                v-model="searchQuery"
                @input="handleSearchInput">
        </div>
        <div class="trending-filters">
            <div class="filter-option" v-for="option in options" :key="option.id">
                <input type="radio" :name="option.name" :id="option.id" :value="option.value" v-model="selectedFilter">
                <label :for="option.id">{{ option.label }}</label>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    data() {
    return {
      selectedFilter: 'weekly',
      options: [
        { id: 'daily', name: 'trend', value: 'daily', label: 'Daily' },
        { id: 'weekly', name: 'trend', value: 'weekly', label: 'Weekly' },
        { id: 'monthly', name: 'trend', value: 'monthly', label: 'Monthly' },
        { id: 'all-time', name: 'trend', value: 'all_time', label: 'All Time' },
      ]
    };
  },
  watch: {
    selectedFilter(newFilter) {
      this.$emit('filter-changed', newFilter);
    },
    searchQuery(newQuery) {
      this.$emit('search', newQuery);
    }
  },
  methods: {
    handleSearchInput() {
    this.$emit('search', this.searchQuery); 
  }
}
}
</script>


<style scoped>
   .filter-box {
    position: fixed;
    top: 40%;
    left: 20px;
    transform: translateY(-50%);
    padding: 20px;
    width: 300px;
    box-sizing: border-box;
    border-radius: 15px;
    background-color: white;
    z-index: 1000;
    box-shadow: rgba(50, 50, 93, 0.25) 0px 13px 27px -5px, rgba(0, 0, 0, 0.3) 0px 8px 16px -8px;
}

.search-area input[type="text"] {
    margin-top: 20px;
    width: 100%;
    padding: 8px;
    margin-bottom: 20px;
    box-sizing: border-box;
    border: 3px solid rgb(214, 214, 214);
}

.trending-filters {
    display: flex;
    flex-direction: column;
}

.filter-option {
    display: flex;
    align-items: center; 
    margin: 5px 0;
}
.filter-option:hover{
    cursor: pointer;
    background-color: #ffd6af;
}

.trending-filters label {
    margin-left: 8px;
    padding: 10px;
}

</style>
