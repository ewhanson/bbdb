import { MainComponentWrapper } from "../MainComponentWrapper.jsx";
import { useEffect, useState } from "preact/hooks";
import { LoadingSpinner } from "../LoadingSpinner.jsx";

export function Unsubscribe({ id }) {
  const [isLoading, setIsLoading] = useState(false);
  const [successMessage, setSuccessMessage] = useState("");
  const [errorMessage, setErrorMessage] = useState("");

  useEffect(async () => {
    setIsLoading(true);
    const res = await fetch(
      `${import.meta.env.VITE_POCKETBASE_URL}/api/unsubscribe?id=${id}`,
      { method: "POST" }
    );

    if (res.status === 200) {
      setErrorMessage("");
      setSuccessMessage(
        "You have successfully unsubscribed from Babygramz notifications."
      );
    } else {
      setSuccessMessage("");
      setErrorMessage(
        "An error occurred while attempting to complete your request. Please try again later."
      );
    }

    setIsLoading(false);
  }, []);

  return (
    <MainComponentWrapper>
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
      {isLoading && <LoadingSpinner />}
    </MainComponentWrapper>
  );
}
