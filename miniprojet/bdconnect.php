
<?php
$name="rossi62u_appli";
$pass="31905588";
$dbname="rossi62u_securite";
$host="devbdd.iutmetz.univ-lorraine.fr";
try{
    $pdo=new PDO('mysql:host=devbdd.iutmetz.univ-lorraine.fr;port=3306;dbname=rossi62u_securite','rossi62u_appli','31905588');
    $pdo=new PDO('mysql:host='.$host.';port=3306;dbname='.$dbname.'',$name , $pass);

    $pdo->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);


} catch (PDOException $e) {
    print "Erreur !: " . $e->getMessage() . "<br/>";
    die();
}
?>
