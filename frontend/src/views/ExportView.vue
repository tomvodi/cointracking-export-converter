<script setup lang="ts">

import {inject} from "vue";
import {useSnackbarStore} from "@/stores/snackbar_store";
import AddExportFile from "@/components/AddExportFile.vue";
import ExportFilesList from "@/components/ExportFilesList.vue";
import {useSettingsStore} from "@/stores/settings_store";
import {WailsApi} from "@/wails/wails_api";
import {wailsClientInjKey} from "@/injection_keys";
import {useApplicationStore} from "@/stores/application_store";

const snackStore = useSnackbarStore()
const settingsStore = useSettingsStore()
const appStore = useApplicationStore()
const wailsClient: WailsApi = inject<WailsApi>(wailsClientInjKey) as WailsApi

const saveBlockpitFile = async () => {
  wailsClient.ExportToBlockpitXlsx().catch((reason: any) => {
    snackStore.showError(`failed saving blockpit file: ${reason}`)
  })
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
      class="mt-4"/>
  <v-btn
      ref="saveBpBtn"
      class="mt-4 mx-5"
      @click="saveBlockpitFile"
      :disabled="!appStore.hasExportFiles"
  >Save Blockpit File
  </v-btn>
</template>