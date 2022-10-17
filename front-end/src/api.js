import fetch from "node-fetch"

export class API {
    constructor (url) {
        if (url === undefined || url === "") {
            url = "http://todo-app.localhost:9090"
            //url = "http://backend:9090"
        }
        
        this.url = url
    }

    completeUrl(path) {
        return `${this.url}${path}`
    }

    getHeaders() {
        return {
        }
    }

    async getTodosRequest() {
        const request = {
            method: 'GET',
            headers: this.getHeaders(),
        }
        
        return fetch(this.completeUrl('/todo/getTodos'), request)
        .then(res => {
            return res.json()
        })
        .then(data => {
            return data
        })
    }

    async createTodoRequest(todo) {
        const request = {
            method: 'POST',
            body: JSON.stringify({
                id: todo.id,
                value: todo.value
            }),
        }

        return fetch(this.completeUrl('/todo/createTodo'), request)
        .then(res => {
            return res.json()
        })
        .then(data => {
            return data
        })
    }

    async markTodoRequest(todo) {
        const request = {
            method: 'POST',
            headers: this.getHeaders(),   
        }

        return fetch(this.completeUrl('/todo/markTodo/' + todo.id), request)
        .then(res => {
            return res.json()
        })
        .then(data => {
            return data
        })
    }

    async deleteTodoRequest(todo) {
        const request = {
            method: 'GET',
            headers: this.getHeaders(),
        }

        return fetch(this.completeUrl('/todo/deleteTodo/' + todo.id), request)
        .then(res => {
            return res.json()
        })
        .then(data => {
            return data
        })
    }

    async deleteAllTodosRequest() {
        const request = {
            method: 'GET',
            headers: this.getHeaders(),   
        }

        return fetch(this.completeUrl('/todo/deleteAllTodos'), request)
        .then(res => {
            return res.json()
        })
        .then(data => {
            return data
        })
    }
}

export default new API()