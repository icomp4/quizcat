<template>
  <AppNavbar />
  <FilterBox />
</template>

<script>
import AppNavbar from './components/AppNavbar.vue'
import FilterBox from './components/FilterBox.vue'

export default {
  name: 'App',
  components: {
    AppNavbar,
    FilterBox
  }
}
</script>

