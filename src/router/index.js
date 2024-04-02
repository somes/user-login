import {createBrowserRouter} from 'react-router-dom';


const router = createBrowserRouter([
  {
    path: '*',
    element: <div style={{textAlign: 'center'}}><h1>404 Not Found</h1></div>
  }
])

export default router;
