import {
  getHasNewPhotos,
  isUploaderLoggedIn,
  isViewerLoggedIn,
  logOut,
} from "../lib/pocketbase.js";
import { useContext, useEffect, useState } from "preact/hooks";
import { AuthContext } from "../lib/AuthContextProvider.js";
import { route } from "preact-router";
import { constants } from "../lib/constants.js";
import { Icon } from "./Icon.jsx";

export function Navbar() {
  const [authData, setAuthData] = useContext(AuthContext);
  const [hasNewPhotos, setHasNewPhotos] = useState(false);

  useEffect(() => {
    getHasNewPhotos().then((res) => setHasNewPhotos(res));
  }, []);

  const doLogout = () => {
    // TODO: See if this should apply not just to viewer
    if (authData.isViewer) {
      logOut();
      setAuthData({
        isViewer: isViewerLoggedIn(),
        isUploader: isUploaderLoggedIn(),
      });
      return route(constants.ROUTES.HOME);
    }
  };

  const isBuildOrderThanOneWeek = () => {
    const oneWeekAgo = Date.now() - 7 * 24 * 60 * 60 * 1000;
    const buildTime = APP_BUILD_DATE;
    return buildTime < oneWeekAgo;
  };

  const shouldDisplayUpdateBadge = () => {
    return !isBuildOrderThanOneWeek() || hasNewPhotos;
  };

  return (
    <div className="navbar bg-base-100">
      <div className="flex-1">
        <a
          href={authData ? constants.ROUTES.FEED : constants.ROUTES.HOME}
          className="btn btn-ghost normal-case text-xl"
        >
          Babygramz👶🎆
        </a>
      </div>
      <div className="flex-none">
        <div className="dropdown dropdown-end">
          <div className="indicator">
            {shouldDisplayUpdateBadge && (
              <div className="badge badge-secondary badge-xs indicator-item mt-1 mr-1"></div>
            )}
            <label tabIndex="0" className="btn btn-square btn-ghost">
              <div>
                <Icon name={constants.ICONS.DOTS_HORIZONTAL} />
              </div>
            </label>
          </div>
          <ul
            tabIndex="0"
            className="mt-3 p-2 shadow menu menu-compact dropdown-content bg-base-100 rounded-box w-52"
          >
            {authData.isViewer && (
              <li>
                <a className={"justify-between"} href={constants.ROUTES.FEED}>
                  Photo feed
                  {hasNewPhotos && (
                    <span className="badge badge-secondary badge-sm">
                      updated
                    </span>
                  )}
                </a>
              </li>
            )}
            <li>
              <a href={constants.ROUTES.ABOUT}>About</a>
            </li>
            <li>
              <a
                href={constants.ROUTES.WHATS_NEW}
                className={"justify-between"}
              >
                What's new
                {!isBuildOrderThanOneWeek() && (
                  <span className="badge badge-secondary badge-sm">
                    updated
                  </span>
                )}
              </a>
            </li>
            {authData.isViewer && (
              <li>
                <a href={constants.ROUTES.NOTIFICATIONS}>
                  Notifications signup
                </a>
              </li>
            )}
            <li>
              {authData.isViewer ? (
                <button onClick={doLogout}>Logout</button>
              ) : (
                <a href={constants.ROUTES.LOGIN}>Login</a>
              )}
            </li>
          </ul>
        </div>
      </div>
    </div>
  );
}
