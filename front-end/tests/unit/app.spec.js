import { shallowMount } from '@vue/test-utils'
import App from '@/App.vue'

describe('App.vue', () => {
    it('assignment given, when, then - 1', async () => {
      const wrapper = shallowMount(App)
      wrapper.vm.$nextTick()
      await new Promise(r => setTimeout(r, 500))
      
      const todo = {
        id: 1,
        value: 'buy some milk',
      }

      expect(wrapper.vm.todoList.length).toBe(0)

      wrapper.find('#todo-input').setValue(todo.value)
      await wrapper.find('#add-todo-button').trigger('click')

      wrapper.vm.$nextTick()
      await new Promise(r => setTimeout(r, 500))

      expect(wrapper.vm.todoList[0].id).toBe(todo.id)
      expect(wrapper.vm.todoList[0].value).toBe(todo.value)
      expect(wrapper.find('#mark-todo-' + todo.id).text()).toBe(todo.value)
      expect(wrapper.vm.todoInput).toBe('')
    })
    it('assignment given, when, then - 2', async () => {
      const wrapper = shallowMount(App)
      wrapper.vm.$nextTick()
      await new Promise(r => setTimeout(r, 500))

      const todo = {
          id: 2,
          value: 'enjoy the assignment'
        }

      wrapper.find('#todo-input').setValue(todo.value)
      await wrapper.find('#add-todo-button').trigger('click')

      wrapper.vm.$nextTick()
      await new Promise(r => setTimeout(r, 500))

      expect(wrapper.vm.todoList[1].id).toBe(todo.id)
      expect(wrapper.vm.todoList[1].value).toBe(todo.value)
      expect(wrapper.find('#mark-todo-' + todo.id).text()).toBe(todo.value)
      expect(wrapper.vm.todoInput).toBe('')
    })
    it('assignment given, when, then - 3', async () => {
      const wrapper = shallowMount(App)
      wrapper.vm.$nextTick()
      await new Promise(r => setTimeout(r, 500))

      //clear todo list
      await wrapper.find('#deleteAllButton').trigger('click')
      wrapper.vm.$nextTick()
      await new Promise(r => setTimeout(r, 500))

      const todo = {
        id: 1,
        value: 'buy some milk',
      }

      wrapper.find('#todo-input').setValue(todo.value)
      await wrapper.find('#add-todo-button').trigger('click')
      wrapper.vm.$nextTick()
      await new Promise(r => setTimeout(r, 500))

      expect(wrapper.vm.todoList[0].id).toBe(todo.id)
      expect(wrapper.vm.todoList[0].value).toBe(todo.value)
      expect(wrapper.find('#mark-todo-' + todo.id).text()).toBe(todo.value)
      expect(wrapper.vm.todoInput).toBe('')

      wrapper.find('#mark-todo-' + todo.id).trigger('click')
      wrapper.vm.$nextTick()
      await new Promise(r => setTimeout(r, 500))

      expect(wrapper.vm.todoList[0].marked).toBe(1)
    })
    it('assignment given, when, then - 4', async () => {
      const wrapper = shallowMount(App)
      wrapper.vm.$nextTick()
      await new Promise(r => setTimeout(r, 500))

      const todo = {
        id: 1,
        value: 'buy some milk',
      }

      wrapper.find('#mark-todo-' + todo.id).trigger('click')
      wrapper.vm.$nextTick()
      await new Promise(r => setTimeout(r, 500))
      expect(wrapper.vm.todoList[0].marked).toBe(0)
    })
    it('assignment given, when, then - 5', async () => {
      const wrapper = shallowMount(App)
      wrapper.vm.$nextTick()
      await new Promise(r => setTimeout(r, 500))

      //clear todo-list
      await wrapper.find('#deleteAllButton').trigger('click')
      wrapper.vm.$nextTick()
      await new Promise(r => setTimeout(r, 500))

      const todo = {
        id: 1,
        value: 'rest for a while',
      }

      wrapper.find('#todo-input').setValue(todo.value)
      await wrapper.find('#add-todo-button').trigger('click')
      wrapper.vm.$nextTick()
      await new Promise(r => setTimeout(r, 500))

      expect(wrapper.vm.todoList[0].id).toBe(todo.id)
      expect(wrapper.vm.todoList[0].value).toBe(todo.value)
      expect(wrapper.find('#mark-todo-' + todo.id).text()).toBe(todo.value)
      expect(wrapper.vm.todoInput).toBe('')

      wrapper.find('#delete-todo-' + todo.id).trigger('click')
      wrapper.vm.$nextTick()
      await new Promise(r => setTimeout(r, 500))
      expect(wrapper.vm.todoList.length).toBe(0)
    })
    it('assignment given, when, then - 5', async () => {
      const wrapper = shallowMount(App)
      wrapper.vm.$nextTick()
      await new Promise(r => setTimeout(r, 500))

      const todos = [
        {
          id: 1,
          value: 'rest for a while',
        },
        {
          id: 2,
          value: 'drink water',
        },
      ]

      wrapper.find('#todo-input').setValue(todos[0].value)
      await wrapper.find('#add-todo-button').trigger('click')
      wrapper.vm.$nextTick()
      await new Promise(r => setTimeout(r, 500))

      expect(wrapper.vm.todoList[0].id).toBe(todos[0].id)
      expect(wrapper.vm.todoList[0].value).toBe(todos[0].value)
      expect(wrapper.find('#mark-todo-' + todos[0].id).text()).toBe(todos[0].value)
      expect(wrapper.vm.todoInput).toBe('')

      wrapper.vm.$nextTick()
      await new Promise(r => setTimeout(r, 500))

      wrapper.find('#todo-input').setValue(todos[1].value)
      await wrapper.find('#add-todo-button').trigger('click')
      wrapper.vm.$nextTick()
      await new Promise(r => setTimeout(r, 500))

      expect(wrapper.vm.todoList[1].id).toBe(todos[1].id)
      expect(wrapper.vm.todoList[1].value).toBe(todos[1].value)
      expect(wrapper.find('#mark-todo-' + todos[1].id).text()).toBe(todos[1].value)
      expect(wrapper.vm.todoInput).toBe('')

      wrapper.vm.$nextTick()
      await new Promise(r => setTimeout(r, 500))

      wrapper.find('#delete-todo-' + todos[0].id).trigger('click')
      wrapper.vm.$nextTick()
      await new Promise(r => setTimeout(r, 500))

      expect(wrapper.vm.todoList.length).toBe(1)
      expect(wrapper.vm.todoList[0].id).toBe(todos[1].id)
      expect(wrapper.vm.todoList[0].value).toBe(todos[1].value)
      expect(wrapper.find('#mark-todo-' + todos[1].id).text()).toBe(todos[1].value)
    })
    it('clear todo list', async () => {
      const wrapper = shallowMount(App)
      wrapper.vm.$nextTick()
      await new Promise(r => setTimeout(r, 500))

      await wrapper.find('#deleteAllButton').trigger('click')
      wrapper.vm.$nextTick()
      await new Promise(r => setTimeout(r, 500))
      expect(wrapper.vm.todoList.length).toBe(0)

    })
  })