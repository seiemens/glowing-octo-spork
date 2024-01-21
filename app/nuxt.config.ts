// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  css: ['@/styles/global.css'],
  alias: {
    '/api/': 'http://localhost:4000/'
  }
});
