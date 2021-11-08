# Examen de Sécurité des applications web 🔐 ![Univ Lorraine](./Logo_Univ.png)

![IUT Metz Project 2021](https://img.shields.io/badge/IUT%20Metz-2021-95a5a6.svg)

### Examen final de sécurité des applications web. Il se découpe en 3 parties :

- Lancement du projet sous forme d'exercice. le but de chaque exercice est de trouver la faille présente
- Écriture d'un rapport sur chaque faille. Pour chaque faille vous devez :
  - Décrire comment reproduire la faille
  - Décrire comment trouver la faille
  - Décrire une solution pour corriger la faille
- Écriture d'un miniprojet sur 2 versions
  - un possédant une faille XSS et SQL injection
  - un ne possédant pas la faille XSS et SQL injection

### Rendu du projet :

Vous devez fork le projet. À l'intérieur du projet vous ajoutez :
- votre rapport de sécurité au format `DOCX` OU `PDF`
- votre miniprojet dans un dossier nommé `miniprojet`
  - le `miniprojet devra possédé un micro notice de lancemant`

### Lancement du projet :
#### Avec docker

1. Installer docker ([lien d'intalation](https://www.docker.com/get-started))
2. Cloner le projet `git clone https://github.com/CharlesLgn/exam-securite.git`
3. Se Rendre dans le projet `cd exam-securite`
4. Lancer le projet `docker-compose up`
5. le site se trouve dans : `localhost:8000`

#### Sans docker (sur windows)
1. Cloner le projet `git clone https://github.com/CharlesLgn/exam-securite.git`
2. récupré le zip dans la release de github
3. déziper le dossier et lancer l'executable
4. lancer le site interter en ouvrant : `exam-securite/projet-front/intex.html`

#### Sans docker (sur unix)
1. instaler go ([lien d'instalation](https://golang.org/dl/))
2. Cloner le projet `git clone https://github.com/CharlesLgn/exam-securite.git`
3. rendez vous dans `exam-securite/projet-back`
4. `go build -o securite_back`
5. `./securite_back`
6. lancer le site interter en ouvrant : `exam-securite/projet-front/intex.html`
