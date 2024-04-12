import {createBrowserRouter} from 'react-router-dom';

import Login from '../pages/Login';

const router = createBrowserRouter([
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
