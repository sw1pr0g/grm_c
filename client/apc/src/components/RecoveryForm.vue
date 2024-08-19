<template>
  <v-card class="card">
    <v-card-title class="headline">Восстановление пароля</v-card-title>
    <v-card-text>
      <v-form ref="form" v-model="valid">
        <v-text-field
          v-model="email"
          :rules="emailRules"
          label="Электронная почта"
          required
        ></v-text-field>
      </v-form>
    </v-card-text>
    <v-card-actions>
      <v-spacer></v-spacer>
      <v-btn color="primary" @click="recoverPassword">Отправить</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';

export default defineComponent({
  name: 'PasswordRecoveryForm',
  setup() {
    const email = ref('');
    const valid = ref(false);

    // Правила для проверки электронной почты
    const emailRules = [
      (v: string) => !!v || 'Электронная почта обязательна',
      (v: string) => /.+@.+\..+/.test(v) || 'Электронная почта должна быть действительной'
    ];

    // Функция отправки запроса на восстановление пароля
    const recoverPassword = () => {
      if (valid.value) {
        // Логика отправки запроса на восстановление пароля
        console.log('Отправка ссылки для восстановления пароля на email:', email.value);
      } else {
        console.log('Форма невалидна');
      }
    };

    return {
      email,
      valid,
      emailRules,
      recoverPassword,
    };
  },
});
</script>

<style scoped>
</style>
