import { MainComponentWrapper } from "../MainComponentWrapper.jsx";

export function About() {
  return (
    <MainComponentWrapper useFooter={true}>
      <article className="prose">
        <h1>About Babygramz</h1>
        <p>
          Thanks for visiting! Babygramz is a personal project created to share
          baby photos with friends and family. It allows you to create a
          password-protected photo feed to share and store photos using the
          infrastructure you choose without added tracking and telemetry built
          in. For now, the project is limited to a single user/photo feed, but
          may be expanded, open-sourced in the future.
        </p>
        <p>
          Babygramz is built with open source technologies including{" "}
          <a href="https://preactjs.com/">Preact</a>,{" "}
          <a href="https://daisyui.com/">DaisyUI</a>, and{" "}
          <a href="https://pocketbase.io">PocketBase</a>.
        </p>
      </article>
    </MainComponentWrapper>
  );
}
