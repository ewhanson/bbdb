import { MainComponentWrapper } from "../MainComponentWrapper.jsx";
import { useUploaderAuthProtected } from "../../lib/customHooks.js";
import { PhotoUploader } from "../uploader/PhotoUploader.jsx";

export function UploaderDashboard() {
  const isUploader = useUploaderAuthProtected();

  if (!isUploader) return null;

  return (
    <MainComponentWrapper useFooter={true}>
      <PhotoUploader />
    </MainComponentWrapper>
  );
}
