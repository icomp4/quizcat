# quizCat

quizCat is a dynamic web application designed for creating and interacting with quizzes. Built with Go and the Fiber framework, it offers a lightweight, yet powerful platform for users to explore, learn or just have fun! Users can create quizzes, rate them, and track quiz plays over various periods. Additionally, QuizCat supports user authentication and categorization of quizzes, enhancing the organization and accessibility of content.

## Features

quizCat provides several endpoints for managing quizzes, users, and categories:

### Quiz Endpoints

- **POST /api/quiz**: Create a new quiz.
- **GET /api/quiz/:id**: Retrieve a quiz by its ID.
- **PUT /api/quiz/:id/complete**: Increment the play count for a quiz.
- **PUT /api/quiz/:id/rate**: Rate a quiz.
- **GET /api/quiz/:id/rating**: Get the rating of a quiz.
- **GET /api/quizzes/top/:period**: Retrieve the top quizzes for a specified period (e.g., daily, weekly, monthly).
- **GET /api/quizzes/search**: Search for quizzes based on a query parameter.

### User Endpoints

- **POST /api/signup**: Sign up a new user.
- **POST /api/login**: Log in an existing user.

### Category Endpoints

- **POST /api/category**: Create a new category for quizzes.
- **GET /api/categories**: Retrieve all categories.
- **POST /api/quiz/:id/category**: Assign a category to a quiz.
- **GET /api/quizzes/category/:category**: Retrieve quizzes by category.
