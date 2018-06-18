import React,{Component} from 'react'

export default class TodoForm extends Component {

    render(){
        return (
            <form onSubmit={this.props.onSubmit}>
                <input type="text" onChange={this.props.onChange} value={this.props.task}/>
                 <button type="submit">Add</button>
            </form>
        )
    }
}