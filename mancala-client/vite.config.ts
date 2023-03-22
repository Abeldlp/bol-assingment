import federation from "@originjs/vite-plugin-federation";
import vue from "@vitejs/plugin-vue";
import { defineConfig } from "vite";
import { quasar, transformAssetUrls } from "@quasar/vite-plugin";

export default defineConfig({
  plugins: [
    vue({
      template: { transformAssetUrls },
    }),
    quasar({
      sassVariables: "src/quasar-variables.sass",
    }),
    federation({
      name: "mancala",
      filename: "remoteEntry.js",
      exposes: {
        "./Main": "./src/App.vue",
      },
      shared: ["vue", "quasar"],
    }),
  ],
  server: {
    host: true,
    port: 3000,
  },
  preview: {
    port: 3000,
  },
  build: {
    minify: false,
    target: "esnext",
    cssCodeSplit: false,
  },
});
