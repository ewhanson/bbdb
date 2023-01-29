import { PhotoFeed } from "../photoFeed/PhotoFeed.jsx";
import { useAuthProtected } from "../../lib/customHooks.js";
import { MainComponentWrapper } from "../MainComponentWrapper.jsx";

export function Feed() {
  const isValid = useAuthProtected();
  if (!isValid) return null;

  return (
    <MainComponentWrapper>
      <PhotoFeed />
    </MainComponentWrapper>
  );
}
