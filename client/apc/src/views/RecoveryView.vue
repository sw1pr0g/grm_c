<template>
  <v-app>
    <v-main>
      <v-container class="container">
        <recovery-form />
      </v-container>
    </v-main>
  </v-app>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import RecoveryForm from '@/components/RecoveryForm.vue'

export default defineComponent({
  name: 'RecoveryView',
  components: {
    RecoveryForm,
  },
});
</script>

<style scoped>
.container {
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>