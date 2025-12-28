import { useEffect, useState } from "react";

function App() {
    const [todos, setTodos] = useState([]);

    useEffect(() => {
        fetch("/api/todos")
            .then((response) => response.json())
            .then((data) => setTodos(data))
            .catch((error) => console.error("Error fetching todos:", error));
    }, []);

    return (
        <div>
            <h1>ToDo List</h1>
            <ul>
                {todos.map(todo => (
                    <li key={todo.id}>{todo.task}</li>
                ))}
            </ul>
        </div>
    );
}

export default App;
