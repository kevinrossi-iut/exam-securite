<html lang="fr-FR">
<head>
    <meta http-equiv="content-type" content="text/html; charset=UTF-8">
    <meta charset="UTF-8">
</head>
<body>
    <header>
        <a href="../secu/index.php">Failles sécurisées</a>
        <h1>FAILLES XSS</h1>
        <a href="login.php">Login</a>
    </header>
    <main>
        <form action ="" method="post">
            <div id="content"/>
            <textarea id="text" name="text" required></textarea>
            <input type="submit" name="submit" value ="Valider"/>
        </form>
        <ul>
            <?php getMessages()?>
        </ul>
    </main>
</body>
</html>
<?php 

if (isset($_POST['submit']))
        addMessage($_POST['text']);

function getMessages(){
    require('../bdconnect.php');
    $result = $pdo->query('SELECT * FROM posts');
    foreach($result as $row){
        echo "<li>".$row['text']."</li>";
    }

}
function addMessage($msg){
    require('../bdconnect.php');
    $sth = $pdo->prepare('INSERT INTO posts(text) VALUES (:msg)');
    $sth->bindValue(":msg",$msg);
    $sth->execute();
    header('Location: .');
}
?>