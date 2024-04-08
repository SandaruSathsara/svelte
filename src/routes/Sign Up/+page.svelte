<script>

    let username = '';
    let email = '';
    let password = '';
    let errors = {
        username: '',
        email: '',
        password: ''
    };

    function handleSubmit() {
        
        errors = {
            username: '',
            email: '',
            password: ''
        };

        if (!username) {
            errors.username = 'Username is required';
        }
        if (!email) {
            errors.email = 'Email is required';
        } else if (!isValidEmail(email)) {
            errors.email = 'Invalid email format';
        }
        if (!password) {
            errors.password = 'Password is required';
        } else if (password.length < 6) {
            errors.password = 'Password must be at least 6 characters long';
        }

        if (Object.values(errors).every(error => !error)) {
            console.log('Submitted:', { username, email, password });
        }
    }

    // @ts-ignore
    function isValidEmail(email) {
        return /\S+@\S+\.\S+/.test(email);
    }

</script>

<div class="container">
    <h2>Sign Up</h2>
    <form on:submit|preventDefault={handleSubmit}>
        <div class="form-group">
            <label for="username">Username:</label>
            <input type="text" id="username" bind:value={username} required>
            {#if errors.username}
                <p style="color: red;">{errors.username}</p>
            {/if}
        </div>
        <div class="form-group">
            <label for="email">Email:</label>
            <input type="email" id="email" bind:value={email} required>
            {#if errors.email}
                <p style="color: red;">{errors.email}</p>
            {/if}
        </div>
        <div class="form-group">
            <label for="password">Password:</label>
            <input type="password" id="password" bind:value={password} required>
            {#if errors.password}
                <p style="color: red;">{errors.password}</p>
            {/if}
        </div>
        <div class="form-group">
            <button type="submit">Sign Up</button>
        </div>
    </form>

    <div class="sign-in-link">
        <p>Already have an account? <a href="/Login">Sign in</a></p>
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
    input[type="email"],
    input[type="password"] {
        width: 100%;
        padding: 10px;
        border: 1px solid #ccc;
        border-radius: 3px;
    }

    button {
        width: 100%;
        padding: 10px;
        background-color:black;
        color: #fff;
        font-weight: bolder;
        border: none;
        border-radius: 3px;
        cursor: pointer;
        font-size: medium;
    }

    button:hover {
        background-color: greenyellow;
    }    

    .sign-in-link {
        text-align: center;
        margin-top: 20px;
    }

    .sign-in-link a {
        color: greenyellow;
        text-decoration: none;
        font-weight: bolder;
    }

    .sign-in-link a:hover {
        text-decoration: underline;
    }

    p {
        font-family: sans-serif;
        margin: 5px 0 0; 
    }

</style>
