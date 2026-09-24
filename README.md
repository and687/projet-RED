# projet-RED
projet RED ymmersion andrea et christophe

# La Tour Infernale

**La Tour Infernale** est un jeu de rôle (RPG) en ligne de commande
développé en Go. Le joueur crée un aventurier, affronte des monstres
étage après étage, gagne de l'expérience et de l'or, puis améliore son
équipement pour progresser.

> Projet développé en Go, avec une interface textuelle et des
> illustrations ASCII dans le terminal.

## Sommaire

-   [Fonctionnalités](#fonctionnalités)
-   [Classes de départ](#classes-de-départ)
-   [Progression et combats](#progression-et-combats)
-   [Marchand et équipement](#marchand-et-équipement)
-   [Prérequis](#prérequis)
-   [Installation et lancement](#installation-et-lancement)
-   [Commandes et menus](#commandes-et-menus)
-   [Structure du projet](#structure-du-projet)
-   [À propos](#à-propos)

## Fonctionnalités

-   Création d'un personnage avec choix du nom et de la classe.
-   Combats au tour par tour contre différents monstres.
-   Progression par niveaux et gain d'expérience.
-   Récompenses en or après les victoires.
-   Inventaire d'objets consommables et capacité extensible.
-   Achat d'armes, d'objets et de matériaux auprès du marchand.
-   Fabrication d'équipements grâce au tailleur.
-   Défense apportée par les pièces d'armure équipées.
-   Mode d'entraînement et mode « Salade de monstres ».
-   Affichage coloré et illustrations ASCII dans le terminal.

## Classes de départ

Au début d'une partie, trois classes sont proposées :

  Classe         Points de vie initiaux / maximum   Or de départ
  ------------ ---------------------------------- --------------
  Chercheur                             100 / 200            150
  Pilote                                 80 / 160            250
  Mercenaire                            150 / 300             50

Chaque personnage commence au niveau 1 et reçoit trois trousses de
soins.

## Progression et combats

Les combats se déroulent au tour par tour. Le joueur peut attaquer à
mains nues, utiliser une arme ou un objet de son inventaire. Les
monstres ripostent ensuite ; leur attaque est doublée à chaque troisième
tour. La défense cumulée du casque, du torse et des jambières réduit les
dégâts reçus.

Vaincre un monstre rapporte de l'or et de l'expérience. Le personnage
peut progresser jusqu'au niveau 10. À chaque niveau gagné, ses points de
vie maximum augmentent de 20 et ses points de vie sont restaurés au
maximum.

### Ennemis du mode histoire

  Ennemi                Dégâts de base    PV   Récompense en or   Expérience
  ------------------- ---------------- ----- ------------------ ------------
  Squelette                          8    55                 35           60
  Zombie                            25    95                 60           90
  Vampire                           60   150                110          120
  Dragon légendaire                 90   350                300          200

Les étages se débloquent en vainquant le monstre de l'étage précédent.
Après la victoire contre le Dragon légendaire, le mode « Salade de
monstres » est débloqué : des monstres sont affrontés successivement et
aléatoirement jusqu'à la défaite ou au départ du joueur.

Le mannequin d'entraînement permet de tester les combats. Les points de
vie du personnage sont restaurés à leur valeur d'avant l'entraînement.

## Marchand et équipement

Le marchand propose notamment des trousses de soins, des fléchettes
empoisonnées, des armes et des matériaux nécessaires à la fabrication.
Certains articles sont disponibles en quantité limitée dans la partie.

### Armes

  Arme                                 Dégâts par utilisation   Prix affiché au marchand
  ---------------------------------- ------------------------ --------------------------
  Double pistolets                                         20                     180 or
  Fusil d'assaut                                           30                     350 or
  Dragon Slayer                                           100                     800 or
  Épée à énergie (contenu spécial)                        200                        ---

### Objets et services

-   **Trousse de soins** : rend jusqu'à 20 PV, sans dépasser les PV
    maximum.
-   **Fléchette empoisonnée** : inflige jusqu'à 30 dégâts sur la durée.
-   **Agrandissement de l'inventaire** : permet d'augmenter sa capacité
    en échange d'or.
-   **Tailleur** : fabrique un casque, un torse ou des jambières à
    partir de matériaux et d'or.

Les équipements fabriqués apportent de la défense et occupent un
emplacement : casque, torse ou jambières.

## Prérequis

-   Go installé sur votre machine.
-   Un terminal compatible avec l'affichage des séquences ANSI pour les
    couleurs.
-   Un système d'exploitation permettant l'exécution d'un programme Go.
    Le code prévoit l'effacement du terminal sous Windows et sur les
    systèmes utilisant la commande `clear`.

Vérifiez votre installation de Go avec :

``` bash
go version
```

## Installation et lancement

Clonez le dépôt :

``` bash
git clone https://github.com/and687/projet-RED.git
cd projet-RED
```

Les fichiers transmis utilisent le package `src`. Pour lancer le jeu, le
dépôt doit également contenir son point d'entrée (`package main` avec
une fonction `main`) et sa configuration Go (`go.mod`) si le projet est
organisé en module.

Une fois ces éléments présents et le point d'entrée configuré pour
appeler la fonction de démarrage du jeu, lancez depuis la racine du
projet :

``` bash
go run .
```

**Remarque :** les fichiers fournis pour la rédaction de ce README ne
comprennent pas de fichier `main.go` ni de `go.mod`. La commande de
lancement devra donc être confirmée avec l'organisation complète du
dépôt.

## Commandes et menus

Le jeu se pilote en saisissant les numéros affichés dans le terminal.

  -----------------------------------------------------------------------
  Menu                                Options principales
  ----------------------------------- -----------------------------------
  Menu principal                      Combat, marchand, tailleur,
                                      inventaire, équipement, quitter

  Combat                              Coup de poing, armes, inventaire,
                                      retour au menu principal

  Marchand                            Acheter des objets, armes,
                                      matériaux ou agrandir l'inventaire

  Tailleur                            Fabriquer des pièces d'équipement
                                      avec les matériaux requis

  Inventaire                          Consulter et utiliser les objets
                                      disponibles
  -----------------------------------------------------------------------

## Structure du projet

Les fichiers Go transmis sont organisés par responsabilité :

  -----------------------------------------------------------------------
  Fichier                             Rôle
  ----------------------------------- -----------------------------------
  `main_menu.go`                      Présentation, choix de classe,
                                      démarrage et menu principal

  `character.go`                      Données du personnage, création,
                                      affichage, expérience et niveaux

  `combat_menu.go`                    Sélection des combats, progression
                                      des étages et mode horde

  `combat_service.go`                 Déroulement des tours et résolution
                                      des combats

  `enemies.go`                        Structure des monstres et
                                      comportement de leurs attaques

  `inventory_menu.go`                 Inventaire, utilisation des objets
                                      et affichage de l'équipement

  `Weapons.go`                        Liste des armes et dégâts associés

  `shop_menu.go`                      Boutique, achats et gestion de
                                      l'inventaire

  `crafting_menu.go`                  Fabrication des équipements

  `color.go`                          Couleurs et styles ANSI pour le
                                      terminal

  `ascii.go`                          Illustrations ASCII des créatures
  -----------------------------------------------------------------------

## À propos

La Tour Infernale est un projet de jeu en Go axé sur les mécaniques de
RPG en terminal : combats, progression, gestion de ressources et
équipement.
