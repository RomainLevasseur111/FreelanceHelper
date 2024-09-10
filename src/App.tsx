import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import { useState } from "react";

function App() {
    return (
        <Router>
            <Routes>
                <Route path="/" element={<Root />} />
            </Routes>
        </Router>
    );
}

function Root() {
    const [name, setName] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");

    const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();

        const requestOptions = {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({
                name: name,
                email: email,
                password: password,
            }),
        };

        fetch("http://localhost:8080/register", requestOptions)
            .then((response) => {
                if (!response.ok) {
                    throw new Error("Failed to register");
                }
                return response.json();
            })
            .then((data) => {
                console.log("Registration successful:", data);
            })
            .catch((error) => {
                console.error("Error during registration:", error);
            });
    };

    return (
        <div>
            <h2>Formulaire d'inscription</h2>
            <form onSubmit={handleSubmit}>
                <div>
                    <label htmlFor="name">Nom :</label>
                    <input
                        type="text"
                        id="name"
                        value={name}
                        onChange={(e) => setName(e.target.value)}
                        required
                    />
                </div>
                <div>
                    <label htmlFor="email">Email :</label>
                    <input
                        type="email"
                        id="email"
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                        required
                    />
                </div>
                <div>
                    <label htmlFor="password">Mot de passe :</label>
                    <input
                        type="password"
                        id="password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        required
                    />
                </div>
                <button type="submit">S'inscrire</button>
            </form>
        </div>
    );
}

export default App;
