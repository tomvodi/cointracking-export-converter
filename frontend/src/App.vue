<script lang="ts" setup>

import {onMounted} from "vue";
import {useSettingsStore} from "./stores/settingsStore";
import {AllTimezones, Timezone} from "../wailsjs/go/config/appConfig";
import {common} from "../wailsjs/go/models";
import Snackbar from "./components/Snackbar.vue";

const store = useSettingsStore()

onMounted(() => {
  const tzPromise = Timezone()
  const allTzPromise = AllTimezones()
  const promises = [
    tzPromise,
    allTzPromise
  ]
  tzPromise.then((loc: string) => {
    store.timezone = loc
  })
  allTzPromise.then((timezones: Array<common.TimezoneData>) => {
    store.allTimezones = timezones
  })

  Promise.allSettled(promises).then((results) => {
    store.settingsLoaded = true
  })
})
</script>

<template>
  <v-overlay
      :model-value="!store.settingsLoaded"
      class="align-center justify-center"
  >
    <v-progress-circular
        color="primary"
        indeterminate
        size="64"
    ></v-progress-circular>
  </v-overlay>
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
        <Snackbar/>
      </v-container>
    </v-main>
  </v-app>
</template>

<style>
</style>
