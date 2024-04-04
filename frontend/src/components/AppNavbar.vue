<template>
  <nav class="navbar">
    <router-link to="/" class="quizHeader">^ QuizCat  ^</router-link>
    <ul class="nav-links">
      <li>
        <router-link to="/test">Explore</router-link>
      </li>
      <li>
        <router-link to="/quizzes">Create</router-link>
      </li>
      <li>
        <router-link to="/create">Sign up</router-link>
      </li>
      <li>
        <router-link to="/login">Login</router-link>
      </li>
    </ul>
  </nav>
</template>

<style>


.navbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px 0;
  box-shadow: 0 3px 6px rgba(0, 0, 0, 0.1);
  border-radius: 10px;
}

.quizHeader {
  font-size: 1.5rem;
  color : var(--text-color-light);
  margin-left: 50%;
  transform: translateX(-50%);
  font-weight: bold;
  text-decoration: none;
}
.quizHeader:hover {
  cursor: pointer;
}

.nav-links {
  display: flex;
  justify-content: space-between;
  width: 20%;
  margin-right: 5%;
}
.nav-links li  {
  list-style: none;
}

.nav-links li a {
  list-style: none;
  padding: 10px;
  color: rgb(36, 36, 36);
  text-decoration: none;
  font-size: medium;
}
.nav-links li a:hover {
  color: rgb(43, 117, 179);
  cursor: pointer;
}



</style>

<script>
export default {
  name: "AppNavbar",
};
</script>