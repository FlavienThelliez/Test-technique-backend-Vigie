# Test-technique-backend-Vigie
Test technique backend de Vigie, Janvier 2026.

Instructions : https://www.notion.so/Test-technique-backend-Vigie-2c37b34ce5ee80bd8580dd7df2fb8854

## instructions pour lancer le programme
Ouvrez un terminal dans le dossier contenant "main.go" et "orders.json"
Dans ce terminal, utilisez la commande suivante :
`go run main.go orders.json`

Vous pouvez aussi remplacer "orders.json" par un autre fichier .json au format de données identique. Si vous n'entrez pas de nom de fichier après main.go, le programme utilisera le fichier orders.json du même dossier.

Manquant d'expérience et devant rendre ce projet pour le 29/01/2026, ce programme ne remplie malheureusement pas les fonctions demandé.
En l'état, il va juste confirmer qu'il est capable d'ouvrir le fichier orders.json passé en argument (ou celui stocké en cas d'absence d'argument), et afficher la totalité des commandes.

Cet affichage m'était nécessaire pour m'assurer d'être capable de récupérer toute les informations et pouvoir les traiter individuellements.

## 6. Questions “mindset”
1. Si ce programme tournait en production, **que surveiller / logger** en priorité ?
Les montants négatifs semblent plus problématique, présentant un risque de fausser les revenus.

2. Si le fichier passait de **10 Ko → 10 Go**, que changerais-tu dans ton approche ?
Pour éviter de manipuler un array go trop large, je découperais par jour lecture des commandes. Il est même envisageable de faire du multithreading pour exploiter le potentiel de Go ainsi. Une version simple serait un thread pour les dates pairs, et un autre pour les dates impairs. Avec les sécurités nécessaires lorsque l'on rassemble le revenu total.

3. Quel est selon toi le **cas de test prioritaire**, et pourquoi ?
Vérifier le format du fichier .json passé en paramètre. Si le format est incorrect, le programme est incapable de l'utiliser.
