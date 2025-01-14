import React, { useState } from "react";
import styles from "./todo.module.css";
import Header from "./header.tsx";

const Todo: React.FC = () => {
    const [tasks, setTasks] = useState<{ text: string; completed: boolean }[]>([]);
    const [newTask, setNewTask] = useState("");

    const handleAddTask = () => {
        if (newTask.trim()) {
            setTasks((prevTasks) => [
                ...prevTasks,
                { text: newTask.trim(), completed: false },
            ]);
            setNewTask("");
        }
    };

    const handleTaskCompletion = (index: number) => {
        const updatedTasks = [...tasks];
        updatedTasks[index].completed = !updatedTasks[index].completed;
        setTasks(updatedTasks);
    };

    const handleDeleteCompletedTasks = () => {
        setTasks(tasks.filter((task) => !task.completed));
    };

    return (
        <>
            <Header/>
            <div className={styles.taskContainer}>
                <div className={styles.dialog} style={{display: tasks.length === 0 ? "block" : "none"}}>
                    The task list is empty
                </div>
                <ul>
                    {tasks.map((task, index) => (
                        <li key={index} className={task.completed ? styles.completed : ""}>
                            <input
                                type="checkbox"
                                checked={task.completed}
                                onChange={() => handleTaskCompletion(index)}
                                className={styles.checkbox}
                            />
                            {task.text}
                            <button
                                onClick={() => setTasks(tasks.filter((_, i) => i !== index))}
                            >
                                ❌
                            </button>
                        </li>
                    ))}
                </ul>
                <input
                    type="text"
                    placeholder="Enter a new task..."
                    value={newTask}
                    onChange={(e) => setNewTask(e.target.value)}
                />
                <button className={styles.addTask} onClick={handleAddTask}>
                    Add Task
                </button>
                <button
                    className={styles.deleteCompleted}
                    onClick={handleDeleteCompletedTasks}
                >
                    Delete Completed Tasks
                </button>
            </div>
        </>

    );
};

export default Todo;
