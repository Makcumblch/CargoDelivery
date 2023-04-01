import { BrowserRouter } from 'react-router-dom'

import useRoutes from "./hooks/useRoutes";
import { useAuth } from './hooks/useAuth';
import { AuthContext } from './contexts/AuthContext';

function App() {
  const auth = useAuth()
  const isAuthenticated = !!auth.token
  const routes = useRoutes(isAuthenticated)
  return (
    <AuthContext.Provider value={{...auth, isAuthenticated}}>
      <BrowserRouter>
        {routes}
      </BrowserRouter>
    </AuthContext.Provider>
  );
}

export default App;
