import { MainComponentWrapper } from "../MainComponentWrapper.jsx";
import { UploaderAuth } from "../auth/UploaderAuth.jsx";
import { useUploaderAuthProtected } from "../../lib/customHooks.js";
import { route } from "preact-router";
import { constants } from "../../lib/constants.js";
import { useEffect } from "preact/hooks";

export function UploaderLogin() {
  const isUploader = useUploaderAuthProtected();

  useEffect(() => {
    if (isUploader) {
      route(constants.ROUTES.UPLOADER.DASHBOARD, true);
    }
  }, [isUploader]);

  if (isUploader) {
    return null;
  }

  return (
    <MainComponentWrapper useFooter={true}>
      <UploaderAuth />
    </MainComponentWrapper>
  );
}
