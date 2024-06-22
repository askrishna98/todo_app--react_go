<!-- # todo_app--react_go -->

# Full-Stack Todo Application with React, Go, and MongoDB

A full-stack Todo application built with React for the frontend, Go for the backend, and MongoDB for the database, enabling users to manage their tasks efficiently.

# Features

1. **Add Task**: Users can create new tasks and add them to their todo list.
2. **Delete Task**: Users can remove tasks from their todo list.
3. **Mark Task as Done**: Users can mark tasks as completed.
4. **Mark Task as Undone**: Users can revert completed tasks back to their pending state.

# Prerequisites

- **Node.js and npm**: Make sure you have Node.js and npm installed on your machine.
- **Go**: Install the Go programming language to run the backend server.
- **MongoDB**: Set up a MongoDB database to store your application's data

# Installation

1. **Clone the Repository**

2. **Backend Setup**

   - Navigate to the `server` directory:

     ```sh
     cd server
     ```

   - Install Go dependencies:

     ```sh
     go mod tidy
     ```

   - Start the Go server:

     ```sh
     go run main.go
     ```

3. **Frontend Setup**

   - Open a new terminal window and navigate to the `todo` directory:

     ```sh
     cd ../todo
     ```

   - Install npm dependencies:

     ```sh
     npm install
     ```

   - Start the Vite development server:

     ```sh
     npm run dev
     ```

# Usage

After completing the installation steps, you should be able to access the frontend application in your web browser at the port specified by Vite and the backend server should be running on `http://localhost:9000`.You can add a new task directly from the frontend interface. Simply type your task in the input field labeled "Create Task" and press Enter or click the "Create Task" button to add it to your todo list.
