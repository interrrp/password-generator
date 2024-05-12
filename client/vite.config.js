import { defineConfig } from "vite";

export default defineConfig({
  server: {
    proxy: {
      "/api/v1": {
        target: "http://localhost",
        rewrite: (path) => path.replace("/api/v1", ""),
      },
    },
  },
});
