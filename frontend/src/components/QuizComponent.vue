<template>
  <div class="parentContainer">
    <div class="quizContainer2" v-if="quiz && quiz.Questions">
      <h1 class="quizTitle2">{{ quiz.title }}</h1>
      <div v-if="quiz.Questions.length > 0">
        <div class="questionContainer" v-if="quiz.Questions[currentQuestionIndex]">
          <p class="questionText">{{ quiz.Questions[currentQuestionIndex].text }}</p>
          <ul class="optionsList">
            <li v-for="(option, optionIndex) in quiz.Questions[currentQuestionIndex].Options" :key="optionIndex" class="optionItem">
              <label class="optionLabel">
                <input type="radio" :name="'question' + currentQuestionIndex" :value="option.id" v-model="selectedOptions[currentQuestionIndex]">
                {{ option.text }}
              </label>
            </li>
          </ul>
        </div>
        <div class="buttonContainer">
          <button class="navBtn" @click="prevQuestion" v-if="currentQuestionIndex > 0">Previous</button>
          <button class="navBtn" @click="nextQuestion" v-if="currentQuestionIndex < quiz.Questions.length - 1">Next</button>
        </div>
      </div>
    </div>
    <button class="submitBtn" @click="calculateScore" v-if="quiz && quiz.Questions && currentQuestionIndex === quiz.Questions.length - 1" :disabled="!allQuestionsAnswered">Submit</button>
  </div>
</template>

<script>
export default {
  props: ['id'],
  data() {
    return {
      quiz: {},
      selectedOptions: [],
      currentQuestionIndex: 0, 
      error: null,
    };
  },
  computed: {
    allQuestionsAnswered() {
      return this.selectedOptions.every(option => option !== null);
    },
  },
  created() {
    fetch(`http://localhost:8080/api/quiz/${this.$route.params.id}`)
      .then(response => response.ok ? response.json() : Promise.reject('Network response was not ok'))
      .then(data => {
        this.quiz = data;
        this.selectedOptions = new Array(data.Questions.length).fill(null);
      })
      .catch(error => {
        this.error = error.toString();
        console.error('There has been a problem with your fetch operation:', error);
      });
  },
  methods: {
    calculateScore() {
      let score = 0;
      this.quiz.Questions.forEach((question, index) => {
        const selectedOption = question.Options.find(option => option.id === this.selectedOptions[index]);
        if (selectedOption && selectedOption.is_correct) {
          score++;
        }
      });

      alert(`Your score is ${score}/${this.quiz.Questions.length}`);

      if (this.allQuestionsAnswered) {
        fetch(`http://localhost:8080/api/quiz/${this.$route.params.id}/complete`, {
          method: 'PUT',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ score }),
        })
        .then(response => response.ok ? response.json() : Promise.reject('Network response was not ok'))
        .catch(error => {
          this.error = error.toString();
          console.error('There has been a problem with your fetch operation:', error);
        });
      }
    },
    nextQuestion() {
      if (this.currentQuestionIndex < this.quiz.Questions.length - 1) {
        this.currentQuestionIndex++;
      }
    },
    prevQuestion() {
      if (this.currentQuestionIndex > 0) {
        this.currentQuestionIndex--;
      }
    },
  },
};
</script>

<style>
.navBtn {
  background-color: #535353;
  border: none;
  color: white;
  padding: 10px 20px;
  margin: 5px;
  font-size: 16px;
  border-radius: 5px; 
  cursor: pointer;
  transition: background-color 0.3s;
}

.navBtn:hover {
  background-color: #777777;
}

.quizContainer2 {
  margin-top: 50px;
  display: flex;
  flex-direction: column;
  width: 100%;
  max-width: 600px;
  padding: 20px;
  background-color: #fff;
  box-shadow: 0 1px 3px rgba(0,0,0,0.12), 0 1px 2px rgba(0,0,0,0.24);
  border-radius: 8px;
  
}

.quizContainer2, .questionContainer, .optionsList {
  align-items: flex-start;
  text-align: left;
}

.submitBtn:disabled {
  background-color: #cccccc;
  cursor: not-allowed;
}
.submitBtn:disabled:hover {
  background-color: #cccccc;
}

.quizTitle2 {
  color: #007BFF;
  margin-bottom: 20px;
  font-size: 36px;
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
  background-color: #e9ecef; 
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
  background-color: #28a745;
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
  margin-top: 10px;
  margin-left: auto;
  margin-right: auto;
}

.submitBtn:disabled {
  background-color: #cccccc;
  cursor: not-allowed;
}
.submitBtn:disabled:hover {
  background-color: #b8b8b8;
}
.submitBtn:hover {
  background-color: #218838; 
}
  </style>
  