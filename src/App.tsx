import { useEffect, useState } from "react";

function App() {
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
