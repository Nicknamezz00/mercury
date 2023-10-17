import react from "@vitejs/plugin-react";
import { resolve } from "path";
import { defineConfig } from "vite";

const devProxyServer = "http://localhost:8081/";
if (process.env.DEV_PROXY_SERVER && process.env.DEV_PROXY_SERVER.length > 0) {
  console.log("Using dev proxy server: ", process.env.DEV_PROXY_SERVER);
}

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    host: "0.0.0.0",
    port: 3001,
    proxy: {
      "^/api": {
        target: devProxyServer,
        xfwd: true,
      },
      "^/o/": {
        target: devProxyServer,
        xfwd: true,
      },
      "^/u/.+/rss.xml": {
        target: devProxyServer,
        xfwd: true,
      },
      "^/explore/rss.xml": {
        target: devProxyServer,
        xfwd: true,
      },
    },
  },
  resolve: {
    alias: {
      "@/": `${resolve(__dirname, "src")}`,
    },
  },
});
