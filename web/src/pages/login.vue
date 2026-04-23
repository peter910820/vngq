<template>
  <v-container class="page-shell login-page">
    <v-row justify="center" class="mt-12">
      <v-col cols="12" sm="10" md="7" lg="5">
        <v-card class="glass-card pa-4 pa-sm-6" variant="tonal">
          <v-card-title class="text-h5 font-weight-bold">登入</v-card-title>
          <v-card-subtitle class="mt-1">歡迎回來，請輸入你的帳號資訊</v-card-subtitle>

          <v-card-text class="pt-6">
            <v-form>
              <v-text-field
                label="Email"
                type="email"
                variant="outlined"
                prepend-inner-icon="mdi-email-outline"
                autocomplete="username"
                class="mb-3"
              />
              <v-text-field
                label="密碼"
                type="password"
                variant="outlined"
                prepend-inner-icon="mdi-lock-outline"
                autocomplete="current-password"
                class="mb-2"
              />

              <div class="d-flex align-center justify-space-between mb-5">
                <v-checkbox label="記住我" density="compact" hide-details />
                <v-btn variant="text" size="small">忘記密碼？</v-btn>
              </div>

              <v-btn block color="primary" size="large" class="mb-3">登入</v-btn>
              <v-btn block variant="outlined" size="large" to="/">
                回到首頁
              </v-btn>
            </v-form>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<style scoped>
.login-page {
  padding-top: 20px;
}
</style>
