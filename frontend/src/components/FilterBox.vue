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
    box-shadow: rgba(100, 100, 111, 0.2) 0px 7px 29px 0px;
}

.search-area input[type="text"] {
    margin-top: 20px;
    width: 100%;
    padding: 8px;
    margin-bottom: 20px;
    box-sizing: border-box;
    border: 2px solid rgb(158, 158, 158);
    border-radius: 25px;
}

.trending-filters {
    display: flex;
    flex-direction: column;
}

.filter-option {
    display: flex;
    align-items: center; 
    margin: 5px 0;
    position: relative; 
}
.filter-option:hover {
    cursor: pointer;
    background-color: #d4d4d4;
}

.trending-filters label {
    margin-left: 8px;
    padding: 10px;
}

.filter-option input[type="radio"] {
    opacity: 0;
    position: absolute;
    left: 10px; 
    width: 20px; 
    height: 100%;
    z-index: 1; 
}

.filter-option label {
    padding-left: 30px; 
    cursor: pointer;
    display: flex;
    align-items: center;
    width: 100%;
}
.filter-option label::before {
    content: '';
    position: absolute;
    left: 10px;
    top: 50%;
    transform: translateY(-50%);
    width: 20px;
    height: 20px;
    border-radius: 50%;
    background: #ddd;
    transition: background 0.3s;
}
.filter-option input[type="radio"]:checked + label::before {
    background: #ff994f;
}
</style>
