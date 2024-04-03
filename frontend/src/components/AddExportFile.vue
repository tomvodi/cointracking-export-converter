<script setup lang="ts">
import {useSettingsStore} from "@/stores/settingsStore";
import TitledPanel from "./TitledPanel.vue";

const store = useSettingsStore()

const $emit = defineEmits<{
  selectFile: [],
}>()

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
        @click="$emit('selectFile')"
    >Select File
    </v-btn>
  </TitledPanel>

</template>

<style scoped>

</style>