import React from 'react'


const TodoList = (props) => {
    if(props.collection){
        return (
            <table>
                <thead>
                    <tr>
                        <th></th>
                        <th>Descrição</th>
                        <th>Ações</th>
                    </tr>
                </thead>
                <tbody>
                    {props.collection.map((item, index) => {
                        return (
                            <tr key={index}>
                               <td>
                                   <input type="checkbox" checked={item.done} data-id={index} 
                                   onChange={props.onChangeDone}/>
                                </td>
                                <td>
                                    {item.description}
                                </td>
                                <td>
                                    <button onClick={e => props.onDelete(index)}>x</button>
                                </td>
                            </tr>
                        )
                    })}
                </tbody>
            </table>
        )
    } else {
        return <p>Não há dados</p>
    }
}
export default TodoList 
