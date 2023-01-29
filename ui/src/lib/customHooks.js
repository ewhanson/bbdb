import { useContext, useEffect, useState } from "preact/hooks";
import { AuthContext } from "./AuthContextProvider.js";
import { route } from "preact-router";
import { constants } from "./constants.js";
import Cookies from "cookies-js";

export const useAuthProtected = () => {
  const [isValid] = useContext(AuthContext);

  useEffect(() => {
    if (isValid) return;

    route(constants.ROUTES.LOGIN, true);
  }, [isValid]);

  return isValid;
};

export const useGetCookie = (key) => {
  const [cookieValue, setCookieValue] = useState("");
  useEffect(() => {
    setCookieValue(Cookies.get(key));
  }, key);

  return cookieValue;
};

export const useInfiniteScroll = (callback, shouldStopExecution) => {
  const [isFetching, setIsFetching] = useState(false);

  useEffect(() => {
    window.addEventListener("scroll", handleScroll);
    return () => window.removeEventListener("scroll", handleScroll);
  }, []);

  useEffect(() => {
    const runCallback = async () => {
      if (!isFetching) return;
      await callback();
    };

    if (shouldStopExecution) {
      setIsFetching(false);
      return;
    }
    runCallback().catch(console.error);
  }, [isFetching]);

  function handleScroll() {
    const current = window.innerHeight + document.documentElement.scrollTop;
    const max = document.documentElement.offsetHeight;
    const isAtBottom = current / max > 0.8;
    if (!isAtBottom || isFetching) return;
    setIsFetching(true);
  }

  return [isFetching, setIsFetching];
};
