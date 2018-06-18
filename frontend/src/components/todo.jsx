import React from 'react'
import axios from 'axios'
import TodoForm from './todoForm'
import TodoList from './todoList'

export default class Todo extends React.Component {

    constructor(props){
        super(props)
        this.state = { tasks: [], userInput: '' }
        this.createTasks = this.createTasks.bind(this)
        this.onChange = this.onChange.bind(this)
    }

    onChange(e){
        const {state} = this
        state.userInput = e.target.value
        this.setState(state)
    }

    refresh(){
       const {state} = this 
       state.userInput = ''
       this.setState(state)
       this.getTasks()
    }

    componentWillMount(){
        this.getTasks()
    }
    
    getTasks(){
        axios.get('http://localhost:8989/tasks/')
        .then(resp => this.setState({...this.state, tasks: resp.data}))
        console.log(this.state.tasks)
    }

    createTasks(e){
        e.preventDefault()
        axios.post('http://localhost:8989/tasks/', { description: this.state.userInput, done: false})
        .then(resp => this.refresh())
    }

    onChangeDone(e){
        const {state} = this 
        const id = e.target.getAttribute('data-id')

        state.tasks[id].done = e.target.checked 
        this.updateTask(state.tasks[id], id)
        this.setState(state)
    }

    updateTask(task, id, callback){
        axios.put(`http://localhost:8989/tasks/${id}`, task)
        .then(resp => callback(resp.data))
    }
    onDelete(id){
        axios.delete(`http://localhost:8989/tasks/${id}`)
        .then(resp => this.refresh())
    }
    render(){
        return (
            <div>
                <TodoForm onSubmit={this.createTasks} 
                onChange={this.onChange} 
                task={this.state.userInput}/>
                <TodoList collection={this.state.tasks} 
                onDelete={this.onDelete.bind(this)}
                onChangeDone={this.onChangeDone.bind(this)}/>
            </div>
        )
    }
}