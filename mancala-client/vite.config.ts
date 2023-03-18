import federation from "@originjs/vite-plugin-federation";
import vue from "@vitejs/plugin-vue";
import { defineConfig } from "vite";

export default defineConfig({
  plugins: [
    vue(),
    federation({
      name: "mancala",
      filename: "remoteEntry.js",
      exposes: {
        "./Main": "./src/App.vue",
      },
      shared: ["vue"],
    }),
  ],
  server: {
    port: 3001,
  },
  preview: {
    port: 3001,
  },
  build: {
    minify: false,
    target: "esnext",
    cssCodeSplit: false,
  },
});
