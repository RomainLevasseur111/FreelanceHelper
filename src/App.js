import "./App.css";
import { Routes, Route } from "react-router-dom";

function App() {
    return (
        <div className="App">
            <header className="App-header">
                <RegisterForm />
            </header>
        </div>
    );
}

function RegisterForm() {
    return (
        <form action="" method="POST">
            <input type="text" name="name" placeholder="Name" />
            <br></br>
            <input type="email" name="mail" placeholder="Email" />
            <br></br>
            <input type="password" name="password" placeholder="Password" />
            <br></br>
            <button name="register" type="submit">
                Register me please!
            </button>
        </form>
    );
}

export default App;
