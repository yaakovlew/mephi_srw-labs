import React, {useEffect} from "react";
import {BrowserRouter as Router, Route, Routes} from "react-router-dom";

import Lab1aPage from "../pages/Lab1aPage.tsx";
import Lab1bPage from "../pages/Lab1bPage.tsx";
import {useDispatch} from "react-redux";
import {setToken} from "../../reducers/authSlice.ts";

const App: React.FC = () => {
    const dispatch = useDispatch();

    useEffect(() => {
        const url = new URL(window.location.href);
        const jwt = url.searchParams.get('jwt');

        if (jwt) {
            const checkIsNeedClearStorage = localStorage.getItem('isDone');
            if (checkIsNeedClearStorage === "1") {
                localStorage.clear();
            }
            localStorage.setItem('token', jwt);
            dispatch(setToken(jwt));

            url.searchParams.delete('jwt');
            window.history.replaceState({}, document.title, url.pathname + url.search);
            window.location.reload();
        }
    }, []);

    return (
        <Router>
            <Routes>
                <Route path={"/"} element={<Lab1bPage/>}/>
                <Route path={"/lab1a"} element={<Lab1aPage/>}/>
                <Route path={"/lab1b"} element={<Lab1bPage/>}/>
            </Routes>
        </Router>

    );
};

export default App;