// @ts-check
const { test, expect } = require("@playwright/test");

test.use({ storageState: "playwright/.auth/viewer.json" });

test("Can access auth-protected viewer routes", async ({ page }) => {
  await page.goto("/");
  await expect(page).toHaveURL("/feed");

  await page.goto("/signup");
  await expect(page.getByRole("heading")).toHaveText("Notifications Signup");
});

test("Can see photos in feed", async ({ page }) => {
  await page.goto("/feed");

  const photo = await page.getByRole("img", { name: "Shows Brewing Coffee" });
  await expect(photo).toBeVisible();

  const photoHeading = await page.getByRole("heading", {
    name: "Brewing coffee",
  });
  await expect(photoHeading).toBeVisible();

  const dateBadge = await page.getByText("Apr 6, 2023");
  await expect(dateBadge).toBeVisible();
  await expect(dateBadge).toHaveAttribute("title", /4\/6\/2023, \d{1,2}:55 PM/);
});

test("Cannot access uploader area", async ({ page }) => {
  await page.goto("/uploader/dashboard");
  await expect(page).toHaveURL("/uploader/login");
});

test.skip("Can scroll through large quantity of photos", async ({ page }) => {
  // TODO: Write test to load more content and add enough content to dummy data to facilitate this
});
