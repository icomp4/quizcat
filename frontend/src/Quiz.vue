<template>
    <AppNavbar />
    <div class="main-container">
        <QuizComponent /> 
    </div>
  </template>
  
  <script>
  import AppNavbar from './components/AppNavbar.vue'
  import QuizComponent from './components/QuizComponent.vue'

  
  export default {
    name: 'QuizPage',
    components: {
      AppNavbar,
      QuizComponent
    },
    props: ['id'],
  data() {
    return {
      quiz: {},
      error: null,
    };
  },
  created() {
    fetch(`http://localhost:8080/api/quiz/${this.id}`)
      .then(response => {
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        return response.json();
      })
      .then(data => {
        this.quiz = data;
      })
      .catch(error => {
        this.error = error;
        console.error('There has been a problem with your fetch operation:', error);
      });
  },
  }
  </script>
  
  <style>
  .main-container {
    display: flex;
    justify-content: center;
    padding: 0 10%;
  }
  </style>
  