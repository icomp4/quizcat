<template>
  <div class="quizContainer">
    <div class="quizGrid">
      <a class="quizCard" v-for="quiz in quizzes" :key="quiz.id" :href="`/quiz/${quiz.id}`">
        <h4 class="quizTitle">{{ quiz.title }}</h4>
        <div class="quizTags">
        </div>
        <div class="quizInfo">
          <p class="quizDate">{{ formatDate(quiz.created_at) }}</p> 
          <p class="quizPlays">{{ quiz.all_time_plays }} plays</p>
        </div>
      </a>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      quizzes: [],
    };
  },
  created: async function() {
    const response = await fetch('http://localhost:8080/api/quizzes/top/weekly');
    const data = await response.json();
    this.quizzes = data;
  },
  methods: {
    formatDate(dateString) {
      const options = { year: '2-digit', month: '2-digit', day: '2-digit' };
      const date = new Date(dateString);
      return date.toLocaleDateString('en-US', options);
    },
  },
};
</script>
<style>
.quizContainer {
    margin-top: 50px;
    display: flex;
    flex-direction: column;
    align-items: center;
    
}
.quizGrid {
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
    gap: 20px;
    width: 100%;
}
.quizCard {
    background: linear-gradient(90deg, rgba(43,155,179,1) 0%, rgba(43,117,179,1) 100%);
    display: flex;
    flex-direction: column;
    justify-content: center; 
    align-items: center; 
    padding: 20px;
    box-shadow: rgba(0, 0, 0, 0.24) 0px 3px 8px;
    border-radius: 10px;
    height: 150px;
    transition: 0.3s all ease-in-out;
    text-decoration: none; 
    color: inherit; 
    max-width: 150px;
}
.quizTitle, .quizDescription {
    margin: 0;
}
.quizTitle {
    font-size: 1.5rem;
    font-weight: bold;
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
}
.quizInfo p {
    margin: 0; 
    margin-top: 4px; 
}

</style>

<style>
.quizContainer {
    margin-top: 50px;
    display: flex;
    flex-direction: column;
    align-items: center;
    
}
.quizGrid {
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
    gap: 20px;
    width: 100%;
}
.quizCard {
    color: white;
    background: linear-gradient(90deg, rgba(43,155,179,1) 0%, rgba(43,117,179,1) 100%);
    display: flex;
    flex-direction: column;
    justify-content: space-between; 
    padding: 20px;
    box-shadow: rgba(0, 0, 0, 0.24) 0px 3px 8px;
    border-radius: 10px;
    height: 150px;
    transition: 0.3s all ease-in-out;
}
.quizTitle, .quizDescription {
    margin: 0;
}
.quizTitle {
    font-size: 1.5rem;
    font-weight: bold;
}

.quizTags {
    display: flex;
    justify-content: flex-end;
    margin-top: 10px; 
}
.tag {
    font-size: small;
    font-weight: bold;
    padding: 5px;
    margin-right: 5px;
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
}
.quizInfo p {
    margin: 0; 
    margin-top: 4px; 
}

</style>