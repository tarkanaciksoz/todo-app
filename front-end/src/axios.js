import axios from 'axios'
import adapter from "axios/lib/adapters/http";

axios.defaults.adapter = adapter;
export class AXIOS {
    constructor (url) {
        if (url === undefined || url === "") {
            url = process.env.VUE_APP_API_URL
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
        }

        return axios(request).then(r => r.data)
    }

    async markTodoRequest(todo) {
        const request = {
            method: 'POST',
            url: this.completeUrl('/todo/markTodo/' + todo.id),
        }

        return axios(request).then(r => r.data)
    }

    async deleteTodoRequest(todo) {
        const request = {
            method: 'POST',
            url: this.completeUrl('/todo/deleteTodo/' + todo.id),
        }

        return axios(request).then(r => r.data)
    }

    async deleteAllTodosRequest() {
        const request = {
            method: 'POST',
            url: this.completeUrl('/todo/deleteAllTodos'),
        }

        return axios(request).then(r => r.data)
    }
}

export default new AXIOS(process.env.VUE_APP_API_URL)