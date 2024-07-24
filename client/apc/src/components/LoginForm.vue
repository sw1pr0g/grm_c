<template>
  <v-card>
    <v-card-title class="headline">Вход в систему</v-card-title>
    <v-card-text>
      <v-form ref="form" v-model="valid">
        <v-text-field
          v-model="email"
          :rules="emailRules"
          label="Электронная почта"
          required
        ></v-text-field>
        <v-text-field
          v-model="password"
          :rules="passwordRules"
          label="Пароль"
          type="password"
          required
        ></v-text-field>
      </v-form>
    </v-card-text>
    <v-card-actions>
      <v-spacer></v-spacer>
      <v-btn color="primary" @click="login">Войти</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'

export default defineComponent({
  name: 'LoginForm',
  setup() {
    const email = ref('')
    const password = ref('')
    const valid = ref(false)

    const emailRules = [
      (v: string) => !!v || 'Электронная почта обязательна',
      (v: string) => /.+@.+\..+/.test(v) || 'Электронная почта должна быть действительной'
    ]

    const passwordRules = [
      (v: string) => !!v || 'Пароль обязателен',
      (v: string) => v.length >= 6 || 'Пароль должен быть не менее 6 символов'
    ]

    const login = () => {
      if (valid.value) {
        // Логика аутентификации
        console.log('Email:', email.value)
        console.log('Password:', password.value)
      } else {
        console.log('Форма невалидна')
      }
    }

    return {
      email,
      password,
      valid,
      emailRules,
      passwordRules,
      login
    }
  }
})
</script>

<style scoped>
.v-card {
  max-width: 1000px;
  min-width: 600px;
  width: 700px;
  margin: 100px auto;
}
</style>
