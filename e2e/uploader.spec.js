// @ts-check
const { test, expect } = require("@playwright/test");

test.use({ storageState: "playwright/.auth/uploader.json" });

// TODO: This is an exact duplicate of the viewer test
test("Can view photo feed", async ({ page }) => {
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

test("Can upload new photo", async ({ page }, workerInfo) => {
  await page.goto("/uploader/dashboard");

  await page.getByPlaceholder("Enter a description").click();
  await page
    .getByPlaceholder("Enter a description")
    .fill(`Forest mushrooms - ${workerInfo.project.name}`);
  await page
    .locator('input[type="file"]')
    .setInputFiles("./e2e/fixtures/mushroom.jpeg");
  await page.getByRole("button", { name: "Submit" }).click();
  await expect(page.getByText("Photo upload successful!")).toBeVisible();

  await page.getByRole("link", { name: /Babygramz/ }).click();

  const photo = await page.getByRole("img", {
    name: `Shows Forest mushrooms - ${workerInfo.project.name}`,
  });
  await expect(photo).toBeVisible();

  const photoHeading = await page.getByRole("heading", {
    name: `Forest mushrooms - ${workerInfo.project.name}`,
  });
  await expect(photoHeading).toBeVisible();

  const dateBadge = await page.getByText("Dec 14, 2022");
  await expect(dateBadge).toBeVisible();
  await expect(dateBadge).toHaveAttribute(
    "title",
    /12\/14\/2022, \d{1,2}:22 PM/
  );
});
