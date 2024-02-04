<script setup lang="ts">
import {computed} from 'vue'
import {OpenExportFile} from "../../wailsjs/go/cointracking/ct";
import TimezoneSelector from "./TimezoneSelector.vue";
import {useSettingsStore} from "../stores/SettingsStore";

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
  <v-sheet class="pa-5" border>
    <v-row>
      <v-col dense>
        <p class="text-h5 mb-3 text-left">Add a CoinTracking export file</p>
        <TimezoneSelector
            :selected-timezone="store.timezone"
            @timezoneChanged="setTimezone"></TimezoneSelector>
        <v-btn
            :disabled="timezoneEmpty"
            @click="selectFile"
        >Select File
        </v-btn>
      </v-col>
    </v-row>
  </v-sheet>
</template>

<style scoped>

</style>