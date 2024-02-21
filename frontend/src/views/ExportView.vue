<script setup lang="ts">

import {ExportToBlockpitXlsx} from "../../wailsjs/go/cointracking/ct";
import {ref} from "vue";
import {common} from "../../wailsjs/go/models";
import {useSnackbarStore} from "../stores/snackbarStore";
import AddExportFile from "../components/AddExportFile.vue";
import ExportFilesList from "../components/ExportFilesList.vue";

const exportEnabled = ref(false)
const snackStore = useSnackbarStore()

const saveBlockpitFile = async () => {
  ExportToBlockpitXlsx().catch((reason: any) => {
    snackStore.showError(`failed saving blockpit file: ${reason}`)
  })
}

const exportedFilesChanged = async (files: Array<common.ExportFileInfo>) => {
  exportEnabled.value = files.length > 0
}

</script>

<template>
  <AddExportFile/>
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