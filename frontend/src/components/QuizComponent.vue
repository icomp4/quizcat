<template>
    <div class="quizContainer2">
      <h1 class="quizTitle2">{{ quiz.title }}</h1>
      <div v-for="(question, index) in quiz.Questions" :key="index" class="questionContainer">
        <p class="questionText">{{ question.text }}</p>
        <ul class="optionsList">
          <li v-for="(option, optionIndex) in question.Options" :key="optionIndex" class="optionItem">
            <label class="optionLabel">
              <input type="radio" :name="'question' + index" :value="option.id" v-model="selectedOptions[index]">
              {{ option.text }}
            </label>
          </li>
        </ul>
      </div>
      <button class="submitBtn" @click="calculateScore" :disabled="!allQuestionsAnswered">Submit</button>
    </div>
  </template>
  
  <script>
  export default {
    props: ['id'],
    data() {
      return {
        quiz: {},
        selectedOptions: [],
        error: null,
      };
    },
    created() {
      fetch(`http://localhost:8080/api/quiz/${this.$route.params.id}`)
        .then(response => {
          if (!response.ok) {
            throw new Error('Network response was not ok');
          }
          return response.json();
        })
        .then(data => {
          this.quiz = data;
          this.selectedOptions = new Array(data.Questions.length).fill(null);
        })
        .catch(error => {
          this.error = error;
          console.error('There has been a problem with your fetch operation:', error);
        });
    },
    computed: {
    allQuestionsAnswered() {
      return this.selectedOptions.every(option => option !== null);
        },
    },
    
    methods: {
    calculateScore() {
      let score = 0;
      const totalQuestions = this.quiz.Questions.length;
      this.quiz.Questions.forEach((question, index) => {
        const selectedOption = question.Options.find(option => option.id === this.selectedOptions[index]);
        if (selectedOption && selectedOption.is_correct) {
          score++;
        }
      });

      fetch(`http://localhost:8080/api/quiz/${this.$route.params.id}/complete`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ score: score }),
      })
      .then(response => {
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        return response.json();
      })
      .then(() => {
        alert(`Your score is ${score}/${totalQuestions}!`);
      })
      .catch(error => {
        console.error('There has been a problem with your fetch operation:', error);
      });
    },
  },
};
</script>
  
  <style>
.quizContainer2 {
  margin-top: 50px;
  display: flex;
  flex-direction: column;
  width: auto;
  max-width: 1280px;
  padding: 20px;
  background-color: #fff;
  box-shadow: rgba(255, 255, 255, 0.1) 0px 1px 1px 0px inset, rgba(50, 50, 93, 0.25) 0px 50px 100px -20px, rgba(0, 0, 0, 0.3) 0px 30px 60px -30px;
  border-radius: 8px;
}

.quizContainer2, .questionContainer, .optionsList {
  align-items: flex-start;
  text-align: left;
}
  
.quizTitle2 {
      color: rgb(43, 117, 179);
      margin-bottom: 20px;
      font-size: 57px;
      font-weight: 600;
  }
  
.questionContainer {
      margin-bottom: 20px;
  }
  
.questionText {
      margin-bottom: 10px;
      font-size: 18px;
      color: #333; 
  }
  
.optionsList {
      padding: 2px;
        list-style: none;
  }
  
.optionItem {
      background-color: #f7f8fa; 
      border-radius: 5px; 
      cursor: pointer;
      transition: background-color 0.3s; 
  }
  
.optionItem:hover {
      background-color: rgb(223, 235, 249); 
  }
  
.optionLabel {
      display: block;
      padding: 10px;
      font-size: 16px;
      color: #333; 
  }
.optionLabel:hover {
      cursor: pointer;
  }
input[type="radio"] {
      margin-right: 10px;
  }
  
.submitBtn {
  background-color: rgb(43, 117, 179);
  border: none;
  color: white;
  padding: 10px 20px;
  text-align: center;
  text-decoration: none;
  display: block; 
  font-size: 16px;
  border-radius: 5px; 
  cursor: pointer;
  transition: background-color 0.3s; 
  margin: 0 auto; 
}

.submitBtn:disabled {
  background-color: #cccccc;
  cursor: not-allowed;
}
.submitBtn:disabled:hover {
  background-color: #b8b8b8;
}
.submitBtn:hover {
      background-color: rgb(35, 95, 145); 
  }
  </style>
  