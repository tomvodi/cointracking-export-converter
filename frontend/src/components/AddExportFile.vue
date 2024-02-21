<script setup lang="ts">
import {OpenExportFile} from "../../wailsjs/go/cointracking/ct";
import {useSettingsStore} from "../stores/settingsStore";
import TitledPanel from "./TitledPanel.vue";

const store = useSettingsStore()


const selectFile = async () => {
  OpenExportFile(store.timezone).catch((reason: any) => {
    console.log("error selecting file: " + reason)
  })
}

</script>
<template>
  <v-alert
      v-if="store.timezoneEmpty"
      title="Missing timezone configuration"
      type="info"
      class="mb-5"
  >
    <p>The timezone for the CoinTracking export files has not been set yet.
      Please got the
      <router-link to="/settings">Settings</router-link>
      and configure the
      timezone with which the file(s) have been exported with.
    </p>
  </v-alert>
  <TitledPanel title="Add a CoinTracking export file">
    <v-btn
        :disabled="store.timezoneEmpty"
        @click="selectFile"
    >Select File
    </v-btn>
  </TitledPanel>

</template>

<style scoped>

</style>