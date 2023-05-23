import { PhotoCard } from "./PhotoCard.jsx";
import { LoadingSpinner } from "../LoadingSpinner.jsx";
import { constants } from "../../lib/constants.js";
import { useEffect, useState } from "preact/hooks";

import { useInfiniteScroll } from "../../lib/customHooks.js";
import { getPhotos } from "../../lib/pocketbase.js";
import { getUniqueArrayBy } from "../../lib/helpers.js";

export function PhotoFeed() {
  const [page, setPage] = useState(0);
  const [sortOrder, setSortOrder] = useState(constants.SORT_ORDER.DATE_TAKEN);

  const [photos, setPhotos] = useState([]);
  const [errorMessage, setErrorMessage] = useState("");
  const [isLastPage, setIsLastPage] = useState(false);

  // set isFetching to trigger loading photos on page load
  useEffect(() => {
    setIsFetching(true);
  }, []);

  // Reset page and re-trigger fetch when changing sort order
  useEffect(() => {
    if (page === 0) return;
    setPage(0);
    setPhotos([]);
    setIsFetching(true);
  }, [sortOrder]);

  const [isFetching, setIsFetching] = useInfiniteScroll(async () => {
    if (isLastPage) return;

    const pageToCheck = page + 1;
    try {
      const results = await getPhotos(pageToCheck, sortOrder);
      const combinedPhotos = [...photos].concat(results.photoData);
      const sanitizedCombinedPhotos = getUniqueArrayBy(combinedPhotos, "id");

      setIsLastPage(results.page >= results.maxPages);

      setPage(results.page);
      setErrorMessage("");
      setPhotos(sanitizedCombinedPhotos);
    } catch (e) {
      setErrorMessage(e.message);
    } finally {
      setIsFetching(false);
    }
  }, isLastPage);

  if (errorMessage !== "") {
    return (
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
    );
  }

  return (
    <>
      <div
        className={
          "card card-compact bg-base-100 w-full sm:w-auto sm:max-w-md sm:min-w-28rem shadow-sm"
        }
      >
        <div className={"card-body filters-container"}>
          <div className={"collapse collapse-arrow"}>
            <input type={"checkbox"} />
            <h2 className={"collapse-title card-title"}>Filters</h2>
            <div
              className={"card-actions collapse-content filters-select-area"}
            >
              <div className={"flex flex-col"}>
                <p className={"pb-4"}>Sort photos by:</p>
                <select
                  className={
                    "select select-bordered select-xs w-full max-w-xs ml-1 filters-select"
                  }
                  onInput={(e) => setSortOrder(e.target.value)}
                  value={sortOrder}
                >
                  <option value={constants.SORT_ORDER.DATE_TAKEN}>
                    Date taken
                  </option>
                  <option value={constants.SORT_ORDER.CREATED}>
                    Date added
                  </option>
                </select>
              </div>
            </div>
          </div>
        </div>
      </div>
      {photos.map((photo) => {
        return (
          <PhotoCard
            key={photo.id}
            url={photo.url}
            displayDate={photo.displayDate}
            altDate={photo.altDate}
            description={photo.description}
            isNew={photo.isNew}
          />
        );
      })}
      {!isLastPage && !isFetching && (
        <button
          onClick={() => setIsFetching(true)}
          className="btn btn-outline btn-sm"
        >
          Load more photos
        </button>
      )}
      {isFetching && !isLastPage && <LoadingSpinner />}
    </>
  );
}
