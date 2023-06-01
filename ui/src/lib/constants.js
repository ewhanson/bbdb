export const constants = {
  ROUTES: {
    HOME: "/",
    LOGIN: "/login",
    FEED: "/feed",
    ABOUT: "/about",
    NOTIFICATIONS: "/signup",
    WHATS_NEW: "/whats-new",
    UNSUBSCRIBE: "/unsubscribe",
    TAG: "/tags/:tagName",
    FOUR_OH_FOUR: "/404",
    UPLOADER: {
      LOGIN: "/uploader/login",
      DASHBOARD: "/uploader/dashboard",
    },
    getTagRoute: function (tagName) {
      return "/tags/" + tagName;
    },
  },
  ICONS: {
    DOTS_HORIZONTAL: "dots_horizontal",
    INFO: "info",
  },
  COOKIE_KEYS: {
    SEEN_NOTIFICATION_PAGE: "seenNotificationPage",
  },
};
