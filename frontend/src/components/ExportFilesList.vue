<script setup lang="ts">

import {EventsOn} from "../../wailsjs/runtime";
import {onMounted, ref} from "vue";
import {GetExportFiles} from "../../wailsjs/go/cointracking/ct";
import {common} from "../../wailsjs/go/models";

const $emit = defineEmits<{
  exportFilesChanged: [files: Array<common.ExportFileInfo>],
}>()

const exportedFiles = ref<Array<common.ExportFileInfo>>([])

onMounted(() => {
  EventsOn("ExportFilesChanged", setExportFiles)

  GetExportFiles().then(setExportFiles).catch((reason: any) => {
    console.log("failed getting export files initially: " + reason)
  })
})

const setExportFiles = async (files: Array<common.ExportFileInfo>) => {
  exportedFiles.value = files
  $emit('exportFilesChanged', exportedFiles.value)
}

</script>

<template>
  <div class="px-5" v-if="exportedFiles.length > 0">
    <p class="text-h6">CoinTracking export files</p>
    <v-list lines="two">
      <v-list-item
          v-for="exportFile in exportedFiles"
          :title="exportFile.fileName"
          :subtitle="`${exportFile.txCount} transactions on ${exportFile.exchanges.length} exchanges`"
      >
        <template v-slot:prepend>
          <v-icon icon="mdi-file-delimited-outline"></v-icon>
        </template>
      </v-list-item>
    </v-list>
  </div>
</template>

<style scoped>

</style>