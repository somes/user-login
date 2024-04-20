import React, {StrictMode} from 'react';
import ReactDOM from 'react-dom/client';
import './index.scss';
import router from "./router";
import {RouterProvider} from 'react-router-dom';
import store from "./store";
import {Provider} from "react-redux";

import 'normalize.css';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <StrictMode>
    <Provider store={store}>
      <RouterProvider router={router}/>
    </Provider>
  </StrictMode>
);
