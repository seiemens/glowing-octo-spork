// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  modules: ['nuxt-security'],
  css: ['@/styles/global.css'],
  ssr:false,
  security: {
    headers: {
      contentSecurityPolicy: {
        'img-src': ['https://chart.googleapis.com/']
      }
    }
  }
});
