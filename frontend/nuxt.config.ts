// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  runtimeConfig: {
    app: {
      baseURL: "/",
      buildAssetsDir: "/_nuxt/",
      cdnURL: "",
    },
    public: {
      baseURL: "",
    },
    apiBaseURL: "",
  },
  typescript: {
    strict: true,
  },
  postcss: {
    plugins: {
      tailwindcss: {},
      autoprefixer: {},
    },
  },
})
