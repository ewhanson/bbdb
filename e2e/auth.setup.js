// auth.setup.js
const { test } = require("@playwright/test");

const viewerFile = "playwright/.auth/viewer.json";

test("authenticate viewer", async ({ page }) => {
  // Perform authentication steps
  await page.goto("/login");
  await page.getByPlaceholder("Password").fill("test_viewer");
  await page.getByRole("button", { name: "Submit" }).click();

  // Wait until the page receives cookies
  await page.waitForURL("/feed");

  // End of authentication steps
  await page.context().storageState({ path: viewerFile });
});

const uploaderFile = "playwright/.auth/uploader.json";

test("authenticate uploader", async ({ page }) => {
  // Perform authentication steps
  await page.goto("/uploader/login");
  await page
    .getByPlaceholder("Enter your email")
    .fill("test_uploader@babygramz.test");
  await page.getByPlaceholder("Enter your password").fill("password1234");
  await page.getByRole("button", { name: "Submit" }).click();

  // Wait until the page receives cookies
  await page.waitForURL("/uploader/dashboard");

  // End of authentication steps
  await page.context().storageState({ path: uploaderFile });
});
