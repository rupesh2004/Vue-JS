<template>
  <div class="container">
    <h1 class="title">TO Do List</h1>
    <div class="input-group">
      <input type="text" class="input-box" placeholder="Enter Title" v-model="title">
      <input type="text" class="input-box" placeholder="Enter Content" v-model="content">
    </div>
    <div class="button-group">
      <!-- Button changes text dynamically -->
      <button class="btn" @click="addTodo">{{ btnTitle }}</button>
      <button class="btn btn-secondary">View Notes</button>
    </div>
    <div class="display-group">
      <ul class="todo-list">
        <li class="todo-item" v-for="items in todos" :key="items.id">
          <div class="todo-content">
            <p class="todo-text">
              <strong>ID: </strong>{{ items.id }}<br />
              <strong>Title: </strong>{{ items.title }}<br />
              <strong>Content: </strong>{{ items.content }}
            </p>
            <div class="action-buttons">
              <button class="btn-edit" @click="editTodo(items)">Edit</button>
              <button class="btn-delete" @click="deleteTodo(items.id)">Delete</button>
            </div>
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  name: "TodoList",
  data() {
    return {
      id: 0,
      todos: [],
      title: null,
      content: null,
      isEditing: false,  // Check if editing
      editingTodoId: null,  // Store the ID of the todo being edited
      btnTitle: "Add Note",  // Default button text
    };
  },
  methods: {
    // Method to add or update todo
    addTodo() {
      if (this.title && this.content) {
        if (this.isEditing) {
          // Find the todo by ID and update its title and content
          const todoIndex = this.todos.findIndex(todo => todo.id === this.editingTodoId);
          if (todoIndex !== -1) {
            this.todos[todoIndex].title = this.title;
            this.todos[todoIndex].content = this.content;
          }

          // Reset the editing state
          this.isEditing = false;
          this.editingTodoId = null;
          this.btnTitle = "Add Note";  // Reset button text
        } else {
          // Add new todo
          this.id++;
          this.todos.push({
            id: this.id,
            title: this.title,
            content: this.content,
          });
        }

        // Clear the inputs
        this.title = null;
        this.content = null;
      } else {
        alert("Please fill in the title and content");
      }
    },

    // Method to edit a todo
    editTodo(todo) {
      // Set the input fields to the selected todo's title and content
      this.title = todo.title;
      this.content = todo.content;

      // Set editing mode
      this.isEditing = true;

      // Store the ID of the todo being edited
      this.editingTodoId = todo.id;

      // Change the button text to "Update Note"
      this.btnTitle = "Update Note";
    },

    // Method to delete a todo
    deleteTodo(todoId) {
      // Remove the todo by its ID
      this.todos = this.todos.filter(todo => todo.id !== todoId);
    },
  },
};
</script>


<style scoped>
.container {
  max-width: 600px;
  margin: 50px auto;
  background: linear-gradient(135deg, #f3f4f6, #e9ecef);
  border-radius: 15px;
  padding: 30px;
  box-shadow: 0 6px 15px rgba(0, 0, 0, 0.2);
  text-align: center;
  font-family: 'Arial', sans-serif;
}

.title {
  font-size: 2.5rem;
  color: #333;
  margin-bottom: 20px;
  text-transform: uppercase;
  letter-spacing: 3px;
  font-weight: bold;
}

.input-group {
  display: flex;
  flex-direction: column;
  gap: 15px;
  margin-bottom: 25px;
}

.input-box {
  padding: 12px 20px;
  font-size: 1rem;
  border: 2px solid #ddd;
  border-radius: 25px;
  transition: border-color 0.3s ease, box-shadow 0.3s ease;
  background: #fff;
  box-shadow: inset 0 2px 5px rgba(0, 0, 0, 0.1);
}

.input-box:focus {
  border-color: #007bff;
  box-shadow: 0 0 8px rgba(0, 123, 255, 0.5);
  outline: none;
}

.button-group {
  display: flex;
  gap: 20px;
  justify-content: center;
}

.btn {
  padding: 12px 25px;
  font-size: 1.1rem;
  color: #fff;
  background: linear-gradient(135deg, #007bff, #0056b3);
  border: none;
  border-radius: 25px;
  cursor: pointer;
  transition: background 0.3s ease, transform 0.2s ease;
  font-weight: bold;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
}

.btn-secondary {
  background: linear-gradient(135deg, #6c757d, #495057);
}

.btn:hover {
  background: linear-gradient(135deg, #0056b3, #003f8a);
  transform: translateY(-3px);
}

.btn-secondary:hover {
  background: linear-gradient(135deg, #495057, #343a40);
  transform: translateY(-3px);
}

.btn:active {
  transform: translateY(0);
  box-shadow: inset 0 2px 5px rgba(0, 0, 0, 0.2);
}

.display-group {
  margin-top: 30px;
  text-align: left;
}

.todo-list {
  list-style-type: none;
  padding: 0;
  margin: 0;
}

.todo-item {
  background: #ffffff;
  padding-left : 10px;
  padding-right:10px;
  border-radius: 10px;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
  transition: transform 0.2s ease;
}

.todo-item:hover {
  transform: scale(1.02);
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.2);
}

.todo-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.todo-text {
  font-size: 1.2rem;
  color: #333;
  margin-right: 20px;
}

.action-buttons {
  display: flex;
  gap: 10px;
}

.btn-edit {
  padding: 8px 15px;
  font-size: 1rem;
  color: #fff;
  background: #28a745;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  transition: background 0.3s ease, transform 0.2s ease;
  font-weight: bold;
}

.btn-delete {
  padding: 8px 15px;
  font-size: 1rem;
  color: #fff;
  background: #dc3545;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  transition: background 0.3s ease, transform 0.2s ease;
  font-weight: bold;
}

.btn-edit:hover {
  background: #218838;
  transform: translateY(-3px);
}

.btn-delete:hover {
  background: #c82333;
  transform: translateY(-3px);
}

.btn-edit:active,
.btn-delete:active {
  transform: translateY(0);
  box-shadow: inset 0 2px 5px rgba(0, 0, 0, 0.2);
}
</style>
