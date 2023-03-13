import { PhotoFeed } from "../photoFeed/PhotoFeed.jsx";
import { useViewerAuthProtected } from "../../lib/customHooks.js";
import { MainComponentWrapper } from "../MainComponentWrapper.jsx";

export function Feed() {
  const isViewer = useViewerAuthProtected();
  if (!isViewer) return null;

  return (
    <MainComponentWrapper>
      <PhotoFeed />
    </MainComponentWrapper>
  );
}
