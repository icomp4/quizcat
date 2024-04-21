<template>
    <div class="quizArea" v-if="quiz.Questions && quiz.Questions.length > 0">
        <div class="timer">
            <p>⏱️{{ quizTime }}</p>
        </div>
      <div class="questionArea">
        <h1>{{ currentQuestion.text }}</h1>
      </div>
      <div class="answersArea">
        <div v-for="answer in currentQuestion.Options" :key="answer.id" @click="checkAnswer(answer)" class="answer" :class="{ 'correct-answer': showCorrectAnswer && answer.is_correct }">
          <p>{{ answer.text }}</p>
        </div>
      </div>
      <div v-if="isQuizEnd">
        <p>Quiz Finished! You scored {{ correctAnswers }}/{{ quiz.Questions.length }}</p>
        <button @click="resetQuiz" class="restartBtn">Restart Quiz</button>
      </div>
      <p>rate this quiz</p>
      <div class="rate">
        <input type="radio" id="star5" name="rate" value="5" v-model="rating" />
        <label for="star5" title="text">5 stars</label>
        <input type="radio" id="star4" name="rate" value="4" v-model="rating" />
        <label for="star4" title="text">4 stars</label>
        <input type="radio" id="star3" name="rate" value="3" v-model="rating" />
        <label for="star3" title="text">3 stars</label>
        <input type="radio" id="star2" name="rate" value="2" v-model="rating" />
        <label for="star2" title="text">2 stars</label>
        <input type="radio" id="star1" name="rate" value="1" v-model="rating" />
        <label for="star1" title="text">1 star</label>
      </div>
      <button @click="submitRating" class="submitRating">submit rating</button>
    </div>
    <div v-else>
      <p>Loading quiz...</p>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        quiz: { Questions: [] },
        currentQuestionIndex: 0,
        correctAnswers: 0,
        isQuizEnd: false,
        selectedAnswer: null,
        showCorrectAnswer: false,
        quizStartTime: null,
        quizTime: '00:00',
        rating: null,
      };
    },
    computed: {
      currentQuestion() {
        return this.quiz.Questions ? this.quiz.Questions[this.currentQuestionIndex] : null;
      },
      questionsLength() {
        return this.quiz.Questions ? this.quiz.Questions.length : 0;
      },
    },
    methods: {
      submitRating() {
      const quizId = this.$route.params.id;
      fetch(`http://localhost:8080/api/quiz/${quizId}/rate`, {
        credentials: 'include', 
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ rating: parseFloat(this.rating) }), 
      })
        .then(response => response.json())
        .then(data => {
          console.log('Rating submitted:', data);
        })
        .catch(error => {
          console.error('Error submitting rating:', error);
        });
    },
      startQuiz() {
        this.quizStartTime = Date.now();
        this.updateQuizTime();
      },
      updateQuizTime() {
        this.quizTimeUpdater = setInterval(() => {
          const totalSeconds = Math.floor((Date.now() - this.quizStartTime) / 1000);
          const minutes = Math.floor(totalSeconds / 60);
          const seconds = totalSeconds % 60;
          this.quizTime = `${minutes.toString().padStart(2, '0')}:${seconds.toString().padStart(2, '0')}`;
        }, 1000);
      },
      endQuiz() {
        clearInterval(this.quizTimeUpdater);
      },
      shuffleArray(array) {
        for (let i = array.length - 1; i > 0; i--) {
          let j = Math.floor(Math.random() * (i + 1));
          [array[i], array[j]] = [array[j], array[i]];
        }
        return array;
      },
      checkAnswer(answer) {
        this.selectedAnswer = answer;
        this.showCorrectAnswer = true;
        if (answer.is_correct && this.isQuizEnd === false) {
          this.correctAnswers++;
        }
        setTimeout(() => {
          if (this.currentQuestionIndex + 1 === this.questionsLength) {
            this.isQuizEnd = true;
            this.endQuiz(); 
          } else {
            this.nextQuestion();
          }
          this.selectedAnswer = null;
          this.showCorrectAnswer = false;
        }, 1500);
      },
      nextQuestion() {
        if (this.currentQuestionIndex + 1 < this.questionsLength) {
          this.currentQuestionIndex++;
        }
      },
      resetQuiz() {
        this.currentQuestionIndex = 0;
        this.correctAnswers = 0;
        this.isQuizEnd = false;
        this.startQuiz();
      },
      fetchQuizData() {
        const quizId = this.$route.params.id;
        fetch(`http://localhost:8080/api/quiz/${quizId}`)
          .then(response => response.json())
          .then(data => {
            this.quiz = data;
            this.quiz.Questions.forEach(question => {
              question.Options = this.shuffleArray(question.Options);
            });
          })
          .catch(error => {
            console.error('Error fetching quiz data:', error);
          });
      },
    },
    mounted() {
      this.fetchQuizData();
      this.startQuiz();
    },
  }
  </script>
  
  <style scoped>
  .answersArea .answer.correct-answer {
    background-color: #4caf50;
    transform: scale(1.1);
    }
    .answersArea .answer.correct-answer:hover {
    background-color: #4caf50;
    }
    .rate {
    float: left;
    height: 46px;
    padding: 0 10px;
    }
    .rate:not(:checked) > input {
        position:absolute;
        top:-9999px;
    }
    .rate:not(:checked) > label {
        float:right;
        width:1em;
        overflow:hidden;
        white-space:nowrap;
        cursor:pointer;
        font-size:30px;
        color:#ccc;
    }
    .rate:not(:checked) > label:before {
        content: '★ ';
    }
    .rate > input:checked ~ label {
        color: #ffc700;    
    }
    .rate:not(:checked) > label:hover,
    .rate:not(:checked) > label:hover ~ label {
        color: #deb217;  
    }
    .rate > input:checked + label:hover,
    .rate > input:checked + label:hover ~ label,
    .rate > input:checked ~ label:hover,
    .rate > input:checked ~ label:hover ~ label,
    .rate > label:hover ~ input:checked ~ label {
        color: #c59b08;
    }
    .submitRating{
        display: block;
        padding: 10px 15px;
        background-color: #252525;
        color: white;
        border: none;
        border-radius: 5px;
        cursor: pointer;
    }
    .submitRating:hover {
        background-color: #474747;
    }
  .quizArea {
        margin-top: 50px;
        display: flex;
        flex-direction: column;
        align-items: center;
    }
  .questionArea {
        margin-bottom: 20px;
    }
    .answersArea {
        display: grid;
        grid-template-columns: repeat(2, 1fr);
        gap: 20px;
        background-color: beige;
        padding: 20px;
        background-color: white;
    }
    .answersArea .answer {
        font-weight: bold;
        color: white;
        cursor: pointer;
        padding: 1px 100px;
        background-color: #ff994f; 
        border-radius: 5px;
        text-align: center;
    }
    .answersArea .answer:hover {
        background-color: #d8874d; 
    }
    .restartBtn {
        display: block;
        margin: 20px auto;
        padding: 10px 20px;
        background-color: #252525;
        color: white;
        border: none;
        border-radius: 5px;
        cursor: pointer;
    }
  </style>
  