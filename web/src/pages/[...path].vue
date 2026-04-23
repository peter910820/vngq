<template>
  <v-container class="page-shell d-flex align-center justify-center py-12">
    <v-card class="glass-card pa-6 text-center" max-width="520" variant="tonal">
      <v-card-title class="text-h4 mb-2">404</v-card-title>
      <v-card-subtitle class="mb-4">找不到這個頁面</v-card-subtitle>
      <v-card-text class="mb-4">
        你前往的路由不存在，請檢查網址或回到首頁。
      </v-card-text>
      <v-card-actions class="justify-center">
        <v-btn to="/" size="large">回到首頁</v-btn>
      </v-card-actions>
    </v-card>
  </v-container>
</template>
