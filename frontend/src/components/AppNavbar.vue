<template>
  <nav class="navbar">
    <router-link to="/" class="quizHeader">^ QuizCat  ^</router-link>
    <ul class="nav-links">
      <li>
        <router-link to="/">Home</router-link>
      </li>
      <li>
        <router-link to="/trending">Trending</router-link>
      </li>
      <li>
        <router-link to="/create">Create</router-link>
      </li>
      <li>
        <router-link to="/signup">Sign Up</router-link>
      </li>
      <li>
        <router-link to="/login">Login</router-link>
      </li>
    </ul>
  </nav>
</template>

<style scoped>
.navbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: rgb(12, 12, 12);
  padding: 10px 0;
}
.quizHeader {
  font-size: 1.5rem;
  color : rgb(194, 194, 194);
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
  color: rgb(194, 194, 194);
  text-decoration: none;
  font-size: medium;
}
.nav-links li a:hover {
  color: #ff994f;
  cursor: pointer;
}
</style>

<script>
export default {
  name: "AppNavbar",
};
</script>