export default defineEventHandler(async (event) => {
  // const { code } = event.context.params
  const { apiBaseURL } = useRuntimeConfig()
  console.log(apiBaseURL)

  const { data, pending, error, refresh } = await $fetch(() => `/`, {
    baseURL: apiBaseURL,
    lazy: true,
    server: true,
    pick: ["data"],
  })

  if (process.server) console.log(data)

  return data
})
