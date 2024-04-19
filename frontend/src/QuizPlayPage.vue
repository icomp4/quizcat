<template>
    <div class="main-container">
      <AppNavbar />
      <div class="content">
        <QuizPlayArea />
      </div>
      <MainFooter />
    </div>
  </template>


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

<script>
import AppNavbar from './components/AppNavbar.vue';
import MainFooter from './components/MainFooter.vue';
import QuizPlayArea from './components/QuizPlayArea.vue';

    export default {
    name: 'App',
    components: {
            AppNavbar,
            MainFooter,
            QuizPlayArea
        }
    }
</script>
  