<script setup lang="ts">
import {ref} from 'vue'
import {OpenExportFile} from "../../wailsjs/go/cointracking/ct";
import TimezoneSelector from "./TimezoneSelector.vue";

const selectedFile = ref<string>('')

const selectFile = async () => {
  OpenExportFile().then((value: string) => {
    console.log("file was selected: " + value)
    selectedFile.value = value
  }).catch((reason: any) => {
    console.log("error selecting file: " + reason)
  })
}

const addExportFile = async () => {
  if (!selectedFile.value?.length) {
    return
  }
}
</script>

<template>
  <v-row>
    <v-col>
      <v-sheet class="pa-2" border>
        <p class="text-h4 pb-2 text-left">Add a CoinTracking export file</p>
        <TimezoneSelector></TimezoneSelector>
        <v-text-field clearable
                      variant="outlined" readonly
                      v-model="selectedFile"
                      @click="selectFile">
          <template #append>
            <v-btn @click="selectFile">Select File</v-btn>
          </template>
        </v-text-field>
        <v-row align="start">
          <v-col class="text-left">
            <v-btn @click="addExportFile">Add File</v-btn>
          </v-col>
          <v-spacer></v-spacer>
        </v-row>
      </v-sheet>

    </v-col>
  </v-row>
</template>

<style scoped>

</style>