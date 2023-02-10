import { DateTime } from "luxon";

export function getBaseUrl() {
  return import.meta.env.VITE_POCKETBASE_URL === undefined
    ? window.location.origin
    : import.meta.env.VITE_POCKETBASE_URL;
}

export function getPocketBaseFileUrl(recordId, filename, size = "large") {
  return `${getBaseUrl()}/api/files/photos/${recordId}/${filename}?size=${size}`;
}

export function getUniqueArrayBy(array, key) {
  return [...new Map(array.map((item) => [item[key], item])).values()];
}

export function getDisplayDateFromFormat(dateString) {
  const nowDate = DateTime.now();
  const photoDate = DateTime.fromSQL(dateString, {
    zone: "UTC",
    setZone: true,
  });

  const diffInDays = nowDate.diff(photoDate, "days");

  if (diffInDays.days > 7) {
    return photoDate.setZone("system").toLocaleString(DateTime.DATE_MED);
  } else {
    return photoDate.setZone("system").toRelative();
  }
}

export function getLocalDateFromFormat(
  dateString,
  format = "yyyy-MM-dd HH:mm:ss.SSS"
) {
  return DateTime.fromFormat(dateString, format, {
    zone: "UTC",
    setZone: true,
  })
    .setZone("system")
    .toLocaleString(DateTime.DATETIME_SHORT);
}
