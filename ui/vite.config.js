import { defineConfig } from "vite";
import preact from "@preact/preset-vite";

// https://vitejs.dev/config/
export default defineConfig(({ command, mode }) => {
  return {
    plugins: [preact()],
    envDir: command === "build" ? "../" : ".",
    define: {
      APP_VERSION: JSON.stringify(process.env.npm_package_version),
      APP_BUILD_DATE: Date.now(),
    },
    assetsInclude: ["**/*.md"],
  };
});
