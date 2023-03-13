import { useContext, useState } from "preact/hooks";
import { AuthContext } from "../../lib/AuthContextProvider.js";
import {
  isUploaderLoggedIn,
  isViewerLoggedIn,
  login,
} from "../../lib/pocketbase.js";
import { route } from "preact-router";
import { constants } from "../../lib/constants.js";

export function UploaderAuth() {
  const [authData, setAuthData] = useContext(AuthContext);

  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [errorMessage, setErrorMessage] = useState("");
  const [isSubmitting, setIsSubmitting] = useState(false);

  const onSubmit = async (event) => {
    event.preventDefault();
    setIsSubmitting(false);
    setErrorMessage("");

    try {
      await login(email, password);
      setAuthData({
        isViewer: isViewerLoggedIn(),
        isUploader: isUploaderLoggedIn(),
      });
      setIsSubmitting(false);
      route(constants.ROUTES.UPLOADER.DASHBOARD);
    } catch (e) {
      console.error({ loginError: e.message });
      setIsSubmitting(false);
      setErrorMessage(e.message);
    }
  };

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

      <div className="card bg-base-100 shadow-xl w-auto">
        <form className="card-body" onSubmit={onSubmit}>
          <h2 className="card-title">Uploader Login</h2>
          <p>
            Users with uploading privileges can login here to upload new photos.
          </p>

          {/* Email */}
          <div className="form-control w-full max-w-lg">
            <label className="label">
              <span className="label-text">Email</span>
            </label>
            <input
              type="email"
              value={email}
              onInput={(e) => setEmail(e.target.value)}
              placeholder="Enter your email"
              className="input input-bordered w-full"
              required
            />
          </div>

          {/* Password */}
          <div className="form-control w-full max-w-lg">
            <label className="label">
              <span className="label-text">Password</span>
            </label>
            <input
              type="password"
              value={password}
              onInput={(e) => setPassword(e.target.value)}
              placeholder="Enter your password"
              className="input input-bordered w-full"
              required
            />
          </div>

          <div className="card-actions justify-end">
            <button
              type="Submit"
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
