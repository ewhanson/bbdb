import { MainComponentWrapper } from "../MainComponentWrapper.jsx";
import { UploaderAuth } from "../auth/UploaderAuth.jsx";

export function UploaderLogin() {
  return (
    <MainComponentWrapper useFooter={true}>
      <UploaderAuth />
    </MainComponentWrapper>
  );
}
