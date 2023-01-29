import { constants } from "../../lib/constants.js";
import { useContext, useEffect } from "preact/hooks";
import { AuthContext } from "../../lib/AuthContextProvider.js";
import { route } from "preact-router";
import { Footer } from "../Footer.jsx";

export function Home() {
  const [isValid] = useContext(AuthContext);

  useEffect(() => {
    if (isValid) {
      route(constants.ROUTES.FEED, true);
    }
  }, []);

  return (
    <>
      <main className="hero bg-base-200" style="min-height: 92.5vh">
        <div className="hero-content text-center">
          <div className="max-w-md">
            <h1 className="text-4xl font-bold">Welcome to Babygramz!</h1>
            <p className="py-6">
              A personal project for sharing baby photos with friends and
              family. Already have an access code? Login by clicking the button
              below!
            </p>
            <a href={constants.ROUTES.LOGIN} className="btn btn-primary">
              Start viewing
            </a>
          </div>
        </div>
      </main>
      <Footer />
    </>
  );
}
