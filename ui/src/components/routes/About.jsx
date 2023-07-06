import { MainComponentWrapper } from "../MainComponentWrapper.jsx";

export function About() {
  return (
    <MainComponentWrapper useFooter={true}>
      <article className="prose">
        <h1>About Babygramz</h1>
        <p>
          Thanks for visiting! Babygramz is a digital scrapbook for sharing
          photos with friends and family. It allows you to create a
          password-protected photo feed with the ethos of a physical scrapbook
          and feel of a modern social media app. For now, the project is limited
          to a single user/photo feed, but I'd like to expand and open-source it
          in the future.
        </p>
        <p>
          Built with ♥️ and <a href="https://preactjs.com/">Preact</a>,{" "}
          <a href="https://daisyui.com/">DaisyUI</a>, and{" "}
          <a href="https://pocketbase.io">PocketBase</a>.
        </p>
      </article>
    </MainComponentWrapper>
  );
}
