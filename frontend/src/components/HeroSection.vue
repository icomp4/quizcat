<template>
    <div class="hero-container">
        <div class="image-container"></div>
        <div class="text-container">
            <h1>Welcome to QuizCat!</h1>
            <p class="caption">Purr-fect Trivia</p>
            <button class="discover-btn">Discover Quizzes</button>
        </div>
    </div>
</template>
<style>
.hero-container {
    display: flex;
    height: 50vh;
    background-color: #fffcf9;
}

.image-container {
    flex: 1;
    background-image: url('../assets/quizcat.png');
    background-repeat: no-repeat;
    margin-top: -50px;
}

.text-container {
    flex: 1;
    display: flex;
    flex-direction: column; 
    justify-content: center; 
    align-items: center; 
    padding: 20px;
    text-align: center; 
    margin-left: -95%;
}

.text-container h1 {
    margin: 0;
    padding-bottom: 20px; 
    font-size: 2rem;
}

.text-container .caption {
    margin: 0;
    padding-bottom: 20px; 
    font-style: italic;
    font-weight: bold;
}

.discover-btn {
    padding: 15px 25px;
    background-color: #ff994f;
    color: white;
    border: none;
    border-radius: 5px;
    cursor: pointer;
    transition: background-color 0.3s;
}
.discover-btn:hover {
    background-color: #c47740;
}
</style>
