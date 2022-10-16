import axios from 'axios'
import adapter from "axios/lib/adapters/http";

axios.defaults.adapter = adapter;
export class API {
    constructor (url) {
        if (url === undefined || url === "") {
            url = "http://localhost:9090"
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
            url: this.completeUrl('/todo/getTodos'),
            headers: this.getHeaders(),
        }

        return axios(request).then(r => r.data)
    }

    async createTodoRequest(todo) {
        const request = {
            method: 'POST',
            url: this.completeUrl('/todo/createTodo'),
            data: JSON.stringify({
                id: todo.id,
                value: todo.value
            }),
            headers: this.getHeaders(),
        }

        return axios(request).then(r => r.data)
    }

    async markTodoRequest(todo) {
        const request = {
            method: 'POST',
            url: this.completeUrl('/todo/markTodo/' + todo.id),
            headers: this.getHeaders(),
            
        }

        return axios(request).then(r => r.data)
    }

    async deleteTodoRequest(todo) {
        const request = {
            method: 'GET',
            url: this.completeUrl('/todo/deleteTodo/' + todo.id),
            headers: this.getHeaders(),
            
        }

        return axios(request).then(r => r.data)
    }

    async deleteAllTodosRequest() {
        const request = {
            method: 'GET',
            url: this.completeUrl('/todo/deleteAllTodos'),
            headers: this.getHeaders(),
            
        }

        return axios(request).then(r => r.data)
    }
}

export default new API()