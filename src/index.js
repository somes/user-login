import React, {StrictMode} from 'react';
import ReactDOM from 'react-dom/client';
import './index.scss';
import router from "./router";
import {RouterProvider} from 'react-router-dom';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <StrictMode>
    <RouterProvider router={router}/>
  </StrictMode>
);
