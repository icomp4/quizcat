<template>
    <div class="main-container">
        <AppNavbar />
      <div class="content">
        <SignupForm />
    </div>
    <MainFooter />
  </div>
  </template>
  
  <script>
  import AppNavbar from './components/AppNavbar.vue';
  import MainFooter from './components/MainFooter.vue';
  import SignupForm from './components/SignupForm.vue';
  
  export default {
    components: {
        AppNavbar,
        MainFooter,
        SignupForm,
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
  