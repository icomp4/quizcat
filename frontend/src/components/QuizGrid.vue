<template>
  <div class="quizContainer">
    <div class="pageHeading">
      <h1>Explore What's <span class="hot">Hot!</span></h1>
      <p>Leap into our most popular quizzes and see what's capturing everyone's curiosity!</p>
    </div>
    <div class="quizGrid">
      <a class="quizCard" v-for="quiz in quizzes" :key="quiz.id" :href="`/quiz/${quiz.id}`" :style="{ backgroundImage: `url(${quiz.picture})` }">
        <h4 class="quizTitle">{{ quiz.title }}</h4>
        <div class="quizInfo">
          <p class="quizPlays">{{ formatPlays(quiz[filter + '_plays']) }} plays</p> 
          <p class="quizRating">{{ quiz.rating.toFixed(1) }}⭐</p>
        </div>
      </a>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    filter: String,
    searchQuery: String,
  },
  data() {
    return {
      quizzes: [],
    };
  },
  watch: {
    filter: {
      immediate: true,
      handler() {
        this.fetchQuizzes();
      },
    },
    searchQuery: {
      immediate: true, 
      handler() {
        this.fetchQuizzes();
      },
    },
  },
  methods: {
    formatDate(dateString) {
    const date = new Date(dateString);
    const month = date.getMonth() + 1;
    const day = date.getDate();
    const year = date.getFullYear().toString().substr(-2);
    return `${month.toString().padStart(2, '0')}/${day.toString().padStart(2, '0')}/${year}`;
  },
  formatPlays(plays) {
    if (plays < 1000) {
      return plays; 
    } else if (plays < 1000000) {
      return (plays / 1000).toFixed(1) + 'k'; 
    } else {
      return (plays / 1000000).toFixed(1) + 'm'; 
    }
  },
    async fetchQuizzes() {
      let url = 'http://localhost:8080/api/quizzes/top/' + this.filter;
      if (this.searchQuery) {
        url = `http://localhost:8080/api/quizzes/search?param=${this.searchQuery}`;
      }
      
      try {
        const response = await fetch(url);
        if (response.ok) {
          this.quizzes = await response.json();
        } else {
          console.error('Failed to fetch quizzes:', response.statusText);
        }
      } catch (error) {
        console.error('Error fetching quizzes:', error);
      }
    }
  },
  
}
</script>

<style>
  .pageHeading {
      text-align: center;
      margin-top: -40px;
      margin-bottom: 30px;
      color: rgb(66, 66, 66);
  }
  .hot {
      color: rgb(255, 124, 30);
      text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.2);
  }
  .pageHeading p {
      font-size: 1rem;
      font-weight: bold;
  }
  .quizContainer {
      display: flex;
      flex-direction: column;
      align-items: center;
      max-width: 60%;
      margin: 0 auto;
      margin-top: 3%;
  }
  .quizGrid {
      display: flex;
      flex-wrap: wrap;
      justify-content: center;
      gap: 20px;
      width: 100%;
      
  }
  .quizCard {
    position: relative;
    color: white;
      display: flex;
      flex-direction: column;
      justify-content: space-between; 
      padding: 20px;
      box-shadow: rgba(0, 0, 0, 0.185) 0px 3px 8px;
      border-radius: 10px;
      height: 150px;
      transition: 0.3s all ease-in-out;
      max-width: 200px;
      min-width: 200px;
      text-decoration: none;
      background-size: cover;
      
  }
  .quizCard::before {
    border-radius: 10px;
      content: "";
      position: absolute;
      top: 0;
      right: 0;
      bottom: 0;
      left: 0;
      background: rgba(0, 0, 0, 0.5); 
      z-index: 1;
    } 

  .quizTitle, .quizTags, .quizInfo {
      position: relative;
      z-index: 2;
    }
  .quizTitle, .quizDescription {
      margin: 0;
    }
    .quizTitle {
    font-size: 1.5rem;
    font-weight: bold;
    display: -webkit-box;
    -webkit-line-clamp: 4;
    -webkit-box-orient: vertical;
    overflow: hidden;
    text-overflow: ellipsis;
  }
  .quizTags {
      display: flex;
      justify-content: center; 
      flex-wrap: wrap; 
      margin-top: 10px; 
  }
  
  .tag {
      font-size: small;
      font-weight: bold;
      padding: 5px;
      margin: 5px; 
      color:  rgb(255, 255, 255);
      background-color: rgba(0, 0, 0, 0.185);
      border-radius: 5px;
  }
  .tag:hover {
      cursor: pointer;
  }
  
  .quizCard:hover {
      cursor: pointer;
      scale: 1.1;
  }
  
  .quizInfo {
  display: flex;
  justify-content: space-between; 
  flex-direction: row;
  align-items: flex-end;
  font-size: small;
  width: 100%;
  gap: 10px; 
}

.quizInfo p {
  margin: 0; 
  margin-top: 4px; 
}

</style>