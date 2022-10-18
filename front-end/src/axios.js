import axios from 'axios'
import adapter from "axios/lib/adapters/http";

axios.defaults.adapter = adapter;
export class AXIOS {
    constructor (url) {
        console.log(process.env.VUE_APP_API_URL)
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
            "Content-type": "application/json; charset=UTF-8"
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
            method: 'DELETE',
            url: this.completeUrl('/todo/deleteTodo/' + todo.id),
            headers: {
                crossDomain: true
            },
        }

        return axios(request).then(r => r.data)
    }

    async deleteAllTodosRequest() {
        const request = {
            method: 'DELETE',
            url: this.completeUrl('/todo/deleteAllTodos'),
            headers: {
                crossDomain: true,
            },
        }

        return axios(request).then(r => r.data)
    }
}

export default new AXIOS()