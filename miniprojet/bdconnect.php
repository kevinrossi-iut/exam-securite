
<?php
$name="";
$pass="";
$dbname="rossi62u_securite";
$host="devbdd.iutmetz.univ-lorraine.fr";
try{
    $pdo=new PDO('mysql:host='.$host.';port=3306;dbname='.$dbname.'',$name , $pass);

    $pdo->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);


} catch (PDOException $e) {
    print "Erreur !: " . $e->getMessage() . "<br/>";
    die();
}
?>
