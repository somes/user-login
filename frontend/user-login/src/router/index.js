import {createBrowserRouter} from 'react-router-dom';

import {AuthRoute} from "@/components/AuthRoute";

import Index from "@/pages/Index";
import Login from '@/pages/Login';
import Home from "@/pages/Home";

const router = createBrowserRouter([
  {
    path: '/',
    element: <AuthRoute><Index/></AuthRoute>,
    children: [
      {
        index: true,
        element: <Home/>
      },
      // {
      //   path: '/article',
      //   element: <Article/>
      // },
      // {
      //   path: '/publish',
      //   element: <Publish/>
      // },
    ]
  },
  {
    path: '/login',
    element: <Login/>
  },
  {
    path: '*',
    element: <div style={{textAlign: 'center'}}><h1>404 Not Found</h1></div>
  }
])

export default router;
