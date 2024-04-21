<template>
    <div class="main-container">
        <AppNavbar />
      <div class="content">
        <AccountComponent />
    </div>
    <MainFooter />
  </div>
  </template>
  
  <script>
  import AppNavbar from './components/AppNavbar.vue';
    import MainFooter from './components/MainFooter.vue';
    import AccountComponent from './components/AccountComponent.vue';
  export default {
    components: {
        AppNavbar,
        MainFooter,
        AccountComponent
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
  