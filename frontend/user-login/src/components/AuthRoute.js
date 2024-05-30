import {Navigate} from "react-router-dom";
import {useSelector} from "react-redux";

export function AuthRoute({children}) {
  // const token = _getToken();
  const token = useSelector(state => state.user.token);
  if (!token) {
    return <Navigate to={'/login'} replace/>;
  } else {
    return children;
  }
}
