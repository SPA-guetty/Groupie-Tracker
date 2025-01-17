# Groupie-Tracker
Le projet Groupie-Tracker consiste à ouvrir et lire une API, afin d'en récupérer ses informations.
Le but du projet est de créer un site web regroupant les données de plusieurs API données, afin de les rendre plus facile à lire et y faire différentes recherches.

# Le site web
- Tout d'abord, une barre de recherche pour rechercher:
  - Un artiste en particulier
  - Une date de concert particulière
  - Un lieu de concert (Dans une ville ou un Pays)

- Le site est composé de différents filtres:
  - Un filtre de tri par nom d'artiste croissant / décroissant
  - Un filtre qui permet de rechercher des artistes par intervalle de date de concert
  - Un filtre triant les artistes par leur date de créations par ordre croissant / décroissant
  - Une sélection de différentes tranches d'années pour la date de création du groupe
 
- Sélection d'un certain nombre d'artistes à afficher:
  - Des variables sont disponibles pour afficher un nombre définit d'artistes
  - Pour une meilleure visibilité s'il y a trop d'artistes à l'écran
 
- Des cartes pour chaque artiste !
    - Chaque carte contient l'image du groupe, son nom, sa date de création ey de son premier album
    - Cliquez dessus pour la retourner et l'agrandir ! Vous verrez ainsi les membres du groupe ainsi que ses dates et lieux de concerts !

# L'activation du site web
- Étape n°1: Importer l'intégralité des codes du Groupie-Tracker sur vscode
    - Soit en l'installant et en ouvrant le dossier sur vscode
    - Soit en l'important directement via vscode grâce à la ligne de code suivante dans le terminal: 
        git clone "https://github.com/SPA-guetty/Groupie-Tracker.git"
      
- Étape n°2: Exécuter le code suivant dans le terminal:
    go run main.go

- Étape n°3: Ouvrez votre navigateur préféré et dans la barre de recherche contenant l'URL, tapez:
    localhost:8080

# Développeurs du site
- Laurent MONNIER: [@laumonnier](https://github.com/laumonnier/)
- Thomas ORJUBIN: [@SPA-guetty](https://github.com/SPA-guetty/)
