import PocketBase from "pocketbase";
import {
  getBaseUrl,
  getDisplayDateFromFormat,
  getLocalDateFromFormat,
  getPocketBaseFileUrl,
} from "./helpers.js";

const client = new PocketBase(getBaseUrl());

export const login = async (password) => {
  const email =
    import.meta.env.VITE_VIEWER_USER === undefined
      ? `${password}.user@babygramz.com`
      : import.meta.env.VITE_VIEWER_USER;
  return await client.collection("users").authWithPassword(email, password);
};

export const logOut = () => {
  client.authStore.clear();
};

export const isUserLoggedIn = () => {
  return client.authStore.isValid;
};

export const getPhotos = async (page, sortOrder, perPage = 10) => {
  try {
    const photoResults = await client
      .collection("photos")
      .getList(page, perPage, {
        sort: sortOrder,
      });

    const maxPages = Math.ceil(photoResults.totalItems / photoResults.perPage);
    if (page > maxPages) {
      return {
        page: photoResults.page,
        maxPages,
        photoData: [],
      };
    }

    const photoData = photoResults.items.map((item) => {
      return {
        id: item.id,
        description: item.description,
        displayDate: getDisplayDateFromFormat(
          item.dateTaken === "" ? item.created : item.dateTaken
        ),
        altDate: getLocalDateFromFormat(
          item.dateTaken === "" ? item.created : item.dateTaken
        ),
        url: getPocketBaseFileUrl(item.id, item.file, item.orientation),
      };
    });

    return {
      page: photoResults.page,
      maxPages,
      photoData,
    };
  } catch (e) {
    throw e;
  } finally {
    await client.collection("users").authRefresh();
  }
};

export const signupForNotifications = async (email, name) => {
  return await client.collection("subscribers").create({ email, name });
};
