import { isUserLoggedIn, logOut } from "../lib/pocketbase.js";
import { useContext } from "preact/hooks";
import { AuthContext } from "../lib/AuthContextProvider.js";
import { route } from "preact-router";
import { constants } from "../lib/constants.js";
import { Icon } from "./Icon.jsx";

export function Navbar() {
  const [isValid, setIsValid] = useContext(AuthContext);

  const doLogout = () => {
    if (isValid) {
      logOut();
      setIsValid(isUserLoggedIn());
      return route(constants.ROUTES.HOME);
    }
  };

  return (
    <div className="navbar bg-base-100">
      <div className="flex-1">
        <a
          href={isValid ? constants.ROUTES.FEED : constants.ROUTES.HOME}
          className="btn btn-ghost normal-case text-xl"
        >
          Babygramz👶🎆
        </a>
      </div>
      <div className="flex-none">
        <div className="dropdown dropdown-end">
          <label tabIndex="0" className="btn btn-square btn-ghost">
            <Icon name={constants.ICONS.DOTS_HORIZONTAL} />
          </label>
          <ul
            tabIndex="0"
            className="mt-3 p-2 shadow menu menu-compact dropdown-content bg-base-100 rounded-box w-52"
          >
            {isValid && (
              <li>
                <a href={constants.ROUTES.FEED}>Photo feed</a>
              </li>
            )}
            <li>
              <a href={constants.ROUTES.ABOUT}>About</a>
            </li>
            {isValid && (
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
              {isValid ? (
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
