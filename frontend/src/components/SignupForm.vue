<template>
    <div>
        <form @submit.prevent="submitForm" class="signupForm">
            <h1 class="signupHeading">Sign Up</h1>
            <div class="form-group">
                <label for="email">Email</label>
                <input type="email" id="email" v-model="email" class="form-control" required/>
            </div>
            <div class="form-group">
                <label for="username">Username</label>
                <input type="text" id="username" v-model="username" class="form-control" required/>
            </div>
            <div class="form-group">
                <label for="password">Password</label>
                <input type="password" id="password" v-model="password" class="form-control" required/>
            </div>
            <button type="submit" class="btn btn-primary">Sign Up</button>
            <p class="alreadyTxt">Already have an account? <a href="login" style="color: #4169E1; text-decoration: none;">Login</a></p>
            <p class="signupFailed" v-html="signupFailedMsg"></p>
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
            signupFailedMsg: '',
        };
    },
    methods: {
        submitForm() {
            const userData = {
                username: this.username,
                email: this.email,
                password: this.password,
            };
            
            fetch('/api/signup', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(userData),
            })
            .then(response => {
                if (response.ok) {
                    window.location.href = '/';
                } else if (response.status === 500) {
                    this.signupFailedMsg = 'This email or username is already in use. If it belongs to you, please <a href="/login" style="color: #4169E1;">sign in</a>.';
                }else{
                    this.signupFailedMsg = 'An error occurred. Please try again.';
                }
            })
            .catch(error => {
                console.error('Error:', error);
                this.signupfailedMsg = 'An error occurred. Please try again.';
            });
        },
    },
};

</script>

<style>
    .alreadyTxt {
        display: block;
        text-align: center;
        margin-top: 0.5rem;
        text-decoration: none;
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
    .signupFailed {
        color: #ed5249;
        text-align: center;
    }
</style>

