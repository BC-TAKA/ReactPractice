// import React from 'react';
// import logo from './logo.svg';
import './App.css';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import LogIn from './pages/LogIn';
import List from './pages/List';

function App() {
  // ページの定義をまとめる
  const pages = [
    {path: "/", element: <LogIn/>},
    {path: "/List", element: <List/>}
  ]

  // TODO: 都度Route宣言しなくても良いようにする
  const pageRoutes = pages.forEach((page) => {
    return (
      <Route path={page.path} element={page.element} />
    )
  })

  return (
    <BrowserRouter>
      <Routes>
        <Route path={`/`} element={<LogIn/>} />
        <Route path={`/List`} element={<List/>} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
