<template>
    <div class="create-quiz-form">
        <h1>Create a Quiz</h1>
        <div v-if="isAuthenticated">
            <form class="createQuiz" @submit.prevent="handleSubmit">
                <input type="text" ref="quiztitle" name="quiztitle" id="quiztitle" placeholder="Quiz Title..." required>
                <input type="text" ref="quizCover" name="quizCover" id="quizCover" placeholder="Quiz Cover Image URL... (optional)">

                <div class="question-wrapper" v-for="(question, qIndex) in quizQuestions" :key="qIndex">
                    <div class="question-group">
                        <div class="question-container">
                            <input type="text" v-model="question.text" :placeholder="'Question ' + (qIndex + 1)" required>
                        </div>
                        <div class="options" v-for="(option, oIndex) in question.options" :key="`option-${qIndex}-${oIndex}`">
                            <div class="option-container">
                                <input type="text" v-model="option.text" :placeholder="'Option ' + (oIndex + 1)" required>
                                <input type="radio" v-model="option.isCorrect" :value="true" :name="`isCorrect-${qIndex}-${oIndex}`">
                                <label class="correctTxt">correct</label>
                                <div class="button-group">
                                    <button class="add-btn" @click.prevent="addOption(qIndex)">+</button>
                                    <button class="remove-btn" @click.prevent="removeOption(qIndex, oIndex)">-</button>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="button-group">
                        <button class="add-btn" @click.prevent="addQuestion">+</button>
                        <button class="remove-btn" @click.prevent="removeQuestion(qIndex)">-</button>
                    </div>
                </div>
                <button type="submit" class="create-btn">Create Quiz</button>
            </form>
        </div>
        <p v-else>Please <router-link to="/login" class="loginLink" style="color: #4169E1; text-decoration: none;">Login</router-link> to create a quiz!</p>
    </div>
</template>


<script>
export default {
    data() {
        return {
            isAuthenticated: false, 
            userID: "",
            quizQuestions: [{ text: '', options: [{ text: '', isCorrect: false }] }],
        };
    },
    mounted() {
        this.checkAuthentication();
    },
    methods: {
        checkAuthentication() {
            fetch('/api/isAuth', {
                credentials: 'include',
            })
            .then(response => response.json())
            .then(data => {
                this.isAuthenticated = data.isAuth;
                this.userID = data.userID;
            });
        },
        addQuestion() {
            this.quizQuestions.push({ text: '', options: [{ text: '', isCorrect: false }] });
        },
        removeQuestion(index) {
            this.quizQuestions.splice(index, 1);
        },
        addOption(qIndex) {
            this.quizQuestions[qIndex].options.push({ text: '', isCorrect: false });
        },
        removeOption(qIndex, oIndex) {
            let options = this.quizQuestions[qIndex].options;
            if (options.length > 1) {
                options.splice(oIndex, 1);
            }
        },
        isEveryQuestionHasCorrectOption() {
        return this.quizQuestions.every(question => 
            question.options.some(option => option.isCorrect)
            );
        },
        handleSubmit() {
            if (!this.isEveryQuestionHasCorrectOption()) {
                alert('Each question must have at least one correct option.');
            return;
            }
            const title = this.$refs.quiztitle ? this.$refs.quiztitle.value : '';
            const cover = this.$refs.quizCover ? this.$refs.quizCover.value : '';
    
            const quizData = {
                title: title,
                author_id: this.userID,
                picture: cover,
                Questions: this.quizQuestions.map(question => ({
                    text: question.text,
                    Options: question.options.map(option => ({
                        text: option.text,
                        is_correct: option.isCorrect
                    }))
                }))
            };

            fetch('/api/quiz', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(quizData),
            })
            .then(response => {
                if (response.ok) {
                    return response.json();
                }
                throw new Error('Network response was not ok.');
            })
            .then(data => {
                console.log('Quiz created successfully!', data);
                alert('Quiz created successfully!');
            })
            .catch(error => {
                console.error('Failed to submit quiz:', error);
                alert('Failed to create quiz');
            });
        }
    }
}

</script>

<style>
.correctTxt {
    font-size: 10;
    font-style: italic;
    color: #5f5f5f;
}
.create-quiz-form {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    margin-top: 5%;
    width: 100%;
}
.question-wrapper {
    display: flex;
    justify-content: space-between;
    align-items: center;
}
.question-group {
    border: 1px solid #ccc; 
    padding: 20px; 
    margin-bottom: 20px; 
    border-radius: 5px; 
}
.createQuiz {
    display: flex;
    flex-direction: column;
    align-items: center;
    width: 100%;
    margin: auto; 
}

.createQuiz input[type="text"] {
    padding: 10px;
    margin-bottom: 10px;
    border-radius: 5px;
    border: 1px solid #ccc;
    width: 100%; 
}

.question-container, .option-container {
    display: flex;
    align-items: center;
    justify-content: space-between;
    width: 100%;
}

.button-group {
    display: flex;
}

.add-btn, .remove-btn {
    padding: 5px 10px;
    margin-left: 5px;
    border: none;
    border-radius: 10px;
    cursor: pointer;
}

.add-btn {
    background-color: rgb(92, 92, 92);
    color: white;
}

.remove-btn {
    background-color: #474747;
    color: white;
}

.create-btn {
    padding: 10px 20px;
    border: none;
    border-radius: 5px;
    background-color: #ff994f;
    color: white;
    cursor: pointer;
    font-weight: bold;
    margin-right: 55px;
}

.create-btn:hover, .add-btn:hover, .remove-btn:hover {
    opacity: 0.9;
}
.options {
    margin-top: 10px; 
}

.option-container {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 5px; 
}
</style>
