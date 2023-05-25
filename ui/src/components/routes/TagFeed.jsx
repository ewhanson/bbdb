import { PhotoFeed } from "../photoFeed/PhotoFeed.jsx";
import { useGetPhotos, useViewerAuthProtected } from "../../lib/customHooks.js";
import { MainComponentWrapper } from "../MainComponentWrapper.jsx";
import { getPhotosByTag } from "../../lib/pocketbase.js";

export function TagFeed({ tagName }) {
  const isViewer = useViewerAuthProtected();
  if (!isViewer) return null;

  const getPhotosCallback = async (page, perPage) => {
    return getPhotosByTag(tagName, page, perPage);
  };
  const [photos, errorMessage, isLastPage, isFetching, setIsFetching] =
    useGetPhotos(getPhotosCallback);

  return (
    <MainComponentWrapper>
      <h1 className="text-2xl font-bold mb-4">#{tagName}</h1>
      <PhotoFeed
        photos={photos}
        errorMessage={errorMessage}
        isLastPage={isLastPage}
        isFetching={isFetching}
        setIsFetching={setIsFetching}
      />
    </MainComponentWrapper>
  );
}
