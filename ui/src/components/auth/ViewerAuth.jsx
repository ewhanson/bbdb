import { useContext, useState } from "preact/hooks";
import {
  isUploaderLoggedIn,
  isViewerLoggedIn,
  viewerLogin,
} from "../../lib/pocketbase.js";
import { AuthContext } from "../../lib/AuthContextProvider.js";
import { route } from "preact-router";
import { constants } from "../../lib/constants.js";

export function ViewerAuth() {
  const [authData, setAuthData] = useContext(AuthContext);

  const [password, setPassword] = useState("");
  const [errorMessage, setErrorMessage] = useState("");
  const [isSubmitting, setIsSubmitting] = useState(false);

  async function onSubmit(event) {
    event.preventDefault();
    setErrorMessage("");
    setIsSubmitting(true);
    try {
      await viewerLogin(password);
      setAuthData({
        isViewer: isViewerLoggedIn(),
        isUploader: isUploaderLoggedIn(),
      });
      setIsSubmitting(false);
      route(constants.ROUTES.FEED);
    } catch (e) {
      console.error({ loginError: e });
      setIsSubmitting(false);
      setErrorMessage(e.message);
    }
  }

  return (
    <>
      {errorMessage.length !== 0 && (
        <div className="alert alert-error max-w-xl shadow-lg">
          <div>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              className="stroke-current flex-shrink-0 h-6 w-6"
              fill="none"
              viewBox="0 0 24 24"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth="2"
                d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z"
              />
            </svg>
            <span>{errorMessage}</span>
          </div>
        </div>
      )}
      <div className="card bg-base-100 shadow-xl w-auto sm:w-96">
        <form className="card-body" onSubmit={onSubmit}>
          <h2 className="card-title">Babygramz Access</h2>
          <div className="form-control w-full max-w-md">
            <label className="label">
              <span className="label-text">
                Please enter the password you were given for access.{" "}
              </span>
            </label>
            <input
              value={password}
              onInput={(e) => {
                setPassword(e.target.value);
              }}
              type="password"
              placeholder="Password"
              className="input input-bordered w-full max-w-md"
            />
            <label className="label">
              <span className="label-text-alt">
                <a className="link" href={constants.ROUTES.UPLOADER.LOGIN}>
                  Access uploader login
                </a>
              </span>
            </label>
          </div>
          <div className="card-actions justify-end">
            <button
              type={"submit"}
              className={`btn btn-primary ${isSubmitting ? "loading" : ""}`}
            >
              Submit
            </button>
          </div>
        </form>
      </div>
    </>
  );
}
