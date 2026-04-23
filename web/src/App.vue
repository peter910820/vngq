<script setup lang="ts">
import ArcNavbar from './components/ArcNavbar.vue'
</script>

<template>
  <v-app>
    <v-main>
      <ArcNavbar />
      <RouterView />
    </v-main>
  </v-app>
</template>
