<script setup lang="ts">

import {inject, ref} from "vue";
import {common} from "@wails/go/models";
import {useSnackbarStore} from "@/stores/snackbarStore";
import AddExportFile from "@/components/AddExportFile.vue";
import ExportFilesList from "@/components/ExportFilesList.vue";
import {useSettingsStore} from "@/stores/settingsStore";
import {WailsApi} from "@/wails/wails_api";
import {wailsClientInjKey} from "@/injection_keys";

const exportEnabled = ref(false)
const snackStore = useSnackbarStore()
const settingsStore = useSettingsStore()
const wailsClient: WailsApi = inject<WailsApi>(wailsClientInjKey) as WailsApi

const saveBlockpitFile = async () => {
  wailsClient.ExportToBlockpitXlsx().catch((reason: any) => {
    snackStore.showError(`failed saving blockpit file: ${reason}`)
  })
}

const exportedFilesChanged = async (files: Array<common.ExportFileInfo>) => {
  exportEnabled.value = files.length > 0
}

const selectFile = async () => {
  wailsClient.OpenExportFile(settingsStore.timezone).catch((reason: any) => {
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