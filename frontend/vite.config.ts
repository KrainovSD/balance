import { presets } from "@krainovsd/presets/rollup";
import vue from "@vitejs/plugin-vue";
import crypto from "crypto";
import path from "path";
import { type PluginOption, defineConfig } from "vite";
import vueDevTools from "vite-plugin-vue-devtools";

const PORT = 3000;
const ssr = process.env.npm_lifecycle_event === "build-ssr:js";
const singleton = process.env.npm_lifecycle_event === "build-singleton:js";
const production = ssr || singleton;

const BALANCE_SERVICE = "http://192.168.135.150:3010"; //

export default defineConfig({
  plugins: [vue({}), vueDevTools(), presets.plugins.visualizer() as PluginOption],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
    },
  },
  publicDir: production ? false : undefined,
  css: {
    modules: {
      generateScopedName: (name, filename, css) => {
        const componentName = filename.split("/").pop()?.split("?").shift()?.replace?.(".vue", "");
        const hash = crypto
          .createHash("sha256")
          .update(name + filename + css)
          .digest("base64")
          .replace(/[^a-z0-9]/gi, "")
          .substring(0, 5);

        return `${componentName}__${name}_${hash}`;
      },
    },
  },
  build: {
    emptyOutDir: true,
    sourcemap: true,
    outDir: ssr ? "../static" : "./dist",
    assetsDir: "./",
    rollupOptions: {
      treeshake: true,
      output: {
        entryFileNames: `bundle.js`,
        chunkFileNames: "[name]-[hash].js",
        assetFileNames: `[name].[ext]`,
      },
    },

    minify: process.env.NODE_ENV === "development" ? false : undefined,
  },
  server: {
    port: PORT,
    proxy: {
      "/api/v1/": {
        target: BALANCE_SERVICE,
        changeOrigin: true,
      },
    },
  },
});
