<template>
  <v-app>
    <v-main>
      <v-container class="container">
        <login-form />
      </v-container>
    </v-main>
  </v-app>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import LoginForm from '@/components/LoginForm.vue'

export default defineComponent({
  name: 'LoginView',
  components: {
    LoginForm,
  },
});
</script>
