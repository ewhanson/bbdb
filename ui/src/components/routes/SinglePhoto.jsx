import { MainComponentWrapper } from "../MainComponentWrapper.jsx";
import { PhotoCard } from "../photoFeed/PhotoCard.jsx";
import { useGetPhoto, useViewerAuthProtected } from "../../lib/customHooks.js";
import { LoadingSpinner } from "../LoadingSpinner.jsx";

export function SinglePhoto({ photoId }) {
  const isViewer = useViewerAuthProtected();
  if (!isViewer) return null;

  const [photo, errorMessage, isFetching] = useGetPhoto(photoId);

  if (errorMessage !== "") {
    return (
      <MainComponentWrapper>
        <div className="alert alert-error shadow-lg max-w-lg">
          <div>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              className="stroke-current flex-shrink-0 h-6 w-6"
              fill="none"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z"
              />
            </svg>
            <span>Oops... {errorMessage}</span>
          </div>
        </div>
      </MainComponentWrapper>
    );
  }

  return (
    <MainComponentWrapper>
      {Object.entries(photo).length !== 0 && (
        <PhotoCard
          id={photo.id}
          url={photo.url}
          displayDate={photo.displayDate}
          altDate={photo.altDate}
          description={photo.description}
          isNew={photo.isNew}
          tags={photo.tags}
        />
      )}
      {isFetching && <LoadingSpinner />}
    </MainComponentWrapper>
  );
}
