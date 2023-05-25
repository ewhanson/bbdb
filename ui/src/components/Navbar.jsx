import {
  isUploaderLoggedIn,
  isViewerLoggedIn,
  logOut,
} from "../lib/pocketbase.js";
import { useContext } from "preact/hooks";
import { AuthContext } from "../lib/AuthContextProvider.js";
import { route } from "preact-router";
import { constants } from "../lib/constants.js";
import { Icon } from "./Icon.jsx";

export function Navbar() {
  const [authData, setAuthData] = useContext(AuthContext);

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
          <label tabIndex="0" className="btn btn-square btn-ghost">
            <Icon name={constants.ICONS.DOTS_HORIZONTAL} />
            {/*<div className="badge badge-accent badge-xs ml-0.5 self-start"></div>*/}
          </label>
          <ul
            tabIndex="0"
            className="mt-3 p-2 shadow menu menu-compact dropdown-content bg-base-100 rounded-box w-52"
          >
            {authData.isViewer && (
              <li>
                <a href={constants.ROUTES.FEED}>Photo feed</a>
              </li>
            )}
            <li>
              <a href={constants.ROUTES.ABOUT}>About</a>
            </li>
            {authData.isViewer && (
              <li>
                <a
                  className={"justify-between"}
                  href={constants.ROUTES.NOTIFICATIONS}
                >
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
