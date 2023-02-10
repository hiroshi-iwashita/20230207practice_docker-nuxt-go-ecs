<script setup lang="ts">
  // const { data, refresh, error, pending } = await useFetch(() => `http://localhost:8080/users`, {
  //   lazy: true,
  //   server: true,
  // })
  // console.log(data.value)

  const result = ref()
  const fetchData = async () => {
    result.value = await fetch("http://localhost:8080/users")
    console.log(result.value)
  }

  fetchData()

  // const payload = {
  //   email: "fsdfas@dfdsf.com",
  //   password: "fasdfjfsdf",
  // }

  // const requestOptions = {
  //   method: "POST",
  //   body: JSON.stringify(payload),
  // }

  // fetch("http://localhost:8080/users/login", requestOptions).then((response) =>
  //   console.log(response)
  // )
</script>
<template>
  <div>
    <NuxtWelcome />
  </div>
</template>
