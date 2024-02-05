<script setup lang="ts">
import {computed} from 'vue'
import {OpenExportFile} from "../../wailsjs/go/cointracking/ct";
import TimezoneSelector from "./TimezoneSelector.vue";
import {useSettingsStore} from "../stores/SettingsStore";
import TitledPanel from "./TitledPanel.vue";

const store = useSettingsStore()

const timezoneEmpty = computed(() => {
  return store.timezone.length == 0
})

const selectFile = async () => {
  OpenExportFile(store.timezone).catch((reason: any) => {
    console.log("error selecting file: " + reason)
  })
}

const setTimezone = async (newTz: string) => {
  store.timezone = newTz
}

</script>

<template>
  <TitledPanel title="Add a CoinTracking export file">
    <TimezoneSelector
        :selected-timezone="store.timezone"
        @timezoneChanged="setTimezone"></TimezoneSelector>
    <v-btn
        :disabled="timezoneEmpty"
        @click="selectFile"
    >Select File
    </v-btn>
  </TitledPanel>

</template>

<style scoped>

</style>