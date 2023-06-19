import PocketBase from "pocketbase";
import {
  getBaseUrl,
  getDisplayDateFromFormat,
  getLocalDateFromFormat,
  getPocketBaseFileUrl,
} from "./helpers.js";

const client = new PocketBase(getBaseUrl());

export const viewerLogin = async (password) => {
  const email =
    import.meta.env.VITE_VIEWER_USER === undefined
      ? `${password}.user@babygramz.com`
      : import.meta.env.VITE_VIEWER_USER;
  return await login(email, password);
};

export const login = async (usernameOrEmail, password) => {
  return await client
    .collection("users")
    .authWithPassword(usernameOrEmail, password);
};

export const logOut = () => {
  client.authStore.clear();
};
export const isViewerLoggedIn = () => {
  return client.authStore.isValid;
};

export const isUploaderLoggedIn = () => {
  return (
    client.authStore.isValid && client.authStore.model?.role === "uploader"
  );
};

export const getPhotosByTag = async (tagName, page, perPage = 10) => {
  try {
    const photoResults = await client.send(`/api/bb/tags/${tagName}`, {
      params: {
        sort: "-dateTaken",
        expand: "tags",
        fields:
          "created,dateTaken,description,file,id,expand.tags.id,expand.tags.name",
      },
    });

    // TODO: Refactor to remove duplication
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
        url: getPocketBaseFileUrl(item.id, item.file),
        isNew: item.expand["photos_queue(photo)"] !== undefined,
        tags:
          item.expand["tags"]?.map((tag) => {
            return {
              id: tag.id,
              name: tag.name,
            };
          }) ?? [],
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

export const getMainFeedPhotos = async (page, perPage = 10) => {
  try {
    const photoResults = await client
      .collection("photos")
      .getList(page, perPage, {
        sort: "-dateTaken",
        expand: "photos_queue(photo),tags",
        fields:
          "created,dateTaken,description,file,id,expand.photos_queue(photo).id,expand.tags.id,expand.tags.name",
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
        url: getPocketBaseFileUrl(item.id, item.file),
        isNew: item.expand["photos_queue(photo)"] !== undefined,
        tags:
          item.expand["tags"]?.map((tag) => {
            return {
              id: tag.id,
              name: tag.name,
            };
          }) ?? [],
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

export const postPhoto = async (description, file, dateTime, tags) => {
  const formData = new FormData();

  formData.append("description", description);
  formData.append("file", file);

  if (dateTime !== "") {
    formData.append("dateTaken", dateTime);
  }

  const parsedTags = parseTagsString(tags);
  const tagIds = await getTagIdsFromString(parsedTags);
  tagIds.forEach((item) => formData.append("tags", item));

  return client.collection("photos").create(formData);
};

export const signupForNotifications = async (email, name) => {
  return await client.collection("subscribers").create({ email, name });
};

/**
 * Gets existing/creates new tags
 *
 * @param {string[]} tags
 * @throws An error message with failed tags
 * @return {Promise<string[]>} Tag IDs
 */
export const getTagIdsFromString = async (tags) => {
  const tagIds = [];
  const failedTagNames = [];

  for (let tagName of tags) {
    try {
      const existingTag = await client
        .collection("tags")
        .getFirstListItem(`name="${tagName}"`);
      tagIds.push(existingTag.id);
    } catch (e) {
      if (e.status === 404) {
        try {
          const newTag = await client
            .collection("tags")
            .create({ name: tagName });
          tagIds.push(newTag.id);
        } catch (e) {
          failedTagNames.push(tagName);
        }
      }
    }
  }

  if (failedTagNames.length > 0) {
    throw new Error(`Failed to create tag(s) for ${failedTagNames.join(", ")}`);
  }

  return tagIds;
};

export const getHasNewPhotos = async () => {
  const records = await client.collection("photos_queue").getFullList({
    fields: "id",
  });

  return records.length > 0;
};

/**
 * Helper function to take string to array
 *
 * @param tagsString
 * @return {string[]}
 */
const parseTagsString = (tagsString) => {
  return tagsString.split(",").map((item) => item.trim());
};
