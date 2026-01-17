import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import { tanstackRouter } from "@tanstack/router-plugin/vite";
import path from "path";

// https://vite.dev/config/
export default defineConfig({
  resolve: {
    alias: {
      "@assets": path.resolve(__dirname, "./src/assets"),
      "@common": path.resolve(__dirname, "./src/common"),
      "@connect": path.resolve(__dirname, "./src/connect"),
      "@features": path.resolve(__dirname, "./src/features"),
    },
  },
  plugins: [
    tanstackRouter({
      target: "react",
      autoCodeSplitting: true,
    }),
    react(),
    /* React compiler
    react({
      babel: {
        plugins: [['babel-plugin-react-compiler']],
      },
    }),
    */
  ],
});
