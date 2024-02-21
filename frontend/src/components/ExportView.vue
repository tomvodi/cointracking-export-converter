<script setup lang="ts">

import ExportFilesList from "./ExportFilesList.vue";
import AddExportFile from "./AddExportFile.vue";
import {ExportToBlockpitXlsx} from "../../wailsjs/go/cointracking/ct";
import {ref} from "vue";
import {common} from "../../wailsjs/go/models";

const exportEnabled = ref(false)
const saveBlockpitFile = async () => {
  ExportToBlockpitXlsx().catch((reason: any) => {
    console.log(`failed saving blockpit file: ${reason}`)
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

<style scoped>

</style>