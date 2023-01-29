import { AuthContext } from "./lib/AuthContextProvider.js";
import { Navbar } from "./components/Navbar.jsx";
import { useState } from "preact/hooks";
import { isUserLoggedIn } from "./lib/pocketbase.js";
import { Router } from "preact-router";
import { Home } from "./components/routes/Home.jsx";
import { Login } from "./components/routes/Login.jsx";
import { Feed } from "./components/routes/Feed.jsx";
import { constants } from "./lib/constants.js";
import { About } from "./components/routes/About.jsx";
import { FourOhFour } from "./components/routes/FourOhFour.jsx";
import { NotificationsSignup } from "./components/routes/NotificationsSignup.jsx";
import { Unsubscribe } from "./components/routes/Unsubscribe.jsx";

export function App() {
  const [isValid, setIsValid] = useState(isUserLoggedIn);

  return (
    <AuthContext.Provider value={[isValid, setIsValid]}>
      <Navbar />
      <Router
        onChange={() => {
          document.activeElement.blur();
        }}
      >
        <Home path={constants.ROUTES.HOME} />
        <Login path={constants.ROUTES.LOGIN} />
        <Feed path={constants.ROUTES.FEED} />
        <About path={constants.ROUTES.ABOUT} />
        <NotificationsSignup path={constants.ROUTES.NOTIFICATIONS} />
        <Unsubscribe path={constants.ROUTES.UNSUBSCRIBE} />
        <FourOhFour path={constants.ROUTES.FOUR_OH_FOUR} default />
      </Router>
    </AuthContext.Provider>
  );
}
