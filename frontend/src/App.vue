<script lang="ts" setup>

import {onMounted} from "vue";
import {useSettingsStore} from "./stores/SettingsStore";
import {Timezone} from "../wailsjs/go/config/appConfig";

const store = useSettingsStore()

onMounted(() => {
  Timezone().then((loc: string) => {
    store.timezone = loc
  })
})
</script>

<template>
  <v-app>
    <v-app-bar :elevation="2" color="blue-grey-lighten-5">
      <v-app-bar-title>CoinTracking CSV Export File Converter</v-app-bar-title>
      <template v-slot:append>
        <router-link to="/settings" custom v-slot="{ navigate }">
          <v-btn @click="navigate" icon="mdi-cog-outline" role="link"></v-btn>
        </router-link>
      </template>
    </v-app-bar>
    <v-main>
      <v-container class="my-4 pa-3">
        <router-view></router-view>
      </v-container>
    </v-main>
  </v-app>
</template>

<style>
</style>
