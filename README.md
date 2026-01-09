# GoDoList

GoDoList est une application de gestion de tâches (To-Do List) développée en langage Go. Ce projet a été réalisé dans un objectif d'apprentissage pour comprendre les fondamentaux du langage et l'interaction avec des systèmes de persistance de données.

## Objectifs du projet

L'objectif principal de ce projet était de pratiquer les concepts suivants :
* La syntaxe et la structure des projets en Go.
* L'utilisation du framework GORM (Object Relational Mapper) pour simplifier les interactions avec la base de données.
* L'implémentation et la gestion d'une base de données relationnelle légère avec SQLite.
* La manipulation des opérations CRUD (Create, Read, Update, Delete) sur des structures de données.

## Technologies utilisées

* Go : Langage de programmation principal.
* SQLite : Moteur de base de données SQL léger.
* GORM : Bibliothèque ORM pour la gestion des schémas et des requêtes.

## Fonctionnalités

Le projet permet de gérer les tâches via les fonctionnalités suivantes :
1. Ajout d'une nouvelle tâche.
2. Affichage de la liste des tâches depuis la base SQLite.
3. Modification du statut d'une tâche (terminée/en cours).
4. Suppression d'une tâche de la base de données.

## Installation et Utilisation

### Prérequis

* Go installé sur votre machine.
* Go Modules activé.

### Étapes

1. Cloner le dépôt :
   ```bash
   git clone [https://github.com/votre-utilisateur/godolist.git](https://github.com/votre-utilisateur/godolist.git)
   cd godolist

2. Installer les dépendances :
   ```bash
   go mod tidy
   
3. Lancer la ToDoList :
   ```bash
   go run .
  
