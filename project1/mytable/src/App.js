import './App.css';
import React, { useEffect, useState } from 'react';
import axios from 'axios';
// test

function App() {
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  const [todos, setTodos] = useState([]);

  const handleTitleChange = (e) =>{
    setTitle(e.target.value);
  }

  const handleDescriptionChange = (e) => {
    setDescription(e.target.value);
  } 

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await axios.post('http://localhost:8080/v1/todo', {
        title,
        description,
      });
      setTodos((prevTodos) => [...prevTodos, response.data]);
      setTitle('');
      setDescription('');
    } catch (error) {
      console.error("error creating todo", error)
    }
  };


  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axios.get('http://localhost:8080/v1/todo');
        setTodos(response.data);
      } catch (error) {
        console.error('Error fetching todos:', error.message);
        // You can set an error state here or show a message to the user
      }
    };
  
    fetchData();
  }, []);
  
  const handleDelete = async(id)=>{
    try {
      console.log(`Deleting todo with ID:`, id);
      await axios.delete(`http://localhost:8080/v1/todo/${id}`);
      console.log(`Todo with ID ${id} deleted successfully`);
      setTodos((prevToDos)=>prevToDos.filter(todo=>todo.ID !== id));
    } catch (error) {
      console.error("error", error)
    }
  };


  return (
    <div className="App">
      <h1>To do list</h1>
      <form onSubmit={handleSubmit}>
      <label>
        Title:
        <input type="text" value={title} onChange={handleTitleChange} />
      </label>
      <br />
      <label>
        Description:
        <input type="text" value={description} onChange={handleDescriptionChange} />
      </label>
      <br />
      <button type="submit">Create Todo</button>
    </form>
      <ul>
        {todos.map(todo => (
          <li key={todo.ID}>
              <strong>{todo.Title}</strong>
              <p>{todo.Description}</p>
              <button onClick={()=>handleDelete(todo.ID)}>Delete</button>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default App;
