<template>
    <footer>
        <div class="footer-container">
            <div class="footer-links">
                <h3>Company</h3>
                <ul>
                    <li><a href="#">About Us</a></li>
                    <li><a href="#">Contact Us</a></li>
                </ul>
            </div>
            <div class="footer-links">
                <h3>Support</h3>
                <ul>
                    <li><a href="#">FAQs</a></li>
                    <li><a href="#">Terms of Service</a></li>
                </ul>
            </div>
            <div class="footer-links">
                <h3>Connect</h3>
                <ul>
                    <li><a href="https://github.com/icomp4">Github</a></li>
                    <li><a href="https://www.linkedin.com/in/isaiahcompere44/">Linkedin</a></li>
                    <li><a href="https://ko-fi.com/icomp4">Buy Me A Coffee</a></li>
                </ul>
            </div>
        </div>
    </footer>
</template>

<style scoped>
footer {
    background-color: #ffffff; 
    color: rgb(0, 0, 0);
    padding: 100px 0;
}

.footer-container {
    display: flex;
    justify-content: space-around;
    align-items: flex-start;
}

.footer-links h3 {
    color: #ff994f; 
    margin-bottom: 10px; 
}

.footer-links ul {
    list-style: none;
    padding: 0;
}

.footer-links li {
    margin-bottom: 5px;
}

.footer-links a {
    color: rgb(54, 54, 54);
    text-decoration: none;
    transition: color 0.3s; 
}

.footer-links a:hover {
    color: #ff994f;
}
</style>
