<html>
    <head>
        <title></title>
    </head>
    <body>
        <form action="/login" method="post">
            <input type="checkbox" name="interest" value="music">Music
            <input type="checkbox" name="interest" value="movies">Movies
            <input type="checkbox" name="interest" value="coding">Coding
            Username: <input type="text" name="username">
            Password: <input type="password" name="password">
            <input type="hidden" name="token" value="{{.}}">
            <input type="submit" value="Login">
        </form>
    </body>
</html>