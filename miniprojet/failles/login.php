<html lang="fr-FR">
<head>
    <meta http-equiv="content-type" content="text/html; charset=UTF-8">
    <meta charset="UTF-8">
    <script>
        window.onload = function() {
            // getMessages
        }
    </script>
</head>
<body>
    <header>
        <h1>FAILLES SQL : ('OR 1=1 -- )SECURISEES </h1>
        <a href="index.php">Accueil</a>
    </header>
    <main>
        <form action ="" method="post">
            <div id="content"/>
            <input type="text" name="name" placeholder="Name"/><br>
            <input type="text" name="password" placeholder="Password"/><br>
            <input type="submit" name="submit" value ="Valider"/>
        </form>

    </main>
</body>
</html>
<?php
if (isset($_POST['submit'])){
        if(checkLogin($_POST['name'],$_POST['password']))
            header("Location: ./success.php");
        else
            header("Location: ./denied.php");

}

function checkLogin($name, $pass){
    require('../bdconnect.php');
    $sql = "SELECT * FROM users WHERE name = '".$name." ' AND password = '".$pass." '";
    $result = $pdo->query($sql);    
    if($result->rowCount()>0) 
        return true;
    return false;
}
?>