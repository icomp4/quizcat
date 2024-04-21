<template>
    <div>
        <form @submit.prevent="submitForm" class="signupForm">
            <h1 class="signupHeading">Login</h1>
            <div class="form-group">
                <label for="username">Username</label>
                <input type="text" id="username" v-model="username" class="form-control" required/>
            </div>
            <div class="form-group">
                <label for="password">Password</label>
                <input type="password" id="password" v-model="password" class="form-control" required/>
            </div>
            <button type="submit" class="btn btn-primary">Login</button>
            <p class="newTxt">New around here? <router-link to="/signup" class="signUpTxt" style="color: #4169E1; text-decoration: none;">Signup</router-link></p>
            <p class="loginFailed" v-html="loginFailedMsg"></p>
        </form>
    </div>
</template>

<script>
export default {
    data() {
        return {
            username: '',
            email: '',
            password: '',
            loginFailedMsg: '',
        };
    },
    methods: {
        submitForm() {
            const userData = {
                username: this.username,
                email: this.email,
                password: this.password,
            };
            
            fetch('/api/login', {
                method: 'POST',
                credentials: 'include',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(userData),
            })
            .then(response => {
                if (response.ok) {
                    window.location.href = '/';
                } else{
                    this.loginFailedMsg = 'Account with this username and password not found. Please try again.';
                }
            })
            .catch(error => {
                console.error('Error:', error);
                this.loginFailedMsg = 'An error occurred. Please try again.';
            });
        },
    },
};

</script>

<style>
    .newTxt {
        display: block;
        text-align: center;
        margin-top: 0.5rem;
        color: #4a4a4a;
    }
    .signupHeading {
        text-align: center;
        margin-top: 40%;
    }
    .signupForm {
        max-width: 400px;
        margin: 0 auto;
    }
    .form-group {
        margin-bottom: 1rem;
    }
    .form-control {
        width: 100%;
        padding: 0.5rem;
        font-size: 1rem;
        border: 1px solid #ccc;
        border-radius: 0.25rem;
    }
    .btn {
        padding: 0.5rem 1rem;
        font-size: 1rem;
        border: none;
        border-radius: 0.25rem;
        background-color: #ff994f;
        color: white;
        margin-left: 40%;

    }
    .btn:hover {
        background-color: #ff7733;
    }
    .loginFailed {
        color: #ed5249;
        text-align: center;
    }
</style>

