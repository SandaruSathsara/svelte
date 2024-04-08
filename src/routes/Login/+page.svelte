<script>
    
</script>

<div class="container">
    <h2>Login</h2>
    <form>
        <div class="form-group">
            <label for="username">Username:</label>
            <input type="text" id="username" required>
        </div>
        <div class="form-group">
            <label for="password">Password:</label>
            <input type="password" id="password" required>
        </div>
        <div class="form-group">
            <button type="submit">Login</button>
        </div>
    </form>

    <div class="sign-up-link">
        <p>Don't have an account? <a href="/Sign Up">Sign Up</a></p>
    </div>

</div>


<style>
    .container {
        max-width: 400px;
        margin: 50px auto;
        padding: 20px;
        background-color: #fff;
        border-radius: 5px;
        box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
    }

    h2 {
        text-align: center;
        margin-bottom: 20px;
    }

    .form-group {
        margin-bottom: 20px;
    }

    label {
        display: block;
        font-weight: bold;
    }

    input[type="text"],
    input[type="password"] {
        width: 100%;
        padding: 10px;
        border: 1px solid #ccc;
        border-radius: 3px;
    }

    button {
        width: 100%;
        padding: 10px;
        background-color: black;
        color: #fff;
        font-weight: bold;
        border: none;
        border-radius: 3px;
        cursor: pointer;
        font-size: medium;
    }

    button:hover {
        background-color: greenyellow;
    }

    .sign-up-link {
        text-align: center;
        margin-top: 20px;
    }

    .sign-up-link a {
        color: greenyellow;
        text-decoration: none;
        font-weight: bolder;
    }

    .sign-up-link a:hover {
        text-decoration: underline;
    }

    p{

       font-family: sans-serif;
     }

</style>
