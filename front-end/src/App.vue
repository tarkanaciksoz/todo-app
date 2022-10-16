<template>
  <div id="app">
    <h1 class="todo-list-application-head">TODO LIST APPLICATION</h1>
    <div class="form-group">
      <input type="text" id="todo-input" name="todo-input" v-model="todoInput" placeholder="Give Todo..."/>
      <button class="add-button" id="add-todo-button" @click="addTodo">ADD</button>
      <button class="delete-button" id="deleteAllButton" @click="deleteAll">DELETE ALL</button>
    </div>
    <div class="todo-list-container">
      <div class="todo-list-head">
        <h2>TODO LIST</h2>
      </div>
      <div class="todo-list" id="todo-list">
        <div v-for="(todo, index) in todoList" :key="index" :index="index + 1" class="todo-row">
          <div class="todo-col-1">
            <p>
              #{{ index + 1}}
            </p>
          </div>
          <div class="todo-col-7">
            <div class="todo" :id="'mark-todo-' + todo.id" @click="markTodo(index)" :style="todo.marked == 1 ? 'text-decoration:line-through' : 'text-decoration:none'">{{ todo.value }}</div>
          </div>
          <div class="todo-col-4">
            <button class="delete-button" :id="'delete-todo-' + todo.id" @click="deleteTodo(todo.id)">DELETE</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import API from '@/api'

export default {
  name: 'App',
  components: {
  },
  created() {
    this.$watch(
      () => {},
      () => {
        this.getTodos()
      },
      {
        immediate: true
      }
    )
  },
  data: function () {
    return {
      todoInput: '',
      todoList: []
    }
  },
  methods: {
    getTodos: async function () {
      const todoListResponse = await API.getTodosRequest()
      if (todoListResponse.success) {
        this.todoList = todoListResponse.data
      }
    },
    addTodo: async function () {
      if (!(this.todoInput.length > 0)) {
        return
      }

      const createTodoResponse = await API.createTodoRequest({
        id: this.todoList.length + 1,
        value: this.todoInput,
        marked: 0
      })

      if(createTodoResponse.success) {
        this.todoList.push(createTodoResponse.data)
        this.todoInput = ''
      }

    },
    deleteTodo: async function (todoId) {
      const todoIndex = this.todoList.map(todo => todo.id).indexOf(todoId)
      const todo = this.todoList[todoIndex]

      const deleteTodoResponse = await API.deleteTodoRequest(todo)
      if(deleteTodoResponse.success) {
        this.todoList.splice(todoIndex, 1)
      }
    },
    deleteAll: async function () {
      const deleteTodoResponse = await API.deleteAllTodosRequest()
      if(deleteTodoResponse.success) {
        this.todoList = []
      }
    },
    markTodo: async function (todoIndex) {
      const todo = this.todoList[todoIndex]
      if (todo.marked === 1) {
        todo.marked = 0
      } else {
        todo.marked = 1
      }

      const todoMarkResponse = await API.markTodoRequest(todo)
      if (todoMarkResponse.success) {
        this.todoList[todoIndex] = todo
      }
    }
  }
}
</script>

<style>
* {
  margin: 0;
  padding: 0;
}

html, body {
  width: 100%;
  height: 100%;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

.todo-list-application-head {
  margin-left: auto;
  margin-right: auto;
  color: tomato;
}

.form-group {
  display: flex;
  gap: 10px;
  margin-top: 5%;
  margin-left: auto;
  margin-right: auto;
  justify-content: center;
}

.add-button {
  background-color: green;
  color: white;
}

.delete-button {
  background-color: red;
  color: white;
  height: 100%;
}

button {
  padding: 10px;
  border-radius: 5px;
  border: solid 1px gray;
  cursor: pointer;
}

button:hover {
  box-shadow: 0 0 10px 1px gray;
}

input {
  width: 45%;
}

.todo-list-head {
  margin-left: auto;
  margin-right: auto;
  margin-top: 100px;
  color: black;
}

.todo-list {
  margin: auto;
  max-width: 60%;
  border: solid 1px black;
  border-radius: 5px;
  box-shadow: 0 0 10px 1px rgb(180, 180, 180);
}

.todo-row {
  width: 100%;
  height: 100%;
  display: flex;
  flex-wrap: wrap;
  border-bottom: solid 1px black;
}

.todo-col-1 {
  width: 10%;
  display: block;
  border-right: solid 1px black;
  justify-content: space-between;
  line-height: 40px;
}

.todo-col-7 {
  width: 64%;
  display: block;
  border-right: solid 1px black;
}

.todo-col-4 {
  width: 25%;
  display: block;
}

.todo {
  text-align: start;
  justify-content: space-between;
  line-height: 40px;
  margin-left: 1%;
}

</style>
