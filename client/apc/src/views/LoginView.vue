<template>
  <v-app>
    <LoginForm />
  </v-app>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import LoginForm from '../components/LoginForm.vue'

export default defineComponent({
  name: 'LoginView',
  components: {
    LoginForm
  }
})
</script>

<style scoped>
.v-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 200vh;
  width: 300vw;
}
</style>
