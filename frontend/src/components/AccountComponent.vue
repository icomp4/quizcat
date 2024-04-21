<template>
    <div class="account-container">
        <h1>Account</h1>
        <div class="account-info">
            <h2>Username: {{ username }}</h2>
            <h2>Email: {{ email }}</h2>
            <h2>Your Quizzes:</h2>
            <div class="quiz-row" v-for="quiz in quizzes" :key="quiz.id">
                <img :src="quiz.picture" alt="Quiz cover" class="quizCover">
                <router-link :to="'/quiz/' + quiz.id" class="quiz-title">{{ quiz.title }}</router-link>
                <p class="quizStats">Rating: {{ quiz.rating.toFixed(1) }}⭐</p>
                <p class="quizStats">Daily Plays: {{ formatPlays(quiz.daily_plays) }}</p>
                <p class="quizStats">Weekly Plays: {{ formatPlays(quiz.weekly_plays) }}</p>
                <p class="quizStats">Monthly Plays: {{ formatPlays(quiz.monthly_plays) }}</p>
                <p class="quizStats">All Time Plays: {{ formatPlays(quiz.all_time_plays) }}</p>
                <button @click="deleteQuiz(quiz.id)" class="delete-button">Delete</button>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            username: "",
            email: "",
            quizzes: [],
        };
    },
    mounted() {
        this.fetchAccountInfo();
        this.fetchUserQuizzes();
    },
    methods: {
        fetchAccountInfo() {
            fetch('http://localhost:8080/api/user', {
                credentials: 'include',
            })
            .then(response => response.json())
            .then(data => {
                this.username = data.username;
                this.email = data.email;
            });
        },
        fetchUserQuizzes() {
            fetch('http://localhost:8080/api/user/quizzes', {
                credentials: 'include',
            })
            .then(response => response.json())
            .then(data => {
                this.quizzes = data;
            });
        },
        deleteQuiz(id) {
        fetch(`http://localhost:8080/api/quiz/${id}`, {
            method: 'DELETE',
            credentials: 'include',
        })
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            this.quizzes = this.quizzes.filter(quiz => quiz.id !== id);
            return response.json();
        })
        .catch(error => {
            console.log('There was a problem with the fetch operation: ' + error.message);
        });
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
    }
}
</script>

<style scoped>
    .account-container {
        display: flex;
        flex-direction: column;
        align-items: center;
    }
    .account-info {
        margin-top: 20px;
    }
    .account-info h2 {
        margin: 10px 0;
    }
    .account-info ul {
        list-style-type: none;
        padding: 0;
    }
    .account-info li {
        margin: 5px 0;
    }
    .account-info a {
        color: #4169E1;
        text-decoration: none;
    }
    .account-info a:hover {
        text-decoration: underline;
    }
    .quiz-row {
        display: flex;
        align-items: center;
        justify-content: space-between;
        margin-bottom: 20px;
        padding: 20px;
        box-shadow: rgba(0, 0, 0, 0.16) 0px 10px 36px 0px, rgba(0, 0, 0, 0.06) 0px 0px 0px 1px;
    }
    .quizCover {
        width: 150px;
        height: 150px;
        margin-right: 20px;
    }
    .quiz-title {
        flex-grow: 1;
        margin-right: 20px;
    }
    .delete-button {
        background-color: #ff6358;
        color: white;
        border: none;
        padding: 10px 20px;
        border-radius: 5px;
        cursor: pointer;
        margin-left: 10px;
    }
    .delete-button:hover {
        background-color: #e57373;
    }
    .quizStats{
    padding-left: 5px; 
    padding-right: 25px;
    text-align: right; 
}
</style>