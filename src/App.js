import "./App.css";
import { Routes, Route } from "react-router-dom";
import RegisterForm from "./register";
import React, { useEffect, useState } from "react";

function App() {
    return (
        <div className="App">
            <header className="App-header">
                <div>
                    <RegisterForm />
                </div>
            </header>
        </div>
    );
}

export default App;
