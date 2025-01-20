import {defineStore} from 'pinia'
export const useTodo = defineStore('todos',{
    state:()=>{
        return {
            todoArray:[]
        }
    },
    actions :{
        addTodos(id,title,content){
            const work ={
                id,title,content
            }
            this.todoArray.push(work)
        },
        deleteTodo(id){
            const index = this.todoArray.findIndex(todo=>todo.id === id)
            if(index !== -1){
                this.todoArray.splice(index,1)
            }

        }
    },
    persist :{
        enabled : true,
        strategies :[{
            Storage : localStorage,
            key : 'todoworks',
            paths : ['todoArray']
        }]

    }
})