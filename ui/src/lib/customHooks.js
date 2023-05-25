import { useContext, useEffect, useState } from "preact/hooks";
import { AuthContext } from "./AuthContextProvider.js";
import { route } from "preact-router";
import { constants } from "./constants.js";
import Cookies from "cookies-js";
import { getUniqueArrayBy } from "./helpers.js";

export const useViewerAuthProtected = () => {
  const [authData] = useContext(AuthContext);

  useEffect(() => {
    if (authData.isViewer) return;

    route(constants.ROUTES.LOGIN, true);
  }, [authData]);

  return authData.isViewer;
};

export const useUploaderAuthProtected = () => {
  const [authData] = useContext(AuthContext);

  useEffect(() => {
    if (authData.isUploader) return;

    route(constants.ROUTES.UPLOADER.LOGIN, true);
  }, [authData]);

  return authData.isUploader;
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

export const useGetPhotos = (getPhotos) => {
  const [page, setPage] = useState(0);

  const [photos, setPhotos] = useState([]);
  const [errorMessage, setErrorMessage] = useState("");
  const [isLastPage, setIsLastPage] = useState(false);

  // set isFetching to trigger loading photos on page load
  useEffect(() => {
    setIsFetching(true);
  }, []);

  const [isFetching, setIsFetching] = useInfiniteScroll(async () => {
    if (isLastPage) return;

    const pageToCheck = page + 1;
    try {
      const results = await getPhotos(pageToCheck);
      const combinedPhotos = [...photos].concat(results.photoData);
      const sanitizedCombinedPhotos = getUniqueArrayBy(combinedPhotos, "id");

      setIsLastPage(results.page >= results.maxPages);

      setPage(results.page);
      setErrorMessage("");
      setPhotos(sanitizedCombinedPhotos);
    } catch (e) {
      setErrorMessage(e.message);
    } finally {
      setIsFetching(false);
    }
  }, isLastPage);

  return [photos, errorMessage, isLastPage, isFetching, setIsFetching];
};
