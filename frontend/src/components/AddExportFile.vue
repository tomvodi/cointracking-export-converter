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
  <v-row>
    <v-col>
      <v-sheet class="pa-2" border>
        <p class="text-h4 pb-2 text-left">Add a CoinTracking export file</p>
        <TimezoneSelector
            :selected-timezone="store.timezone"
            @timezoneChanged="setTimezone"></TimezoneSelector>
        <v-row align="start">
          <v-col class="text-left">
            <v-btn
                :disabled="timezoneEmpty"
                @click="selectFile"
            >Add CoinTracking Export File
            </v-btn>
          </v-col>
          <v-spacer></v-spacer>
        </v-row>
      </v-sheet>
    </v-col>
  </v-row>
</template>

<style scoped>

</style>