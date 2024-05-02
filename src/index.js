import React, {StrictMode} from 'react';
import ReactDOM from 'react-dom/client';
import {RouterProvider} from 'react-router-dom';
import {Provider} from "react-redux";

import './index.scss';
import 'normalize.css';

import router from "@/router";
import store from "@/store";

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <StrictMode>
    <Provider store={store}>
      <RouterProvider router={router}/>
    </Provider>
  </StrictMode>
);
