// @ts-check
const { test, expect } = require("@playwright/test");

test("Can access basic routes", async ({ page }) => {
  await page.goto("/");
  await expect(page).toHaveTitle("Babygramz");

  await page.goto("/login");
  await expect(page.getByRole("heading")).toHaveText("Babygramz Access");

  await page.goto("/about");
  await expect(page.getByRole("heading")).toHaveText("About Babygramz");
});

test("Unauthenticated users cannot access auth-protected routes", async ({
  page,
}) => {
  await page.goto("/feed");
  await expect(page).toHaveURL("/login");

  await page.goto("/uploader/dashboard");
  await expect(page).toHaveURL("/uploader/login");
});
