<template>
    <v-app>
      <v-main>
        <v-container class="container">
          <signup-form />
        </v-container>
      </v-main>
    </v-app>
  </template>
  
  <script lang="ts">
  import { defineComponent } from 'vue'
  import SignupForm from '@/components/SignupForm.vue'
  
  export default defineComponent({
    name: 'SignupView',
    components: {
      SignupForm
    }
  })
  </script>