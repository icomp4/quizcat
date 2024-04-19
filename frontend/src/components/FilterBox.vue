<template>
    <div class="filter-box">
        <div class="search-area">
            <input type="text" placeholder="Search quizzes..." name="search" id="search-quizzes">
        </div>
        <div class="trending-filters">
            <input type="radio" name="trend" id="daily" value="daily">
            <label for="daily">Daily</label>
            
            <input type="radio" name="trend" id="weekly" value="weekly">
            <label for="weekly">Weekly</label>
            
            <input type="radio" name="trend" id="monthly" value="monthly">
            <label for="monthly">Monthly</label>
            
            <input type="radio" name="trend" id="all-time" value="all-time">
            <label for="all-time">All Time</label>
        </div>
    </div>
</template>

<style>
    .filter-box {
        margin-top: -10%;
        margin-left: 20px;
        border: 2px solid #ff994f;
        padding: 20px;
        width: 300px;
        box-sizing: border-box;
        border-radius: 15px;
    }
    
    .search-area input[type="text"] {
        width: 100%;
        padding: 8px;
        margin-bottom: 20px; 
        box-sizing: border-box; 
    }
    
    .trending-filters {
        display: flex;
        flex-direction: column;
    }
    
    .trending-filters input[type="radio"] {
        margin: 5px 0;
    }
    
    .trending-filters label {
        margin-left: 8px;
    }
</style>
