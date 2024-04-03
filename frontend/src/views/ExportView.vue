<script setup lang="ts">

import {ExportToBlockpitXlsx} from "@wails/go/blockpit/bp";
import {ref} from "vue";
import {common} from "@wails/go/models";
import {useSnackbarStore} from "@/stores/snackbarStore";
import AddExportFile from "@/components/AddExportFile.vue";
import ExportFilesList from "@/components/ExportFilesList.vue";
import {OpenExportFile} from "@wails/go/cointracking/ct";
import {useSettingsStore} from "@/stores/settingsStore";

const exportEnabled = ref(false)
const snackStore = useSnackbarStore()
const settingsStore = useSettingsStore()

const saveBlockpitFile = async () => {
  ExportToBlockpitXlsx().catch((reason: any) => {
    snackStore.showError(`failed saving blockpit file: ${reason}`)
  })
}

const exportedFilesChanged = async (files: Array<common.ExportFileInfo>) => {
  exportEnabled.value = files.length > 0
}

const selectFile = async () => {
  OpenExportFile(settingsStore.timezone).catch((reason: any) => {
    snackStore.showError(reason)
  })
}

</script>

<template>
  <AddExportFile @selectFile="selectFile"/>
  <ExportFilesList
      class="mt-4"
      @exportFilesChanged="exportedFilesChanged"/>
  <v-btn
      class="mt-4 mx-5"
      @click="saveBlockpitFile"
      :disabled="!exportEnabled"
  >Save Blockpit File
  </v-btn>
</template>