<script lang="ts" setup>
  const { data, refresh, error, pending } = await useFetch(() => `http://localhost:8080`, {
    lazy: true,
    server: true,
  })
  console.log(data.value)
</script>
<template>
  <div>
    {{ data }}
    <NuxtWelcome />
  </div>
</template>
