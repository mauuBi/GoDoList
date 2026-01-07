# Go SQLite To-Do List

A robust Command Line Interface (CLI) task manager built with **Golang**. 

## The "Progress" Goal
This isn't just another "Hello World" project. This repository was created with the specific goal of **real-world skill progression**. It focuses on:
* **Modular Architecture**: Moving beyond single-file scripts to multi-package projects.
* **Database Persistence**: Implementing an ORM (GORM) with SQLite to handle data lifecycle.
* **Pointer Mastery**: Understanding how to share database instances across the app without memory overhead.
* **Data Structures**: Utilizing `structs` for models and `Hashmaps` for optimized in-memory data handling.

## Features

* **Persistent Storage**: Full integration with SQLite using **GORM**.
* **User Isolation**: Tasks are linked to a specific `UserID`.
* **Intelligent CLI**: Uses `bufio.Scanner` to handle full sentences (e.g., "Buy groceries for dinner").
* **Automatic Migrations**: The database schema updates automatically upon launch.
* **Custom Stringer**: Specialized methods to display task data clearly in the terminal.

## Tech Stack

* **Language**: Go (Golang)
* **Database**: SQLite
* **ORM**: [GORM](https://gorm.io/)
* **Input Handling**: `bufio` (Standard Library)

## Project Structure

```text
.
├── main.go               # Entry point & CLI Loop
├── handleDatabase/       # GORM logic & CRUD operations
│   └── database.go
├── structhandler/        # Task models & String methods
│   └── task.go
├── .gitignore            # Keeps your .db and binaries out of Git
└── go.mod                # Dependency management
