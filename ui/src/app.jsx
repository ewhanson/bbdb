import { AuthContext } from "./lib/AuthContextProvider.js";
import { Navbar } from "./components/Navbar.jsx";
import { useState } from "preact/hooks";
import { isUploaderLoggedIn, isViewerLoggedIn } from "./lib/pocketbase.js";
import { Router } from "preact-router";
import { Home } from "./components/routes/Home.jsx";
import { Login } from "./components/routes/Login.jsx";
import { Feed } from "./components/routes/Feed.jsx";
import { constants } from "./lib/constants.js";
import { About } from "./components/routes/About.jsx";
import { FourOhFour } from "./components/routes/FourOhFour.jsx";
import { NotificationsSignup } from "./components/routes/NotificationsSignup.jsx";
import { Unsubscribe } from "./components/routes/Unsubscribe.jsx";
import { UploaderLogin } from "./components/routes/UploaderLogin.jsx";
import { UploaderDashboard } from "./components/routes/UploaderDashboard.jsx";
import { TagFeed } from "./components/routes/TagFeed.jsx";
import { WhatsNew } from "./components/routes/WhatsNew.jsx";

export function App() {
  const [authData, setAuthData] = useState({
    isViewer: isViewerLoggedIn(),
    isUploader: isUploaderLoggedIn(),
  });

  return (
    <AuthContext.Provider value={[authData, setAuthData]}>
      <Navbar />
      <Router
        onChange={() => {
          document.activeElement.blur();
        }}
      >
        <Home path={constants.ROUTES.HOME} />
        <Login path={constants.ROUTES.LOGIN} />
        <Feed path={constants.ROUTES.FEED} />
        <TagFeed path={constants.ROUTES.TAG} />
        <About path={constants.ROUTES.ABOUT} />
        <WhatsNew path={constants.ROUTES.WHATS_NEW} />
        <NotificationsSignup path={constants.ROUTES.NOTIFICATIONS} />
        <Unsubscribe path={constants.ROUTES.UNSUBSCRIBE} />
        <UploaderLogin path={constants.ROUTES.UPLOADER.LOGIN} />
        <UploaderDashboard path={constants.ROUTES.UPLOADER.DASHBOARD} />
        <FourOhFour path={constants.ROUTES.FOUR_OH_FOUR} default />
      </Router>
    </AuthContext.Provider>
  );
}
