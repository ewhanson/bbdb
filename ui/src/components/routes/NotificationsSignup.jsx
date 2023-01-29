import { MainComponentWrapper } from "../MainComponentWrapper.jsx";
import { useAuthProtected } from "../../lib/customHooks.js";
import { useState } from "preact/hooks";
import { signupForNotifications } from "../../lib/pocketbase.js";

export function NotificationsSignup() {
  const isValid = useAuthProtected();

  const [email, setEmail] = useState("");
  const [errorMessage, setErrorMessage] = useState("");
  const [successMessage, setSuccessMessage] = useState("");
  const [isSubmitting, setIsSubmitting] = useState(false);

  const onSubmit = async (event) => {
    event.preventDefault();
    setErrorMessage("");
    setSuccessMessage("");
    setIsSubmitting(true);

    try {
      // Function here
      await signupForNotifications(email);
      setIsSubmitting(false);
      setSuccessMessage("Sign up successful! Check your email for details.");
      setEmail("");
    } catch (e) {
      console.error({ notificationSignupError: e });
      setIsSubmitting(false);
      setErrorMessage(e.message);
    }
  };

  if (!isValid) return null;

  return (
    <MainComponentWrapper useFooter={true}>
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

      {successMessage.length !== 0 && (
        <div className="alert alert-success max-w-xl shadow-lg">
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
                d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"
              />
            </svg>
            <span>{successMessage}</span>
          </div>
        </div>
      )}

      <div className="card bg-base-100 shadow-xl w-auto">
        <form className={"card-body"} onSubmit={onSubmit}>
          <h2 className={"card-title"}>Notifications Signup</h2>
          <div className={"form-control w-full max-w-lg"}>
            <label className={"label"}>
              <span className={"label-text"}>
                Sign up to receive an email when there are new photos. We'll
                email you at most once a day.
              </span>
            </label>
            <input
              type={"email"}
              value={email}
              onInput={(e) => setEmail(e.target.value)}
              placeholder={"Enter your email"}
              className={"input input-bordered w-full"}
              required
            />
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
    </MainComponentWrapper>
  );
}
