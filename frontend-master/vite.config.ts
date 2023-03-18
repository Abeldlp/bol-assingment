import federation from "@originjs/vite-plugin-federation";
import vue from "@vitejs/plugin-vue";
import { defineConfig } from "vite";

export default defineConfig({
  plugins: [
    vue(),
    federation({
      name: "frontendMaster",
      filename: "remoteEntry.js",
      remotes: {
        mancala: "http://localhost:3001/assets/remoteEntry.js",
      },
      shared: ["vue"],
    }),
  ],
  server: {
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
