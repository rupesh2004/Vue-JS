<template>
  <div class="container">
    <h1 class="title">To-Do List</h1>
    <div class="input-group">
      <input
        type="text"
        class="input-box"
        placeholder="Enter Title"
        v-model="title"
      />
      <input
        type="text"
        class="input-box"
        placeholder="Enter Content"
        v-model="content"
      />
    </div>
    <div class="button-group">
      <!-- Button dynamically changes based on edit mode -->
      <button class="btn" @click="id ? saveTodo() : addTodo()">
        {{ id ? "Save Changes" : "Add Task" }}
      </button>
      <button class="btn btn-secondary" @click="viewTodos">View Notes</button>
    </div>
    <div class="display-group">
      <ul class="todo-list">
        <li class="todo-item" v-for="items in todoList.todoArray" :key="items.id">
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
import { useTodo } from "../store.js";
import { ref } from "vue";

export default {
  name: "TodoList",
  setup() {
    const todoList = useTodo();
    const title = ref("");
    const content = ref("");
    const id = ref(0); // Keeps track of the current ID for editing

    const addTodo = () => {
      const maxId = Math.max(0, ...todoList.todoArray.map((todo) => todo.id));
      todoList.addTodos(maxId + 1, title.value, content.value);
      alert("Note Added!");
      title.value = "";
      content.value = "";
    };

    const deleteTodo = (id) => {
      todoList.deleteTodo(id);
      alert("Note Deleted!");
    };

    const editTodo = (item) => {
      id.value = item.id; // Assign the current item's ID to track edits
      title.value = item.title;
      content.value = item.content;
    };

    const saveTodo = () => {
      const index = todoList.todoArray.findIndex((todo) => todo.id === id.value);
      if (index !== -1) {
        todoList.todoArray[index].title = title.value;
        todoList.todoArray[index].content = content.value;
        alert("Note Edited Successfully!");
        title.value = "";
        content.value = "";
        id.value = 0;
      } else {
        alert("Note not found!");
      }
    };

    const viewTodos = () => {
      console.log(todoList.todoArray);
    };

    return {
      todoList,
      title,
      content,
      id,
      addTodo,
      deleteTodo,
      editTodo,
      saveTodo,
      viewTodos,
    };
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
  margin-bottom: 15px;
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
