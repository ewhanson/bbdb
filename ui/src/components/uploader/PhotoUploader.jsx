import { useUploaderAuthProtected } from "../../lib/customHooks.js";
import { useEffect, useState } from "preact/hooks";
import { getTagIdsFromString, postPhoto } from "../../lib/pocketbase.js";
import { tryGetDateTimeFromImage } from "../../lib/helpers.js";

export function PhotoUploader() {
  const isUploader = useUploaderAuthProtected();

  const [description, setDescription] = useState("");
  const [file, setFile] = useState(null);
  const [tags, setTags] = useState("");
  const [dateTime, setDateTime] = useState("");
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [errorMessage, setErrorMessage] = useState("");
  const [successMessage, setSuccessMessage] = useState("");

  const onSubmit = async (event) => {
    event.preventDefault();
    setIsSubmitting(true);
    setErrorMessage("");
    setSuccessMessage("");

    try {
      await postPhoto(description, file, dateTime, tags);
      setIsSubmitting(false);
      setSuccessMessage("Photo upload successful!");
      setDescription("");
      setFile(null);
    } catch (e) {
      console.error({ error: e.message });
      setIsSubmitting(false);
      setErrorMessage(e.message);
      setSuccessMessage("");
    }
  };

  if (!isUploader) return null;

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
        <form className="card-body" onSubmit={onSubmit}>
          <h2 className="card-title">Upload a photo</h2>
          <p>
            Use this form to upload a new photo.
            <br />
            The date taken will be inferred from the photo directly.
          </p>

          {/* Title */}
          <div className="form-control w-full max-w-lg">
            <label className="label">
              <span className="label-text">Description</span>
            </label>
            <input
              type="text"
              value={description}
              onInput={(e) => setDescription(e.target.value)}
              placeholder="Enter a description"
              className="input input-bordered w-full"
              required
            />
          </div>

          {/* File */}
          <div className="form-control w-full max-w-sm"></div>
          <label className="label">
            <span className="label-text">Photo</span>
          </label>
          <input
            type="file"
            value={file}
            onInput={async (e) => {
              setFile(e.target?.files[0]);
              const currentDateTime = await tryGetDateTimeFromImage(
                e.target?.files[0]
              );
              if (currentDateTime !== null) {
                setDateTime(currentDateTime);
              }
            }}
            className="file-input file-input-bordered w-full"
          />

          {/* Tags */}
          <div className="form-control w-full max-w-lg">
            <label className="label">
              <span className="label-text">Tags</span>
            </label>
            <input
              type="text"
              value={tags}
              onInput={(e) => setTags(e.target.value)}
              placeholder="Enter a comma-separated list"
              className="input input-bordered w-full"
              required
            />
          </div>

          <div className="card-actions justify-end pt-4">
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
