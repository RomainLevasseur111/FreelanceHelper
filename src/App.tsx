import { Route, Routes } from "react-router-dom";
import { useEffect, useState } from "react";

function App() {
    return (
        <Routes>
            <Route path="/" element={<Root />} />
        </Routes>
    );
}

function Root() {
    const [data, setData] = useState<string>("");

    useEffect(() => {
        fetch("/test")
            .then((resp) => (resp.ok ? resp.json() : ""))
            .then(setData);
    }, []);

    return (
        <>
            <p>FreelanceHelper</p>
            <div>{data}</div>
        </>
    );
}

export default App;
