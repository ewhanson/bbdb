import { DateTime } from "luxon";
import ExifReader from "exifreader";

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
  format = "yyyy-MM-dd HH:mm:ss.SSS'Z'"
) {
  return DateTime.fromFormat(dateString, format, {
    zone: "UTC",
    setZone: true,
  })
    .setZone("system")
    .toLocaleString(DateTime.DATETIME_SHORT);
}

export async function tryGetDateTimeFromImage(file) {
  const tags = await ExifReader.load(file, { includeUnknown: true });

  let dateTime = "";
  let offset = "";

  if (
    tags.hasOwnProperty("DateTimeOriginal") &&
    tags.DateTimeOriginal.value[0]
  ) {
    dateTime = tags.DateTimeOriginal.value[0];
  } else if (tags.hasOwnProperty("DateTime") && tags.DateTime.value[0]) {
    dateTime = tags.DateTime.value[0];
  } else if (
    tags.hasOwnProperty("DateTimeDigitized") &&
    tags.DateTimeDigitized.value[0]
  ) {
    dateTime = tags.DateTimeDigitized.value[0];
  }

  if (
    tags.hasOwnProperty("OffsetTimeOriginal") &&
    tags.OffsetTimeOriginal.value[0]
  ) {
    offset = tags.OffsetTimeOriginal.value[0];
  } else if (tags.hasOwnProperty("OffsetTime") && tags.OffsetTime.value[0]) {
    offset = tags.OffsetTime.value[0];
  } else if (
    tags.hasOwnProperty("OffsetTimeDigitized") &&
    tags.OffsetTimeDigitized.value[0]
  ) {
    offset = tags.OffsetTimeDigitized.value[0];
  }

  return DateTime.fromFormat(dateTime + offset, "yyyy:MM:dd HH:mm:ssZZ")
    .toUTC()
    .toISO();
}
