import { Route, Routes } from "react-router-dom";
import RegisterPage from "./components/reg.tsx";
import Todo from "./components/todo.tsx";

function App() {
    return (
        <Routes>
            <Route path="/login" element={<RegisterPage />} />
            <Route path="/reg" element={<RegisterPage/>}/>
            <Route path="/todo" element={<Todo/>}/>
        </Routes>
    );
}

export default App;
